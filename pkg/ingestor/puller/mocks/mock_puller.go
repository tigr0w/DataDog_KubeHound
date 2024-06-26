// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DataPuller is an autogenerated mock type for the DataPuller type
type DataPuller struct {
	mock.Mock
}

type DataPuller_Expecter struct {
	mock *mock.Mock
}

func (_m *DataPuller) EXPECT() *DataPuller_Expecter {
	return &DataPuller_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx, basePath
func (_m *DataPuller) Close(ctx context.Context, basePath string) error {
	ret := _m.Called(ctx, basePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, basePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataPuller_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type DataPuller_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
//   - basePath string
func (_e *DataPuller_Expecter) Close(ctx interface{}, basePath interface{}) *DataPuller_Close_Call {
	return &DataPuller_Close_Call{Call: _e.mock.On("Close", ctx, basePath)}
}

func (_c *DataPuller_Close_Call) Run(run func(ctx context.Context, basePath string)) *DataPuller_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataPuller_Close_Call) Return(_a0 error) *DataPuller_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataPuller_Close_Call) RunAndReturn(run func(context.Context, string) error) *DataPuller_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Extract provides a mock function with given fields: ctx, archivePath
func (_m *DataPuller) Extract(ctx context.Context, archivePath string) error {
	ret := _m.Called(ctx, archivePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, archivePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataPuller_Extract_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Extract'
type DataPuller_Extract_Call struct {
	*mock.Call
}

// Extract is a helper method to define mock.On call
//   - ctx context.Context
//   - archivePath string
func (_e *DataPuller_Expecter) Extract(ctx interface{}, archivePath interface{}) *DataPuller_Extract_Call {
	return &DataPuller_Extract_Call{Call: _e.mock.On("Extract", ctx, archivePath)}
}

func (_c *DataPuller_Extract_Call) Run(run func(ctx context.Context, archivePath string)) *DataPuller_Extract_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataPuller_Extract_Call) Return(_a0 error) *DataPuller_Extract_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataPuller_Extract_Call) RunAndReturn(run func(context.Context, string) error) *DataPuller_Extract_Call {
	_c.Call.Return(run)
	return _c
}

// Pull provides a mock function with given fields: ctx, clusterName, runID
func (_m *DataPuller) Pull(ctx context.Context, clusterName string, runID string) (string, error) {
	ret := _m.Called(ctx, clusterName, runID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, clusterName, runID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, clusterName, runID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterName, runID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataPuller_Pull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pull'
type DataPuller_Pull_Call struct {
	*mock.Call
}

// Pull is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - runID string
func (_e *DataPuller_Expecter) Pull(ctx interface{}, clusterName interface{}, runID interface{}) *DataPuller_Pull_Call {
	return &DataPuller_Pull_Call{Call: _e.mock.On("Pull", ctx, clusterName, runID)}
}

func (_c *DataPuller_Pull_Call) Run(run func(ctx context.Context, clusterName string, runID string)) *DataPuller_Pull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *DataPuller_Pull_Call) Return(_a0 string, _a1 error) *DataPuller_Pull_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataPuller_Pull_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *DataPuller_Pull_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDataPuller interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataPuller creates a new instance of DataPuller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataPuller(t mockConstructorTestingTNewDataPuller) *DataPuller {
	mock := &DataPuller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
