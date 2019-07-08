// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	boshdirector "github.com/pivotal-cf/on-demand-service-broker/boshdirector"
)

func AnyBoshdirectorBoshTask() boshdirector.BoshTask {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(boshdirector.BoshTask))(nil)).Elem()))
	var nullValue boshdirector.BoshTask
	return nullValue
}

func EqBoshdirectorBoshTask(value boshdirector.BoshTask) boshdirector.BoshTask {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue boshdirector.BoshTask
	return nullValue
}
