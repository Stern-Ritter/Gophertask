package client

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"

	"github.com/Stern-Ritter/gophertask/internal/model"
)

type SubtaskService interface {
	CreateSubtask(epicID string, name string, description string, duration time.Duration, startedAt time.Time) error
	UpdateSubtask(id string, epicID string, name string, description string, status model.TaskStatus, duration time.Duration, startedAt time.Time) error
	DeleteSubtask(subtaskID string) error
	GetSubtaskByID(subtaskID string) (*pb.SubtaskV1, error)
	GetAllSubtasks() ([]*pb.SubtaskV1, error)
}

type SubtaskServiceImpl struct {
	subtaskClient pb.SubtaskServiceV1Client
}

func NewSubtaskService(subtaskClient pb.SubtaskServiceV1Client) SubtaskService {
	return &SubtaskServiceImpl{
		subtaskClient: subtaskClient,
	}
}

func (s *SubtaskServiceImpl) CreateSubtask(epicID string, name string, description string, duration time.Duration, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AddSubtaskRequestV1{
		EpicId:      epicID,
		Name:        name,
		Description: description,
		Duration:    durationpb.New(duration),
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.subtaskClient.AddSubtask(ctx, req)
	return err
}

func (s *SubtaskServiceImpl) UpdateSubtask(id string, epicID string, name string, description string, status model.TaskStatus,
	duration time.Duration, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateSubtaskRequestV1{
		Id:          id,
		EpicId:      epicID,
		Name:        name,
		Description: description,
		Status:      model.MapTaskStatusToMessageTaskStatus(status),
		Duration:    durationpb.New(duration),
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.subtaskClient.UpdateSubtask(ctx, req)
	return err
}

func (s *SubtaskServiceImpl) DeleteSubtask(subtaskID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteSubtaskRequestV1{
		Id: subtaskID,
	}

	_, err := s.subtaskClient.DeleteSubtask(ctx, req)
	return err
}

func (s *SubtaskServiceImpl) GetSubtaskByID(subtaskID string) (*pb.SubtaskV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetSubtaskByIDRequestV1{
		Id: subtaskID,
	}

	resp, err := s.subtaskClient.GetSubtaskByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSubtask(), nil
}

func (s *SubtaskServiceImpl) GetAllSubtasks() ([]*pb.SubtaskV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetSubtasksRequestV1{}
	resp, err := s.subtaskClient.GetSubtasks(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSubtasks(), nil
}
