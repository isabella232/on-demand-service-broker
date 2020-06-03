// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/broker"
)

type FakeUAAClient struct {
	CreateClientStub        func(string, string) (map[string]string, error)
	createClientMutex       sync.RWMutex
	createClientArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createClientReturns struct {
		result1 map[string]string
		result2 error
	}
	createClientReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	DeleteClientStub        func(string) error
	deleteClientMutex       sync.RWMutex
	deleteClientArgsForCall []struct {
		arg1 string
	}
	deleteClientReturns struct {
		result1 error
	}
	deleteClientReturnsOnCall map[int]struct {
		result1 error
	}
	GetClientStub        func(string) (map[string]string, error)
	getClientMutex       sync.RWMutex
	getClientArgsForCall []struct {
		arg1 string
	}
	getClientReturns struct {
		result1 map[string]string
		result2 error
	}
	getClientReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	UpdateClientStub        func(string, string) (map[string]string, error)
	updateClientMutex       sync.RWMutex
	updateClientArgsForCall []struct {
		arg1 string
		arg2 string
	}
	updateClientReturns struct {
		result1 map[string]string
		result2 error
	}
	updateClientReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) CreateClient(arg1 string, arg2 string) (map[string]string, error) {
	fake.createClientMutex.Lock()
	ret, specificReturn := fake.createClientReturnsOnCall[len(fake.createClientArgsForCall)]
	fake.createClientArgsForCall = append(fake.createClientArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateClient", []interface{}{arg1, arg2})
	fake.createClientMutex.Unlock()
	if fake.CreateClientStub != nil {
		return fake.CreateClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) CreateClientCallCount() int {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return len(fake.createClientArgsForCall)
}

func (fake *FakeUAAClient) CreateClientCalls(stub func(string, string) (map[string]string, error)) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = stub
}

func (fake *FakeUAAClient) CreateClientArgsForCall(i int) (string, string) {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	argsForCall := fake.createClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) CreateClientReturns(result1 map[string]string, result2 error) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = nil
	fake.createClientReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) CreateClientReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = nil
	if fake.createClientReturnsOnCall == nil {
		fake.createClientReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.createClientReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) DeleteClient(arg1 string) error {
	fake.deleteClientMutex.Lock()
	ret, specificReturn := fake.deleteClientReturnsOnCall[len(fake.deleteClientArgsForCall)]
	fake.deleteClientArgsForCall = append(fake.deleteClientArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteClient", []interface{}{arg1})
	fake.deleteClientMutex.Unlock()
	if fake.DeleteClientStub != nil {
		return fake.DeleteClientStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteClientReturns
	return fakeReturns.result1
}

func (fake *FakeUAAClient) DeleteClientCallCount() int {
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	return len(fake.deleteClientArgsForCall)
}

func (fake *FakeUAAClient) DeleteClientCalls(stub func(string) error) {
	fake.deleteClientMutex.Lock()
	defer fake.deleteClientMutex.Unlock()
	fake.DeleteClientStub = stub
}

func (fake *FakeUAAClient) DeleteClientArgsForCall(i int) string {
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	argsForCall := fake.deleteClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) DeleteClientReturns(result1 error) {
	fake.deleteClientMutex.Lock()
	defer fake.deleteClientMutex.Unlock()
	fake.DeleteClientStub = nil
	fake.deleteClientReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUAAClient) DeleteClientReturnsOnCall(i int, result1 error) {
	fake.deleteClientMutex.Lock()
	defer fake.deleteClientMutex.Unlock()
	fake.DeleteClientStub = nil
	if fake.deleteClientReturnsOnCall == nil {
		fake.deleteClientReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteClientReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUAAClient) GetClient(arg1 string) (map[string]string, error) {
	fake.getClientMutex.Lock()
	ret, specificReturn := fake.getClientReturnsOnCall[len(fake.getClientArgsForCall)]
	fake.getClientArgsForCall = append(fake.getClientArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceInstanceClient", []interface{}{arg1})
	fake.getClientMutex.Unlock()
	if fake.GetClientStub != nil {
		return fake.GetClientStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetClientCallCount() int {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return len(fake.getClientArgsForCall)
}

func (fake *FakeUAAClient) GetClientCalls(stub func(string) (map[string]string, error)) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = stub
}

func (fake *FakeUAAClient) GetClientArgsForCall(i int) string {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	argsForCall := fake.getClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) GetClientReturns(result1 map[string]string, result2 error) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = nil
	fake.getClientReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetClientReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = nil
	if fake.getClientReturnsOnCall == nil {
		fake.getClientReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.getClientReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) UpdateClient(arg1 string, arg2 string) (map[string]string, error) {
	fake.updateClientMutex.Lock()
	ret, specificReturn := fake.updateClientReturnsOnCall[len(fake.updateClientArgsForCall)]
	fake.updateClientArgsForCall = append(fake.updateClientArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UpdateClient", []interface{}{arg1, arg2})
	fake.updateClientMutex.Unlock()
	if fake.UpdateClientStub != nil {
		return fake.UpdateClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) UpdateClientCallCount() int {
	fake.updateClientMutex.RLock()
	defer fake.updateClientMutex.RUnlock()
	return len(fake.updateClientArgsForCall)
}

func (fake *FakeUAAClient) UpdateClientCalls(stub func(string, string) (map[string]string, error)) {
	fake.updateClientMutex.Lock()
	defer fake.updateClientMutex.Unlock()
	fake.UpdateClientStub = stub
}

func (fake *FakeUAAClient) UpdateClientArgsForCall(i int) (string, string) {
	fake.updateClientMutex.RLock()
	defer fake.updateClientMutex.RUnlock()
	argsForCall := fake.updateClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) UpdateClientReturns(result1 map[string]string, result2 error) {
	fake.updateClientMutex.Lock()
	defer fake.updateClientMutex.Unlock()
	fake.UpdateClientStub = nil
	fake.updateClientReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) UpdateClientReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.updateClientMutex.Lock()
	defer fake.updateClientMutex.Unlock()
	fake.UpdateClientStub = nil
	if fake.updateClientReturnsOnCall == nil {
		fake.updateClientReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.updateClientReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	fake.updateClientMutex.RLock()
	defer fake.updateClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ broker.UAAClient = new(FakeUAAClient)