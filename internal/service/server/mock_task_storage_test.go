// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/storage/server/task_storage.go
//
// Generated by this command:
//
//	mockgen -source=./internal/storage/server/task_storage.go -destination ./internal/service/server/mock_task_storage_test.go -package server
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	model "github.com/Stern-Ritter/gophertask/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskStorage is a mock of TaskStorage interface.
type MockTaskStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTaskStorageMockRecorder
}

// MockTaskStorageMockRecorder is the mock recorder for MockTaskStorage.
type MockTaskStorageMockRecorder struct {
	mock *MockTaskStorage
}

// NewMockTaskStorage creates a new mock instance.
func NewMockTaskStorage(ctrl *gomock.Controller) *MockTaskStorage {
	mock := &MockTaskStorage{ctrl: ctrl}
	mock.recorder = &MockTaskStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskStorage) EXPECT() *MockTaskStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskStorage) Create(ctx context.Context, task model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTaskStorageMockRecorder) Create(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskStorage)(nil).Create), ctx, task)
}

// Delete mocks base method.
func (m *MockTaskStorage) Delete(ctx context.Context, taskID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskStorageMockRecorder) Delete(ctx, taskID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskStorage)(nil).Delete), ctx, taskID)
}

// GetAll mocks base method.
func (m *MockTaskStorage) GetAll(ctx context.Context, userID string) ([]model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, userID)
	ret0, _ := ret[0].([]model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTaskStorageMockRecorder) GetAll(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTaskStorage)(nil).GetAll), ctx, userID)
}

// GetByID mocks base method.
func (m *MockTaskStorage) GetByID(ctx context.Context, taskID string) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, taskID)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTaskStorageMockRecorder) GetByID(ctx, taskID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTaskStorage)(nil).GetByID), ctx, taskID)
}

// Update mocks base method.
func (m *MockTaskStorage) Update(ctx context.Context, task model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTaskStorageMockRecorder) Update(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskStorage)(nil).Update), ctx, task)
}
