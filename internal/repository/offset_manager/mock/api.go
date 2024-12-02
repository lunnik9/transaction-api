// Code generated by MockGen. DO NOT EDIT.
// Source: local/transaction/internal/repository/offset_manager (interfaces: BlockOffsetManager)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlockOffsetManager is a mock of BlockOffsetManager interface.
type MockBlockOffsetManager struct {
	ctrl     *gomock.Controller
	recorder *MockBlockOffsetManagerMockRecorder
}

// MockBlockOffsetManagerMockRecorder is the mock recorder for MockBlockOffsetManager.
type MockBlockOffsetManagerMockRecorder struct {
	mock *MockBlockOffsetManager
}

// NewMockBlockOffsetManager creates a new mock instance.
func NewMockBlockOffsetManager(ctrl *gomock.Controller) *MockBlockOffsetManager {
	mock := &MockBlockOffsetManager{ctrl: ctrl}
	mock.recorder = &MockBlockOffsetManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockOffsetManager) EXPECT() *MockBlockOffsetManagerMockRecorder {
	return m.recorder
}

// GetOffset mocks base method.
func (m *MockBlockOffsetManager) GetOffset(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffset", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffset indicates an expected call of GetOffset.
func (mr *MockBlockOffsetManagerMockRecorder) GetOffset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffset", reflect.TypeOf((*MockBlockOffsetManager)(nil).GetOffset), arg0)
}

// GetProcessed mocks base method.
func (m *MockBlockOffsetManager) GetProcessed(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessed", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessed indicates an expected call of GetProcessed.
func (mr *MockBlockOffsetManagerMockRecorder) GetProcessed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessed", reflect.TypeOf((*MockBlockOffsetManager)(nil).GetProcessed), arg0)
}

// SetNext mocks base method.
func (m *MockBlockOffsetManager) SetNext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNext indicates an expected call of SetNext.
func (mr *MockBlockOffsetManagerMockRecorder) SetNext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNext", reflect.TypeOf((*MockBlockOffsetManager)(nil).SetNext), arg0, arg1)
}