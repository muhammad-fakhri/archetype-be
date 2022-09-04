// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/contract.go

// Package mockrepository is a generated GoMock package.
package mockrepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	constant "github.com/muhammad-fakhri/archetype-be/internal/constant"
	model "github.com/muhammad-fakhri/archetype-be/internal/model"
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

// GetSystemConfigAll mocks base method.
func (m *MockRepository) GetSystemConfigAll(ctx context.Context) ([]*model.SystemConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemConfigAll", ctx)
	ret0, _ := ret[0].([]*model.SystemConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemConfigAll indicates an expected call of GetSystemConfigAll.
func (mr *MockRepositoryMockRecorder) GetSystemConfigAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemConfigAll", reflect.TypeOf((*MockRepository)(nil).GetSystemConfigAll), ctx)
}

// GetSystemConfigByName mocks base method.
func (m *MockRepository) GetSystemConfigByName(ctx context.Context, name constant.SystemConfig, config interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemConfigByName", ctx, name, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSystemConfigByName indicates an expected call of GetSystemConfigByName.
func (mr *MockRepositoryMockRecorder) GetSystemConfigByName(ctx, name, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemConfigByName", reflect.TypeOf((*MockRepository)(nil).GetSystemConfigByName), ctx, name, config)
}

// PublishSystemConfig mocks base method.
func (m *MockRepository) PublishSystemConfig(ctx context.Context, data map[constant.SystemConfig]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishSystemConfig", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishSystemConfig indicates an expected call of PublishSystemConfig.
func (mr *MockRepositoryMockRecorder) PublishSystemConfig(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishSystemConfig", reflect.TypeOf((*MockRepository)(nil).PublishSystemConfig), ctx, data)
}

// SendSystemConfigReport mocks base method.
func (m *MockRepository) SendSystemConfigReport(ctx context.Context, details string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSystemConfigReport", ctx, details)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSystemConfigReport indicates an expected call of SendSystemConfigReport.
func (mr *MockRepositoryMockRecorder) SendSystemConfigReport(ctx, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSystemConfigReport", reflect.TypeOf((*MockRepository)(nil).SendSystemConfigReport), ctx, details)
}

// UpdateSystemConfig mocks base method.
func (m *MockRepository) UpdateSystemConfig(ctx context.Context, name constant.SystemConfig, config interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSystemConfig", ctx, name, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSystemConfig indicates an expected call of UpdateSystemConfig.
func (mr *MockRepositoryMockRecorder) UpdateSystemConfig(ctx, name, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSystemConfig", reflect.TypeOf((*MockRepository)(nil).UpdateSystemConfig), ctx, name, config)
}
