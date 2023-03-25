// Code generated by MockGen. DO NOT EDIT.
// Source: organization_belongings.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yukitaka/longlong/internal/domain/entity"
)

// MockOrganizationBelongings is a mock of OrganizationBelongings interface.
type MockOrganizationBelongings struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationBelongingsMockRecorder
}

// MockOrganizationBelongingsMockRecorder is the mock recorder for MockOrganizationBelongings.
type MockOrganizationBelongingsMockRecorder struct {
	mock *MockOrganizationBelongings
}

// NewMockOrganizationBelongings creates a new mock instance.
func NewMockOrganizationBelongings(ctrl *gomock.Controller) *MockOrganizationBelongings {
	mock := &MockOrganizationBelongings{ctrl: ctrl}
	mock.recorder = &MockOrganizationBelongingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationBelongings) EXPECT() *MockOrganizationBelongingsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockOrganizationBelongings) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockOrganizationBelongingsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOrganizationBelongings)(nil).Close))
}

// Entry mocks base method.
func (m *MockOrganizationBelongings) Entry(organizationId, userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entry", organizationId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Entry indicates an expected call of Entry.
func (mr *MockOrganizationBelongingsMockRecorder) Entry(organizationId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entry", reflect.TypeOf((*MockOrganizationBelongings)(nil).Entry), organizationId, userId)
}

// Leave mocks base method.
func (m *MockOrganizationBelongings) Leave(organizationId, userId int64, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", organizationId, userId, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave.
func (mr *MockOrganizationBelongingsMockRecorder) Leave(organizationId, userId, reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockOrganizationBelongings)(nil).Leave), organizationId, userId, reason)
}

// Members mocks base method.
func (m *MockOrganizationBelongings) Members(organizationId int64) (*[]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members", organizationId)
	ret0, _ := ret[0].(*[]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Members indicates an expected call of Members.
func (mr *MockOrganizationBelongingsMockRecorder) Members(organizationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockOrganizationBelongings)(nil).Members), organizationId)
}
