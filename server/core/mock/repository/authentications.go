// Code generated by MockGen. DO NOT EDIT.
// Source: authentications.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthentications is a mock of Authentications interface.
type MockAuthentications struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationsMockRecorder
}

// MockAuthenticationsMockRecorder is the mock recorder for MockAuthentications.
type MockAuthenticationsMockRecorder struct {
	mock *MockAuthentications
}

// NewMockAuthentications creates a new mock instance.
func NewMockAuthentications(ctrl *gomock.Controller) *MockAuthentications {
	mock := &MockAuthentications{ctrl: ctrl}
	mock.recorder = &MockAuthenticationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthentications) EXPECT() *MockAuthenticationsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAuthentications) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockAuthenticationsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAuthentications)(nil).Close))
}

// Create mocks base method.
func (m *MockAuthentications) Create(identify, token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", identify, token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthenticationsMockRecorder) Create(identify, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthentications)(nil).Create), identify, token)
}

// FindToken mocks base method.
func (m *MockAuthentications) FindToken(identify string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindToken", identify)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindToken indicates an expected call of FindToken.
func (mr *MockAuthenticationsMockRecorder) FindToken(identify interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindToken", reflect.TypeOf((*MockAuthentications)(nil).FindToken), identify)
}

// StoreOAuth2Info mocks base method.
func (m *MockAuthentications) StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreOAuth2Info", identify, accessToken, refreshToken, expiry)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreOAuth2Info indicates an expected call of StoreOAuth2Info.
func (mr *MockAuthenticationsMockRecorder) StoreOAuth2Info(identify, accessToken, refreshToken, expiry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreOAuth2Info", reflect.TypeOf((*MockAuthentications)(nil).StoreOAuth2Info), identify, accessToken, refreshToken, expiry)
}

// UpdateToken mocks base method.
func (m *MockAuthentications) UpdateToken(id int, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", id, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockAuthenticationsMockRecorder) UpdateToken(id, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockAuthentications)(nil).UpdateToken), id, token)
}
