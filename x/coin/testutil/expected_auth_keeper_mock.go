// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cosmos/cosmos-sdk/x/auth/keeper (interfaces: AccountKeeperI)

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/adminoid/cosmos-sdk/crypto/types"
	types0 "github.com/adminoid/cosmos-sdk/types"
	types1 "github.com/adminoid/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeperI is a mock of AccountKeeperI interface.
type MockAccountKeeperI struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperIMockRecorder
}

// MockAccountKeeperIMockRecorder is the mock recorder for MockAccountKeeperI.
type MockAccountKeeperIMockRecorder struct {
	mock *MockAccountKeeperI
}

// NewMockAccountKeeperI creates a new mock instance.
func NewMockAccountKeeperI(ctrl *gomock.Controller) *MockAccountKeeperI {
	mock := &MockAccountKeeperI{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeperI) EXPECT() *MockAccountKeeperIMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeperI) GetAccount(arg0 types0.Context, arg1 types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperIMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeperI)(nil).GetAccount), arg0, arg1)
}

// GetNextAccountNumber mocks base method.
func (m *MockAccountKeeperI) GetNextAccountNumber(arg0 types0.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextAccountNumber", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNextAccountNumber indicates an expected call of GetNextAccountNumber.
func (mr *MockAccountKeeperIMockRecorder) GetNextAccountNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextAccountNumber", reflect.TypeOf((*MockAccountKeeperI)(nil).GetNextAccountNumber), arg0)
}

// GetPubKey mocks base method.
func (m *MockAccountKeeperI) GetPubKey(arg0 types0.Context, arg1 types0.AccAddress) (types.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey", arg0, arg1)
	ret0, _ := ret[0].(types.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubKey indicates an expected call of GetPubKey.
func (mr *MockAccountKeeperIMockRecorder) GetPubKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockAccountKeeperI)(nil).GetPubKey), arg0, arg1)
}

// GetSequence mocks base method.
func (m *MockAccountKeeperI) GetSequence(arg0 types0.Context, arg1 types0.AccAddress) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSequence", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSequence indicates an expected call of GetSequence.
func (mr *MockAccountKeeperIMockRecorder) GetSequence(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSequence", reflect.TypeOf((*MockAccountKeeperI)(nil).GetSequence), arg0, arg1)
}

// HasAccount mocks base method.
func (m *MockAccountKeeperI) HasAccount(arg0 types0.Context, arg1 types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAccount indicates an expected call of HasAccount.
func (mr *MockAccountKeeperIMockRecorder) HasAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccount", reflect.TypeOf((*MockAccountKeeperI)(nil).HasAccount), arg0, arg1)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeperI) IterateAccounts(arg0 types0.Context, arg1 func(types1.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", arg0, arg1)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperIMockRecorder) IterateAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeperI)(nil).IterateAccounts), arg0, arg1)
}

// NewAccount mocks base method.
func (m *MockAccountKeeperI) NewAccount(arg0 types0.Context, arg1 types1.AccountI) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountKeeperIMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountKeeperI)(nil).NewAccount), arg0, arg1)
}

// NewAccountWithAddress mocks base method.
func (m *MockAccountKeeperI) NewAccountWithAddress(arg0 types0.Context, arg1 types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccountWithAddress", arg0, arg1)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// NewAccountWithAddress indicates an expected call of NewAccountWithAddress.
func (mr *MockAccountKeeperIMockRecorder) NewAccountWithAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccountWithAddress", reflect.TypeOf((*MockAccountKeeperI)(nil).NewAccountWithAddress), arg0, arg1)
}

// RemoveAccount mocks base method.
func (m *MockAccountKeeperI) RemoveAccount(arg0 types0.Context, arg1 types1.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAccount", arg0, arg1)
}

// RemoveAccount indicates an expected call of RemoveAccount.
func (mr *MockAccountKeeperIMockRecorder) RemoveAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccount", reflect.TypeOf((*MockAccountKeeperI)(nil).RemoveAccount), arg0, arg1)
}

// SetAccount mocks base method.
func (m *MockAccountKeeperI) SetAccount(arg0 types0.Context, arg1 types1.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", arg0, arg1)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperIMockRecorder) SetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeperI)(nil).SetAccount), arg0, arg1)
}
