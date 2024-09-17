package server

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config "github.com/Stern-Ritter/gophertask/internal/config/server"
	er "github.com/Stern-Ritter/gophertask/internal/errors"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name                 string
		req                  *pb.SignUpRequestV1
		expectedResp         *pb.SignUpResponseV1
		authServiceSignUpErr error
		expectedErr          error
	}{
		{
			name: "should return conflict error if user already exists",
			req: &pb.SignUpRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignUpErr: er.NewConflictError("user already exists", nil),
			expectedErr:          status.Error(codes.AlreadyExists, "user already exists"),
		},
		{
			name: "should sign up successfully",
			req: &pb.SignUpRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignUpErr: nil,
			expectedResp: &pb.SignUpResponseV1{
				Token: "token",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockAuthService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("token", tt.authServiceSignUpErr).Times(1)

			resp, err := s.SignUp(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name                 string
		req                  *pb.SignInRequestV1
		expectedResp         *pb.SignInResponseV1
		authServiceSignInErr error
		expectedErr          error
	}{
		{
			name: "should return unauthorized error if credentials are incorrect",
			req: &pb.SignInRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignInErr: er.NewUnauthorizedError("unauthorized", nil),
			expectedErr:          status.Error(codes.Unauthenticated, "unauthorized"),
		},
		{
			name: "should sign in successfully",
			req: &pb.SignInRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignInErr: nil,
			expectedResp: &pb.SignInResponseV1{
				Token: "token",
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuthService.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("token", tc.authServiceSignInErr).Times(1)

			resp, err := s.SignIn(context.Background(), tc.req)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, tc.expectedErr, err, "should return error: %s, got: %s", tc.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tc.expectedResp, resp, "should return response: %v, got: %v", tc.expectedResp, resp)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddTaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateTask     bool
		createTaskErr      error
		expectedResp       *pb.AddTaskResponseV1
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.AddTaskRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if task service returns error",
			req:                &pb.AddTaskRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			callCreateTask:     true,
			createTaskErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should add task successfully",
			req:                &pb.AddTaskRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			callCreateTask:     true,
			expectedResp:       &pb.AddTaskResponseV1{},
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callCreateTask {
				mockTaskService.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(tt.createTaskErr)
			}

			resp, err := s.AddTask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, tt.expectedResp, resp, "expected response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.UpdateTaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetTaskByID    bool
		getTaskByIDErr     error
		callUpdateTask     bool
		updateTaskErr      error
		expectedResp       *pb.UpdateTaskResponseV1
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.UpdateTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if task is not found",
			req:                &pb.UpdateTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			getTaskByIDErr:     er.NewNotFoundError("task not found", nil),
			expectedErr:        status.Error(codes.NotFound, "task not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if update task service returns error",
			req:                &pb.UpdateTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			callUpdateTask:     true,
			updateTaskErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should update task successfully",
			req:                &pb.UpdateTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			callUpdateTask:     true,
			expectedResp:       &pb.UpdateTaskResponseV1{},
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetTaskByID {
				mockTaskService.EXPECT().GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Task{}, tt.getTaskByIDErr)
			}
			if tt.callUpdateTask {
				mockTaskService.EXPECT().UpdateTask(gomock.Any(), gomock.Any()).Return(tt.updateTaskErr)
			}

			resp, err := s.UpdateTask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, tt.expectedResp, resp, "expected response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteTaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteTask     bool
		deleteTaskErr      error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.DeleteTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if task is not found",
			req:                &pb.DeleteTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteTask:     true,
			deleteTaskErr:      er.NewNotFoundError("task not found", nil),
			expectedErr:        status.Error(codes.NotFound, "task not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return permission denied error if user does not have permission",
			req:                &pb.DeleteTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteTask:     true,
			deleteTaskErr:      er.NewForbiddenError("forbidden access", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "forbidden access"),
			expectedCode:       codes.PermissionDenied,
		},
		{
			name:               "should return internal error if task service returns unknown error",
			req:                &pb.DeleteTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteTask:     true,
			deleteTaskErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should delete task successfully",
			req:                &pb.DeleteTaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteTask:     true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteTask {
				mockTaskService.EXPECT().DeleteTask(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteTaskErr)
			}

			resp, err := s.DeleteTask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, &pb.DeleteTaskResponseV1{}, resp, "expected response: %v, got: %v", resp)
			}
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetTaskByIDRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetTaskByID    bool
		getTaskByIDErr     error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetTaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if task is not found",
			req:                &pb.GetTaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			getTaskByIDErr:     er.NewNotFoundError("task not found", nil),
			expectedErr:        status.Error(codes.NotFound, "task not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return permission denied error if user does not have permission",
			req:                &pb.GetTaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			getTaskByIDErr:     er.NewForbiddenError("forbidden access", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "forbidden access"),
			expectedCode:       codes.PermissionDenied,
		},
		{
			name:               "should return internal error if task service returns unknown error",
			req:                &pb.GetTaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			getTaskByIDErr:     errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should get task successfully",
			req:                &pb.GetTaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetTaskByID:    true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetTaskByID {
				mockTaskService.EXPECT().GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.Task{}, tt.getTaskByIDErr)
			}

			_, err := s.GetTaskByID(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestGetTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetTasksRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetAllTasks    bool
		getAllTasksErr     error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetTasksRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if task service returns error",
			req:                &pb.GetTasksRequestV1{},
			callGetCurrentUser: true,
			callGetAllTasks:    true,
			getAllTasksErr:     errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should return tasks successfully",
			req:                &pb.GetTasksRequestV1{},
			callGetCurrentUser: true,
			callGetAllTasks:    true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetAllTasks {
				mockTaskService.EXPECT().GetAllTasks(gomock.Any(), gomock.Any()).Return([]model.Task{}, tt.getAllTasksErr)
			}

			_, err := s.GetTasks(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestAddSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddSubtaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateSubtask  bool
		createSubtaskErr   error
		expectedResp       *pb.AddSubtaskResponseV1
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.AddSubtaskRequestV1{EpicId: "1", Name: "TestSubtask"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if subtask service returns error",
			req:                &pb.AddSubtaskRequestV1{EpicId: "1", Name: "TestSubtask"},
			callGetCurrentUser: true,
			callCreateSubtask:  true,
			createSubtaskErr:   errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should add subtask successfully",
			req:                &pb.AddSubtaskRequestV1{EpicId: "1", Name: "TestSubtask"},
			callGetCurrentUser: true,
			callCreateSubtask:  true,
			expectedResp:       &pb.AddSubtaskResponseV1{},
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callCreateSubtask {
				mockSubtaskService.EXPECT().CreateSubtask(gomock.Any(), gomock.Any()).Return(tt.createSubtaskErr)
			}

			resp, err := s.AddSubtask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, tt.expectedResp, resp, "expected response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestUpdateSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.UpdateSubtaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetSubtaskByID bool
		getSubtaskByIDErr  error
		callUpdateSubtask  bool
		updateSubtaskErr   error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.UpdateSubtaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if subtask is not found",
			req:                &pb.UpdateSubtaskRequestV1{EpicId: "1", Id: "subtask1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			getSubtaskByIDErr:  er.NewNotFoundError("subtask not found", nil),
			expectedErr:        status.Error(codes.NotFound, "subtask not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if update subtask service returns error",
			req:                &pb.UpdateSubtaskRequestV1{EpicId: "1", Id: "1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			callUpdateSubtask:  true,
			updateSubtaskErr:   errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should update subtask successfully",
			req:                &pb.UpdateSubtaskRequestV1{EpicId: "1", Id: "1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			callUpdateSubtask:  true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetSubtaskByID {
				mockSubtaskService.EXPECT().GetSubtaskByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Subtask{}, tt.getSubtaskByIDErr)
			}
			if tt.callUpdateSubtask {
				mockSubtaskService.EXPECT().UpdateSubtask(gomock.Any(), gomock.Any()).Return(tt.updateSubtaskErr)
			}

			_, err := s.UpdateSubtask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestDeleteSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteSubtaskRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteSubtask  bool
		deleteSubtaskErr   error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.DeleteSubtaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if subtask is not found",
			req:                &pb.DeleteSubtaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteSubtask:  true,
			deleteSubtaskErr:   er.NewNotFoundError("subtask not found", nil),
			expectedErr:        status.Error(codes.NotFound, "subtask not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if subtask service returns unknown error",
			req:                &pb.DeleteSubtaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteSubtask:  true,
			deleteSubtaskErr:   errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should delete subtask successfully",
			req:                &pb.DeleteSubtaskRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteSubtask:  true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteSubtask {
				mockSubtaskService.EXPECT().DeleteSubtask(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteSubtaskErr)
			}

			resp, err := s.DeleteSubtask(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, &pb.DeleteSubtaskResponseV1{}, resp, "expected response: %v, got: %v", resp)
			}
		})
	}
}

func TestGetSubtaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetSubtaskByIDRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetSubtaskByID bool
		getSubtaskByIDErr  error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetSubtaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if subtask is not found",
			req:                &pb.GetSubtaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			getSubtaskByIDErr:  er.NewNotFoundError("subtask not found", nil),
			expectedErr:        status.Error(codes.NotFound, "subtask not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return permission denied error if user does not have permission",
			req:                &pb.GetSubtaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			getSubtaskByIDErr:  er.NewForbiddenError("forbidden access", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "forbidden access"),
			expectedCode:       codes.PermissionDenied,
		},
		{
			name:               "should get subtask successfully",
			req:                &pb.GetSubtaskByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetSubtaskByID: true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetSubtaskByID {
				mockSubtaskService.EXPECT().GetSubtaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.Subtask{}, tt.getSubtaskByIDErr)
			}

			_, err := s.GetSubtaskByID(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestGetSubtasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetSubtasksRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetAllSubtasks bool
		getAllSubtasksErr  error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetSubtasksRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if subtask service returns error",
			req:                &pb.GetSubtasksRequestV1{},
			callGetCurrentUser: true,
			callGetAllSubtasks: true,
			getAllSubtasksErr:  errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should return subtasks successfully",
			req:                &pb.GetSubtasksRequestV1{},
			callGetCurrentUser: true,
			callGetAllSubtasks: true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetAllSubtasks {
				mockSubtaskService.EXPECT().GetAllSubtasks(gomock.Any(), gomock.Any()).Return([]model.Subtask{}, tt.getAllSubtasksErr)
			}

			_, err := s.GetSubtasks(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestAddEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddEpicRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateEpic     bool
		createEpicErr      error
		expectedResp       *pb.AddEpicResponseV1
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.AddEpicRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if epic service fails",
			req:                &pb.AddEpicRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			callCreateEpic:     true,
			createEpicErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should add epic successfully",
			req:                &pb.AddEpicRequestV1{Name: "Test"},
			callGetCurrentUser: true,
			callCreateEpic:     true,
			expectedResp:       &pb.AddEpicResponseV1{},
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callCreateEpic {
				mockEpicService.EXPECT().CreateEpic(gomock.Any(), gomock.Any()).Return(tt.createEpicErr)
			}

			resp, err := s.AddEpic(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, tt.expectedResp, resp, "expected response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestUpdateEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.UpdateEpicRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetEpicByID    bool
		getEpicByIDErr     error
		callUpdateEpic     bool
		updateEpicErr      error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.UpdateEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if epic is not found",
			req:                &pb.UpdateEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			getEpicByIDErr:     er.NewNotFoundError("epic not found", nil),
			expectedErr:        status.Error(codes.NotFound, "epic not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if update epic service fails",
			req:                &pb.UpdateEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			callUpdateEpic:     true,
			updateEpicErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should update epic successfully",
			req:                &pb.UpdateEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			callUpdateEpic:     true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetEpicByID {
				mockEpicService.EXPECT().GetEpicByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Epic{}, tt.getEpicByIDErr)
			}
			if tt.callUpdateEpic {
				mockEpicService.EXPECT().UpdateEpic(gomock.Any(), gomock.Any()).Return(tt.updateEpicErr)
			}

			_, err := s.UpdateEpic(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestDeleteEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteEpicRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteEpic     bool
		deleteEpicErr      error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.DeleteEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if epic is not found",
			req:                &pb.DeleteEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteEpic:     true,
			deleteEpicErr:      er.NewNotFoundError("epic not found", nil),
			expectedErr:        status.Error(codes.NotFound, "epic not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if delete epic service fails",
			req:                &pb.DeleteEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteEpic:     true,
			deleteEpicErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should delete epic successfully",
			req:                &pb.DeleteEpicRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callDeleteEpic:     true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteEpic {
				mockEpicService.EXPECT().DeleteEpic(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteEpicErr)
			}

			resp, err := s.DeleteEpic(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
				assert.Equal(t, &pb.DeleteEpicResponseV1{}, resp, "expected response: %v, got: %v", resp)
			}
		})
	}
}

func TestGetEpicByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetEpicByIDRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetEpicByID    bool
		getEpicByIDErr     error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetEpicByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return not found error if epic is not found",
			req:                &pb.GetEpicByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			getEpicByIDErr:     er.NewNotFoundError("epic not found", nil),
			expectedErr:        status.Error(codes.NotFound, "epic not found"),
			expectedCode:       codes.NotFound,
		},
		{
			name:               "should return internal error if get epic service fails",
			req:                &pb.GetEpicByIDRequestV1{Id: "1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			getEpicByIDErr:     errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should get epic successfully",
			req:                &pb.GetEpicByIDRequestV1{Id: "epic1"},
			callGetCurrentUser: true,
			callGetEpicByID:    true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetEpicByID {
				mockEpicService.EXPECT().GetEpicByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.Epic{}, tt.getEpicByIDErr)
			}

			_, err := s.GetEpicByID(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}

func TestGetEpics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockTaskService := NewMockTaskService(ctrl)
	mockSubtaskService := NewMockSubtaskService(ctrl)
	mockEpicService := NewMockEpicService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockTaskService, mockSubtaskService, mockEpicService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetEpicsRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetAllEpics    bool
		getAllEpicsErr     error
		expectedErr        error
		expectedCode       codes.Code
	}{
		{
			name:               "should return unauthenticated error if user service fails",
			req:                &pb.GetEpicsRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
			expectedCode:       codes.Unauthenticated,
		},
		{
			name:               "should return internal error if get all epics service fails",
			req:                &pb.GetEpicsRequestV1{},
			callGetCurrentUser: true,
			callGetAllEpics:    true,
			getAllEpicsErr:     errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
			expectedCode:       codes.Internal,
		},
		{
			name:               "should return epics successfully",
			req:                &pb.GetEpicsRequestV1{},
			callGetCurrentUser: true,
			callGetAllEpics:    true,
			expectedErr:        nil,
			expectedCode:       codes.OK,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetAllEpics {
				mockEpicService.EXPECT().GetAllEpics(gomock.Any(), gomock.Any()).Return([]model.Epic{}, tt.getAllEpicsErr)
			}

			_, err := s.GetEpics(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "expected error: %v, got: %v", tt.expectedErr, err)
				st, _ := status.FromError(err)
				assert.Equal(t, tt.expectedCode, st.Code(), "expected status code: %v, got: %v", tt.expectedCode, st.Code())
			} else {
				assert.NoError(t, err, "expected no error, but got: %v", err)
			}
		})
	}
}
