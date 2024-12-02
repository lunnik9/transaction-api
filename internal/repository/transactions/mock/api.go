// Code generated by MockGen. DO NOT EDIT.
// Source: local/transaction/internal/repository/transactions (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain "local/transaction/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddSubscriber mocks base method.
func (m *MockRepository) AddSubscriber(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscriber", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubscriber indicates an expected call of AddSubscriber.
func (mr *MockRepositoryMockRecorder) AddSubscriber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscriber", reflect.TypeOf((*MockRepository)(nil).AddSubscriber), arg0, arg1)
}

// AddSubscriberTransactions mocks base method.
func (m *MockRepository) AddSubscriberTransactions(arg0 context.Context, arg1 map[string][]domain.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscriberTransactions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubscriberTransactions indicates an expected call of AddSubscriberTransactions.
func (mr *MockRepositoryMockRecorder) AddSubscriberTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscriberTransactions", reflect.TypeOf((*MockRepository)(nil).AddSubscriberTransactions), arg0, arg1)
}

// GetSubscriberAddresses mocks base method.
func (m *MockRepository) GetSubscriberAddresses(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriberAddresses", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriberAddresses indicates an expected call of GetSubscriberAddresses.
func (mr *MockRepositoryMockRecorder) GetSubscriberAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriberAddresses", reflect.TypeOf((*MockRepository)(nil).GetSubscriberAddresses), arg0)
}

// GetSubscriberTransactions mocks base method.
func (m *MockRepository) GetSubscriberTransactions(arg0 context.Context, arg1 string) ([]domain.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriberTransactions", arg0, arg1)
	ret0, _ := ret[0].([]domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriberTransactions indicates an expected call of GetSubscriberTransactions.
func (mr *MockRepositoryMockRecorder) GetSubscriberTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriberTransactions", reflect.TypeOf((*MockRepository)(nil).GetSubscriberTransactions), arg0, arg1)
}
