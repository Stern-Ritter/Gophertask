package client

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"

	"github.com/Stern-Ritter/gophertask/internal/model"
)

type TaskService interface {
	CreateTask(name string, description string, duration time.Duration, startedAt time.Time) error
	UpdateTask(id string, name string, description string, status model.TaskStatus, duration time.Duration, startedAt time.Time) error
	DeleteTask(taskID string) error
	GetTaskByID(taskID string) (*pb.TaskV1, error)
	GetAllTasks() ([]*pb.TaskV1, error)
}

type TaskServiceImpl struct {
	taskClient pb.TaskServiceV1Client
}

func NewTaskService(taskClient pb.TaskServiceV1Client) TaskService {
	return &TaskServiceImpl{
		taskClient: taskClient,
	}
}

func (s *TaskServiceImpl) CreateTask(name string, description string, duration time.Duration, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AddTaskRequestV1{
		Name:        name,
		Description: description,
		Duration:    durationpb.New(duration),
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.taskClient.AddTask(ctx, req)
	return err
}

func (s *TaskServiceImpl) UpdateTask(id string, name string, description string, status model.TaskStatus,
	duration time.Duration, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateTaskRequestV1{
		Id:          id,
		Name:        name,
		Description: description,
		Status:      model.MapTaskStatusToMessageTaskStatus(status),
		Duration:    durationpb.New(duration),
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.taskClient.UpdateTask(ctx, req)
	return err
}

func (s *TaskServiceImpl) DeleteTask(taskID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteTaskRequestV1{
		Id: taskID,
	}

	_, err := s.taskClient.DeleteTask(ctx, req)
	return err
}

func (s *TaskServiceImpl) GetTaskByID(taskID string) (*pb.TaskV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetTaskByIDRequestV1{
		Id: taskID,
	}

	resp, err := s.taskClient.GetTaskByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetTask(), nil
}

func (s *TaskServiceImpl) GetAllTasks() ([]*pb.TaskV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetTasksRequestV1{}
	resp, err := s.taskClient.GetTasks(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetTasks(), nil
}
