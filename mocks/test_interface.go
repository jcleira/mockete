package mocks

import (
	"context"
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
	ctx  context.Context
	arg0 string

	r_0 string
	r_e error

	notCalled bool
}

func (m *DoSomethingMock) DoSomething(ctx context.Context, arg0 string) (string, error) {
	return m.r_0, m.r_e
}

func (m *DoSomethingMock) Return(r_0 string, r_e error) *DoSomethingMock {
	m.r_0 = r_0
	m.r_e = r_e
	return m
}

func DoSomethingNotCalled() *DoSomethingMock {
	return &DoSomethingMock{
		notCalled: true,
	}
}

func DoSomething(ctx context.Context, arg0 string) *DoSomethingMock {
	return &DoSomethingMock{
		ctx:  ctx,
		arg0: arg0,
	}
}

type FinishSomethingMock struct {
	ctx  context.Context
	arg0 string

	r_e error

	notCalled bool
}

func (m *FinishSomethingMock) FinishSomething(ctx context.Context, arg0 string) error {
	return m.r_e
}

func (m *FinishSomethingMock) Return(r_e error) *FinishSomethingMock {
	m.r_e = r_e
	return m
}

func FinishSomethingNotCalled() *FinishSomethingMock {
	return &FinishSomethingMock{
		notCalled: true,
	}
}

func FinishSomething(ctx context.Context, arg0 string) *FinishSomethingMock {
	return &FinishSomethingMock{
		ctx:  ctx,
		arg0: arg0,
	}
}
