// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/repository/organizations.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yukitaka/longlong/internal/entity"
)

// MockOrganizations is a mock of Organizations interface.
type MockOrganizations struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsMockRecorder
}

// MockOrganizationsMockRecorder is the mock recorder for MockOrganizations.
type MockOrganizationsMockRecorder struct {
	mock *MockOrganizations
}

// NewMockOrganizations creates a new mock instance.
func NewMockOrganizations(ctrl *gomock.Controller) *MockOrganizations {
	mock := &MockOrganizations{ctrl: ctrl}
	mock.recorder = &MockOrganizationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizations) EXPECT() *MockOrganizationsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganizations) Create(name string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(int)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationsMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizations)(nil).Create), name)
}

// Find mocks base method.
func (m *MockOrganizations) Find(arg0 int) (*entity.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*entity.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockOrganizationsMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockOrganizations)(nil).Find), arg0)
}