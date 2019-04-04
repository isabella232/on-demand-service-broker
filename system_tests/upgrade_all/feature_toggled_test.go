// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upgrade_all_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/pborman/uuid"
	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/cf_helpers"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/service_helpers"
)

var _ = Describe("upgrade-all-service-instances errand using all the features available", func() {

	const canaryOrg = "canary_org"
	const canarySpace = "canary_space"
	const standardOrg = "system"
	const standardSpace = "system"

	var (
		brokerInfo         bosh_helpers.BrokerInfo
		uniqueID           string
		nonCanariesDetails []appDetails
		canaryDetails      appDetails
	)

	BeforeEach(func() {
		uniqueID = uuid.New()[:8]

		brokerInfo = bosh_helpers.DeployAndRegisterBroker(
			"feature-toggled-upgrade-"+uniqueID,
			bosh_helpers.BrokerDeploymentOptions{},
			service_helpers.Redis,
			[]string{
				"service_catalog_with_lifecycle.yml",
				"add_canary_filter.yml",
			},
		)

		cf_helpers.CreateOrg(canaryOrg)
		cf_helpers.CreateSpace(canaryOrg, canarySpace)

		cf_helpers.TargetOrgAndSpace(standardOrg, standardSpace)
	})

	AfterEach(func() {
		for _, appDtls := range append(nonCanariesDetails, canaryDetails) {
			cf_helpers.UnbindAndDeleteApp(appDtls.appName, appDtls.serviceName)
		}

		bosh_helpers.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)

		cf_helpers.DeleteSpace(canaryOrg, canarySpace)
		cf_helpers.DeleteOrg(canaryOrg)

		cf_helpers.TargetOrgAndSpace(standardOrg, standardSpace)
	})

	It("succeeds", func() {
		nonCanaryServices := 2
		planName := "redis-with-post-deploy"

		appDtlsCh := make(chan appDetails, nonCanaryServices)
		appPath := cf_helpers.GetAppPath(service_helpers.Redis)

		performInParallel(func() {
			appDtlsCh <- deployService(brokerInfo.ServiceOffering, planName, appPath)
		}, nonCanaryServices)

		close(appDtlsCh)
		for dtls := range appDtlsCh {
			nonCanariesDetails = append(nonCanariesDetails, dtls)
		}

		cf_helpers.TargetOrgAndSpace(canaryOrg, canarySpace)
		canaryDetails = deployService(brokerInfo.ServiceOffering, planName, appPath)
		cf_helpers.TargetOrgAndSpace(standardOrg, standardSpace)

		By("changing the name of instance group and disabling persistence", func() {
			brokerInfo = bosh_helpers.DeployAndRegisterBroker(
				"feature-toggled-upgrade-"+uniqueID,
				bosh_helpers.BrokerDeploymentOptions{},
				service_helpers.Redis,
				[]string{
					"service_catalog_with_lifecycle_updated.yml",
					"add_canary_filter.yml",
				})
		})

		By("running the upgrade-all errand")
		session := bosh_helpers.RunErrand(brokerInfo.DeploymentName, "upgrade-all-service-instances")

		By("upgrading the canary instance first, followed by the rest in parallel")
		Expect(session).To(SatisfyAll(
			gbytes.Say(`\[%s\] Starting to process service instance`, canaryDetails.serviceGUID),
			gbytes.Say(`\[%s\] Result: Service Instance operation success`, canaryDetails.serviceGUID),
			gbytes.Say("FINISHED CANARIES"),
			gbytes.Say(`\[%s\] Starting to process service instance`, nonCanariesDetails[0].serviceGUID),
			gbytes.Say(`\[%s\] Starting to process service instance`, nonCanariesDetails[1].serviceGUID),
			gbytes.Say(`\[%s\] Result: Service Instance operation success`, nonCanariesDetails[0].serviceGUID),
			gbytes.Say("FINISHED PROCESSING Status: SUCCESS"),
		))

		By("checking the other service instance upgrade completed", func() {
			Expect(string(session.Out.Contents())).To(
				ContainSubstring(`[%s] Result: Service Instance operation success`, nonCanariesDetails[1].serviceGUID),
			)
		})

		for _, appDtls := range append(nonCanariesDetails, canaryDetails) {
			By("verifying the update changes were applied to the instance", func() {
				manifest := bosh_helpers.GetManifest(appDtls.serviceDeploymentName)
				instanceGroupProperties := bosh_helpers.FindInstanceGroupProperties(&manifest, "redis")
				Expect(instanceGroupProperties["redis"].(map[interface{}]interface{})["persistence"]).To(Equal("no"))
			})

			By("checking apps still have access to the data previously stored in their service", func() {
				Expect(cf_helpers.GetFromTestApp(appDtls.appURL, "uuid")).To(Equal(appDtls.uuid))
			})

			By("checking lifecycle errands executed as expected", func() {
				expectedBoshTasksOrder := []string{"create deployment", "run errand", "create deployment", "run errand", "create deployment", "run errand"}

				boshTasks := bosh_helpers.TasksForDeployment(appDtls.serviceDeploymentName)
				Expect(boshTasks).To(HaveLen(4))

				for i, task := range boshTasks {
					Expect(task.State).To(Equal(boshdirector.TaskDone))
					Expect(task.Description).To(ContainSubstring(expectedBoshTasksOrder[3-i]))
				}
			})
		}
	})
})