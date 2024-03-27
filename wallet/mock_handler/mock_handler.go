// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	wallet "github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	gomock "github.com/golang/mock/gomock"
)

// MockStorer is a mock of Storer interface.
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer.
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance.
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStorer) Create(wallet *wallet.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStorerMockRecorder) Create(wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorer)(nil).Create), wallet)
}

// DeleteOne mocks base method.
func (m *MockStorer) DeleteOne(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockStorerMockRecorder) DeleteOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockStorer)(nil).DeleteOne), id)
}

// UpdateOne mocks base method.
func (m *MockStorer) UpdateOne(update *wallet.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockStorerMockRecorder) UpdateOne(update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockStorer)(nil).UpdateOne), update)
}

// Wallets mocks base method.
func (m *MockStorer) Wallets(filter wallet.Filter) ([]wallet.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wallets", filter)
	ret0, _ := ret[0].([]wallet.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wallets indicates an expected call of Wallets.
func (mr *MockStorerMockRecorder) Wallets(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wallets", reflect.TypeOf((*MockStorer)(nil).Wallets), filter)
}
