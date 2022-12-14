// Code generated by MockGen. DO NOT EDIT.
// Source: internal/component/ratelimiter/contract.go

// Package mockcomponent is a generated GoMock package.
package mockcomponent

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRateLimiter is a mock of RateLimiter interface.
type MockRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimiterMockRecorder
}

// MockRateLimiterMockRecorder is the mock recorder for MockRateLimiter.
type MockRateLimiterMockRecorder struct {
	mock *MockRateLimiter
}

// NewMockRateLimiter creates a new mock instance.
func NewMockRateLimiter(ctrl *gomock.Controller) *MockRateLimiter {
	mock := &MockRateLimiter{ctrl: ctrl}
	mock.recorder = &MockRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimiter) EXPECT() *MockRateLimiterMockRecorder {
	return m.recorder
}

// Allow mocks base method.
func (m *MockRateLimiter) Allow(maxAttempt int, attemptWaitDuration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow", maxAttempt, attemptWaitDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Allow indicates an expected call of Allow.
func (mr *MockRateLimiterMockRecorder) Allow(maxAttempt, attemptWaitDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockRateLimiter)(nil).Allow), maxAttempt, attemptWaitDuration)
}

// Finish mocks base method.
func (m *MockRateLimiter) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockRateLimiterMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockRateLimiter)(nil).Finish))
}
