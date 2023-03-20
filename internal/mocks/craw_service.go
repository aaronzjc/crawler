// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/aaronzjc/mu/internal/application/dto"
	mock "github.com/stretchr/testify/mock"
)

// CrawService is an autogenerated mock type for the CrawService type
type CrawService struct {
	mock.Mock
}

type CrawService_Expecter struct {
	mock *mock.Mock
}

func (_m *CrawService) EXPECT() *CrawService_Expecter {
	return &CrawService_Expecter{mock: &_m.Mock}
}

// Craw provides a mock function with given fields: _a0, _a1
func (_m *CrawService) Craw(_a0 context.Context, _a1 *dto.Site) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.Site) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CrawService_Craw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Craw'
type CrawService_Craw_Call struct {
	*mock.Call
}

// Craw is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *dto.Site
func (_e *CrawService_Expecter) Craw(_a0 interface{}, _a1 interface{}) *CrawService_Craw_Call {
	return &CrawService_Craw_Call{Call: _e.mock.On("Craw", _a0, _a1)}
}

func (_c *CrawService_Craw_Call) Run(run func(_a0 context.Context, _a1 *dto.Site)) *CrawService_Craw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dto.Site))
	})
	return _c
}

func (_c *CrawService_Craw_Call) Return(_a0 error) *CrawService_Craw_Call {
	_c.Call.Return(_a0)
	return _c
}

// PickAgent provides a mock function with given fields: _a0, _a1
func (_m *CrawService) PickAgent(_a0 context.Context, _a1 *dto.Site) (*dto.Node, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.Node
	if rf, ok := ret.Get(0).(func(context.Context, *dto.Site) *dto.Node); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.Site) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CrawService_PickAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PickAgent'
type CrawService_PickAgent_Call struct {
	*mock.Call
}

// PickAgent is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *dto.Site
func (_e *CrawService_Expecter) PickAgent(_a0 interface{}, _a1 interface{}) *CrawService_PickAgent_Call {
	return &CrawService_PickAgent_Call{Call: _e.mock.On("PickAgent", _a0, _a1)}
}

func (_c *CrawService_PickAgent_Call) Run(run func(_a0 context.Context, _a1 *dto.Site)) *CrawService_PickAgent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dto.Site))
	})
	return _c
}

func (_c *CrawService_PickAgent_Call) Return(_a0 *dto.Node, _a1 error) *CrawService_PickAgent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewCrawService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCrawService creates a new instance of CrawService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCrawService(t mockConstructorTestingTNewCrawService) *CrawService {
	mock := &CrawService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}