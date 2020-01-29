// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gqlschema "github.com/kyma-incubator/compass/components/provisioner/pkg/gqlschema"
import mock "github.com/stretchr/testify/mock"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ValidateInput provides a mock function with given fields: input
func (_m *Validator) ValidateInput(input gqlschema.ProvisionRuntimeInput) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(gqlschema.ProvisionRuntimeInput) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateTenant provides a mock function with given fields: runtimeID, tenant
func (_m *Validator) ValidateTenant(runtimeID string, tenant string) error {
	ret := _m.Called(runtimeID, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(runtimeID, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
