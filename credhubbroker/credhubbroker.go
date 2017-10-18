package credhubbroker

import (
	"context"
	"fmt"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/apiserver"
)

type CredHubBroker struct {
	apiserver.CombinedBrokers
	credStore CredentialStore
}

func New(broker apiserver.CombinedBrokers, credStore CredentialStore) *CredHubBroker {
	return &CredHubBroker{
		CombinedBrokers: broker,
		credStore:       credStore,
	}
}

func (b *CredHubBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	binding, err := b.CombinedBrokers.Bind(ctx, instanceID, bindingID, details)
	if err != nil {
		return brokerapi.Binding{}, err
	}
	key := fmt.Sprintf("/c/%s/%s/%s/credentials", details.ServiceID, instanceID, bindingID)
	err = b.credStore.Set(key, binding.Credentials)
	if err != nil {
		return brokerapi.Binding{}, err
	}
	return binding, nil
}