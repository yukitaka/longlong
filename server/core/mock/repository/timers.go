// Code generated by MockGen. DO NOT EDIT.
// Source: timers.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
)

// MockTimers is a mock of Timers interface.
type MockTimers struct {
	ctrl     *gomock.Controller
	recorder *MockTimersMockRecorder
}

// MockTimersMockRecorder is the mock recorder for MockTimers.
type MockTimersMockRecorder struct {
	mock *MockTimers
}

// NewMockTimers creates a new mock instance.
func NewMockTimers(ctrl *gomock.Controller) *MockTimers {
	mock := &MockTimers{ctrl: ctrl}
	mock.recorder = &MockTimersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimers) EXPECT() *MockTimersMockRecorder {
	return m.recorder
}
