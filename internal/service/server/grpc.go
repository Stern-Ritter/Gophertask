package server

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	er "github.com/Stern-Ritter/gophertask/internal/errors"
	"github.com/Stern-Ritter/gophertask/internal/model"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequestV1) (*pb.SignUpResponseV1, error) {
	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.MessageToSignUpRequest(in)
	token, err := s.AuthService.SignUp(ctx, req)
	if err != nil {
		var conflictErr er.ConflictError
		if errors.As(err, &conflictErr) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	resp := pb.SignUpResponseV1{
		Token: token,
	}

	return &resp, nil
}
func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequestV1) (*pb.SignInResponseV1, error) {
	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.MessageToSignInRequest(in)
	token, err := s.AuthService.SignIn(ctx, req)

	if err != nil {
		var unauthorizedErr er.UnauthorizedError
		if errors.As(err, &unauthorizedErr) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, "internal server error")
	}

	resp := pb.SignInResponseV1{
		Token: token,
	}

	return &resp, nil
}

func (s *Server) AddTask(ctx context.Context, in *pb.AddTaskRequestV1) (*pb.AddTaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	task := model.AddTaskRequestToTask(in)
	task.SetUserID(user.ID)

	err = s.TaskService.CreateTask(ctx, task)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddTaskResponseV1{}, nil
}

func (s *Server) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequestV1) (*pb.UpdateTaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	taskID := in.Id
	t, err := s.TaskService.GetTaskByID(ctx, user.ID, taskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	updatedTask := model.UpdateTaskRequestToTask(in, t)
	err = s.TaskService.UpdateTask(ctx, updatedTask)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.UpdateTaskResponseV1{}, nil
}

func (s *Server) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequestV1) (*pb.DeleteTaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	taskID := in.Id
	err = s.TaskService.DeleteTask(ctx, user.ID, taskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteTaskResponseV1{}, nil
}

func (s *Server) GetTaskByID(ctx context.Context, in *pb.GetTaskByIDRequestV1) (*pb.GetTaskByIDResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	taskID := in.Id
	t, err := s.TaskService.GetTaskByID(ctx, user.ID, taskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	task := model.TaskToTaskMessage(t)

	return &pb.GetTaskByIDResponseV1{Task: task}, nil
}

func (s *Server) GetTasks(ctx context.Context, in *pb.GetTasksRequestV1) (*pb.GetTasksResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	t, err := s.TaskService.GetAllTasks(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	tasks := model.TasksToRepeatedTaskMessage(t)
	resp := pb.GetTasksResponseV1{
		Tasks: tasks,
	}

	return &resp, nil
}

func (s *Server) AddSubtask(ctx context.Context, in *pb.AddSubtaskRequestV1) (*pb.AddSubtaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	subtask := model.AddSubtaskRequestToSubtask(in)
	subtask.SetUserID(user.ID)

	err = s.SubtaskService.CreateSubtask(ctx, subtask)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddSubtaskResponseV1{}, nil
}

func (s *Server) UpdateSubtask(ctx context.Context, in *pb.UpdateSubtaskRequestV1) (*pb.UpdateSubtaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	subtaskID := in.Id
	st, err := s.SubtaskService.GetSubtaskByID(ctx, user.ID, subtaskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	updatedSubtask := model.UpdateSubtaskRequestToSubtask(in, st)
	err = s.SubtaskService.UpdateSubtask(ctx, updatedSubtask)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.UpdateSubtaskResponseV1{}, nil
}

func (s *Server) DeleteSubtask(ctx context.Context, in *pb.DeleteSubtaskRequestV1) (*pb.DeleteSubtaskResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	subtaskID := in.Id
	err = s.SubtaskService.DeleteSubtask(ctx, user.ID, subtaskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteSubtaskResponseV1{}, nil
}

func (s *Server) GetSubtaskByID(ctx context.Context, in *pb.GetSubtaskByIDRequestV1) (*pb.GetSubtaskByIDResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	subtaskID := in.Id
	st, err := s.SubtaskService.GetSubtaskByID(ctx, user.ID, subtaskID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	subtask := model.SubtaskToSubtaskMessage(st)

	return &pb.GetSubtaskByIDResponseV1{Subtask: subtask}, nil
}

func (s *Server) GetSubtasks(ctx context.Context, in *pb.GetSubtasksRequestV1) (*pb.GetSubtasksResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	st, err := s.SubtaskService.GetAllSubtasks(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	subtasks := model.SubtaskToRepeatedSubtaskMessage(st)
	resp := pb.GetSubtasksResponseV1{
		Subtasks: subtasks,
	}

	return &resp, nil
}

func (s *Server) AddEpic(ctx context.Context, in *pb.AddEpicRequestV1) (*pb.AddEpicResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	epic := model.AddEpicRequestToEpic(in)
	epic.SetUserID(user.ID)

	err = s.EpicService.CreateEpic(ctx, epic)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddEpicResponseV1{}, nil
}

func (s *Server) UpdateEpic(ctx context.Context, in *pb.UpdateEpicRequestV1) (*pb.UpdateEpicResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	epicID := in.Id
	e, err := s.EpicService.GetEpicByID(ctx, user.ID, epicID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	updatedEpic := model.UpdateEpicRequestToEpic(in, e)
	err = s.EpicService.UpdateEpic(ctx, updatedEpic)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.UpdateEpicResponseV1{}, nil
}

func (s *Server) DeleteEpic(ctx context.Context, in *pb.DeleteEpicRequestV1) (*pb.DeleteEpicResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	epicID := in.Id
	err = s.EpicService.DeleteEpic(ctx, user.ID, epicID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteEpicResponseV1{}, nil
}

func (s *Server) GetEpicByID(ctx context.Context, in *pb.GetEpicByIDRequestV1) (*pb.GetEpicByIDResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	epicID := in.Id
	e, err := s.EpicService.GetEpicByID(ctx, user.ID, epicID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	epic := model.EpicToEpicMessage(e)

	return &pb.GetEpicByIDResponseV1{Epic: epic}, nil
}

func (s *Server) GetEpics(ctx context.Context, in *pb.GetEpicsRequestV1) (*pb.GetEpicsResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	e, err := s.EpicService.GetAllEpics(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	epics := model.EpicsToRepeatedEpicMessage(e)
	resp := pb.GetEpicsResponseV1{
		Epics: epics,
	}

	return &resp, nil
}
