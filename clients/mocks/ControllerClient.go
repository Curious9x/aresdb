// Code generated by mockery v1.0.0
package mocks

import common "github.com/uber/aresdb/metastore/common"
import mock "github.com/stretchr/testify/mock"

// ControllerClient is an autogenerated mock type for the ControllerClient type
type ControllerClient struct {
	mock.Mock
}

// GetAllSchema provides a mock function with given fields: namespace
func (_m *ControllerClient) GetAllSchema(namespace string) ([]common.Table, error) {
	ret := _m.Called(namespace)

	var r0 []common.Table
	if rf, ok := ret.Get(0).(func(string) []common.Table); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Table)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaHash provides a mock function with given fields: namespace
func (_m *ControllerClient) GetSchemaHash(namespace string) (string, error) {
	ret := _m.Called(namespace)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
