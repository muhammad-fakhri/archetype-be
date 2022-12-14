// Code generated by MockGen. DO NOT EDIT.
// Source: internal/component/localcache/contract.go

// Package mockcomponent is a generated GoMock package.
package mockcomponent

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocalCache is a mock of LocalCache interface.
type MockLocalCache struct {
	ctrl     *gomock.Controller
	recorder *MockLocalCacheMockRecorder
}

// MockLocalCacheMockRecorder is the mock recorder for MockLocalCache.
type MockLocalCacheMockRecorder struct {
	mock *MockLocalCache
}

// NewMockLocalCache creates a new mock instance.
func NewMockLocalCache(ctrl *gomock.Controller) *MockLocalCache {
	mock := &MockLocalCache{ctrl: ctrl}
	mock.recorder = &MockLocalCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalCache) EXPECT() *MockLocalCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLocalCache) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLocalCacheMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLocalCache)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockLocalCache) Get(key string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLocalCacheMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLocalCache)(nil).Get), key)
}

// Set mocks base method.
func (m *MockLocalCache) Set(key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockLocalCacheMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLocalCache)(nil).Set), key, value)
}
