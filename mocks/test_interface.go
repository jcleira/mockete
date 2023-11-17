package mocks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockTestInterface struct {
	*DoSomethingMock
	*FinishSomethingMock
}

func NewMockTestInterface(
	doSomething *DoSomethingMock,
	finishSomething *FinishSomethingMock) *MockTestInterface {
	return &MockTestInterface{
		DoSomethingMock:     doSomething,
		FinishSomethingMock: finishSomething,
	}
}

type DoSomethingMock struct {
	mock.Mock
}

func (_m *DoSomethingMock) DoSomething(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DoSomethingMock) Return(returnArguments ...interface{}) *DoSomethingMock {
	if len(_m.Mock.ExpectedCalls) == 0 {
		return _m
	}

	_m.Mock.ExpectedCalls[0].Return(returnArguments...)
	return _m
}

func DoSomething(t *testing.T, ctx context.Context, arg0 string) *DoSomethingMock {
	mock := &DoSomethingMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	mock.On("DoSomething", ctx, arg0).Return(arg0, nil)

	return mock
}

type FinishSomethingMock struct {
	mock.Mock
}

func (_m *FinishSomethingMock) FinishSomething(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func FinishSomething(t *testing.T, ctx context.Context, arg0 string) *FinishSomethingMock {
	mock := &FinishSomethingMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	mock.On("FinishSomething", ctx, arg0)

	return mock
}

func (_m *FinishSomethingMock) Return(returnArguments ...interface{}) *FinishSomethingMock {
	_m.Mock.ExpectedCalls[0].Return(returnArguments...)
	return _m
}

func FinishSomethingNotCalled(t *testing.T, returnArguments ...interface{}) *FinishSomethingMock {
	mock := &FinishSomethingMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	mock.AssertNotCalled(t, "FinishSomething")

	return mock
}
