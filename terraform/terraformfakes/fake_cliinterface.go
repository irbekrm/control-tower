// Code generated by counterfeiter. DO NOT EDIT.
package terraformfakes

import (
	"sync"

	"github.com/EngineerBetter/concourse-up/terraform"
)

type FakeCLIInterface struct {
	ApplyStub        func(terraform.InputVars) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 terraform.InputVars
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	BuildOutputStub        func(terraform.InputVars) (terraform.Outputs, error)
	buildOutputMutex       sync.RWMutex
	buildOutputArgsForCall []struct {
		arg1 terraform.InputVars
	}
	buildOutputReturns struct {
		result1 terraform.Outputs
		result2 error
	}
	buildOutputReturnsOnCall map[int]struct {
		result1 terraform.Outputs
		result2 error
	}
	DestroyStub        func(terraform.InputVars) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 terraform.InputVars
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCLIInterface) Apply(arg1 terraform.InputVars) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 terraform.InputVars
	}{arg1})
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.applyReturns
	return fakeReturns.result1
}

func (fake *FakeCLIInterface) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeCLIInterface) ApplyCalls(stub func(terraform.InputVars) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeCLIInterface) ApplyArgsForCall(i int) terraform.InputVars {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCLIInterface) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCLIInterface) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCLIInterface) BuildOutput(arg1 terraform.InputVars) (terraform.Outputs, error) {
	fake.buildOutputMutex.Lock()
	ret, specificReturn := fake.buildOutputReturnsOnCall[len(fake.buildOutputArgsForCall)]
	fake.buildOutputArgsForCall = append(fake.buildOutputArgsForCall, struct {
		arg1 terraform.InputVars
	}{arg1})
	fake.recordInvocation("BuildOutput", []interface{}{arg1})
	fake.buildOutputMutex.Unlock()
	if fake.BuildOutputStub != nil {
		return fake.BuildOutputStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.buildOutputReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCLIInterface) BuildOutputCallCount() int {
	fake.buildOutputMutex.RLock()
	defer fake.buildOutputMutex.RUnlock()
	return len(fake.buildOutputArgsForCall)
}

func (fake *FakeCLIInterface) BuildOutputCalls(stub func(terraform.InputVars) (terraform.Outputs, error)) {
	fake.buildOutputMutex.Lock()
	defer fake.buildOutputMutex.Unlock()
	fake.BuildOutputStub = stub
}

func (fake *FakeCLIInterface) BuildOutputArgsForCall(i int) terraform.InputVars {
	fake.buildOutputMutex.RLock()
	defer fake.buildOutputMutex.RUnlock()
	argsForCall := fake.buildOutputArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCLIInterface) BuildOutputReturns(result1 terraform.Outputs, result2 error) {
	fake.buildOutputMutex.Lock()
	defer fake.buildOutputMutex.Unlock()
	fake.BuildOutputStub = nil
	fake.buildOutputReturns = struct {
		result1 terraform.Outputs
		result2 error
	}{result1, result2}
}

func (fake *FakeCLIInterface) BuildOutputReturnsOnCall(i int, result1 terraform.Outputs, result2 error) {
	fake.buildOutputMutex.Lock()
	defer fake.buildOutputMutex.Unlock()
	fake.BuildOutputStub = nil
	if fake.buildOutputReturnsOnCall == nil {
		fake.buildOutputReturnsOnCall = make(map[int]struct {
			result1 terraform.Outputs
			result2 error
		})
	}
	fake.buildOutputReturnsOnCall[i] = struct {
		result1 terraform.Outputs
		result2 error
	}{result1, result2}
}

func (fake *FakeCLIInterface) Destroy(arg1 terraform.InputVars) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 terraform.InputVars
	}{arg1})
	fake.recordInvocation("Destroy", []interface{}{arg1})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeCLIInterface) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeCLIInterface) DestroyCalls(stub func(terraform.InputVars) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeCLIInterface) DestroyArgsForCall(i int) terraform.InputVars {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCLIInterface) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCLIInterface) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCLIInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.buildOutputMutex.RLock()
	defer fake.buildOutputMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCLIInterface) recordInvocation(key string, args []interface{}) {
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

var _ terraform.CLIInterface = new(FakeCLIInterface)
