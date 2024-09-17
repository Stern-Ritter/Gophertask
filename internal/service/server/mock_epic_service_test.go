// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/server/epic_service.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/server/epic_service.go -destination ./internal/service/server/mock_epic_service_test.go -package server
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	model "github.com/Stern-Ritter/gophertask/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockEpicService is a mock of EpicService interface.
type MockEpicService struct {
	ctrl     *gomock.Controller
	recorder *MockEpicServiceMockRecorder
}

// MockEpicServiceMockRecorder is the mock recorder for MockEpicService.
type MockEpicServiceMockRecorder struct {
	mock *MockEpicService
}

// NewMockEpicService creates a new mock instance.
func NewMockEpicService(ctrl *gomock.Controller) *MockEpicService {
	mock := &MockEpicService{ctrl: ctrl}
	mock.recorder = &MockEpicServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpicService) EXPECT() *MockEpicServiceMockRecorder {
	return m.recorder
}

// CreateEpic mocks base method.
func (m *MockEpicService) CreateEpic(ctx context.Context, epic model.Epic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEpic", ctx, epic)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEpic indicates an expected call of CreateEpic.
func (mr *MockEpicServiceMockRecorder) CreateEpic(ctx, epic any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEpic", reflect.TypeOf((*MockEpicService)(nil).CreateEpic), ctx, epic)
}

// DeleteEpic mocks base method.
func (m *MockEpicService) DeleteEpic(ctx context.Context, userID, epicID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEpic", ctx, userID, epicID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEpic indicates an expected call of DeleteEpic.
func (mr *MockEpicServiceMockRecorder) DeleteEpic(ctx, userID, epicID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEpic", reflect.TypeOf((*MockEpicService)(nil).DeleteEpic), ctx, userID, epicID)
}

// GetAllEpics mocks base method.
func (m *MockEpicService) GetAllEpics(ctx context.Context, userID string) ([]model.Epic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEpics", ctx, userID)
	ret0, _ := ret[0].([]model.Epic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEpics indicates an expected call of GetAllEpics.
func (mr *MockEpicServiceMockRecorder) GetAllEpics(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEpics", reflect.TypeOf((*MockEpicService)(nil).GetAllEpics), ctx, userID)
}

// GetEpicByID mocks base method.
func (m *MockEpicService) GetEpicByID(ctx context.Context, userID, epicID string) (model.Epic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpicByID", ctx, userID, epicID)
	ret0, _ := ret[0].(model.Epic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEpicByID indicates an expected call of GetEpicByID.
func (mr *MockEpicServiceMockRecorder) GetEpicByID(ctx, userID, epicID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpicByID", reflect.TypeOf((*MockEpicService)(nil).GetEpicByID), ctx, userID, epicID)
}

// UpdateEpic mocks base method.
func (m *MockEpicService) UpdateEpic(ctx context.Context, epic model.Epic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEpic", ctx, epic)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEpic indicates an expected call of UpdateEpic.
func (mr *MockEpicServiceMockRecorder) UpdateEpic(ctx, epic any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEpic", reflect.TypeOf((*MockEpicService)(nil).UpdateEpic), ctx, epic)
}
