// Code generated by mockery v2.23.2. DO NOT EDIT.

package service

import (
	domain "github.com/pklimuk-eng-thesis/device/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockDeviceService is an autogenerated mock type for the DeviceService type
type MockDeviceService struct {
	mock.Mock
}

type MockDeviceService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeviceService) EXPECT() *MockDeviceService_Expecter {
	return &MockDeviceService_Expecter{mock: &_m.Mock}
}

// GetInfo provides a mock function with given fields:
func (_m *MockDeviceService) GetInfo() *domain.Device {
	ret := _m.Called()

	var r0 *domain.Device
	if rf, ok := ret.Get(0).(func() *domain.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Device)
		}
	}

	return r0
}

// MockDeviceService_GetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInfo'
type MockDeviceService_GetInfo_Call struct {
	*mock.Call
}

// GetInfo is a helper method to define mock.On call
func (_e *MockDeviceService_Expecter) GetInfo() *MockDeviceService_GetInfo_Call {
	return &MockDeviceService_GetInfo_Call{Call: _e.mock.On("GetInfo")}
}

func (_c *MockDeviceService_GetInfo_Call) Run(run func()) *MockDeviceService_GetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceService_GetInfo_Call) Return(_a0 *domain.Device) *MockDeviceService_GetInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceService_GetInfo_Call) RunAndReturn(run func() *domain.Device) *MockDeviceService_GetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleEnabled provides a mock function with given fields:
func (_m *MockDeviceService) ToggleEnabled() *domain.Device {
	ret := _m.Called()

	var r0 *domain.Device
	if rf, ok := ret.Get(0).(func() *domain.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Device)
		}
	}

	return r0
}

// MockDeviceService_ToggleEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleEnabled'
type MockDeviceService_ToggleEnabled_Call struct {
	*mock.Call
}

// ToggleEnabled is a helper method to define mock.On call
func (_e *MockDeviceService_Expecter) ToggleEnabled() *MockDeviceService_ToggleEnabled_Call {
	return &MockDeviceService_ToggleEnabled_Call{Call: _e.mock.On("ToggleEnabled")}
}

func (_c *MockDeviceService_ToggleEnabled_Call) Run(run func()) *MockDeviceService_ToggleEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceService_ToggleEnabled_Call) Return(_a0 *domain.Device) *MockDeviceService_ToggleEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceService_ToggleEnabled_Call) RunAndReturn(run func() *domain.Device) *MockDeviceService_ToggleEnabled_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDeviceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDeviceService creates a new instance of MockDeviceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDeviceService(t mockConstructorTestingTNewMockDeviceService) *MockDeviceService {
	mock := &MockDeviceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
