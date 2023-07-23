// Code generated by MockGen. DO NOT EDIT.
// Source: stddevapi (interfaces: RandomBackend)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRandomBackend is a mock of RandomBackend interface.
type MockRandomBackend struct {
	ctrl     *gomock.Controller
	recorder *MockRandomBackendMockRecorder
}

// MockRandomBackendMockRecorder is the mock recorder for MockRandomBackend.
type MockRandomBackendMockRecorder struct {
	mock *MockRandomBackend
}

// NewMockRandomBackend creates a new mock instance.
func NewMockRandomBackend(ctrl *gomock.Controller) *MockRandomBackend {
	mock := &MockRandomBackend{ctrl: ctrl}
	mock.recorder = &MockRandomBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRandomBackend) EXPECT() *MockRandomBackendMockRecorder {
	return m.recorder
}

// GetRandomIntegers mocks base method.
func (m *MockRandomBackend) GetRandomIntegers(arg0 context.Context, arg1 int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomIntegers", arg0, arg1)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomIntegers indicates an expected call of GetRandomIntegers.
func (mr *MockRandomBackendMockRecorder) GetRandomIntegers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomIntegers", reflect.TypeOf((*MockRandomBackend)(nil).GetRandomIntegers), arg0, arg1)
}
