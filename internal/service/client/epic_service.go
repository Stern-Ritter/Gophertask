package client

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type EpicService interface {
	CreateEpic(name string, description string, startedAt time.Time) error
	UpdateEpic(id string, name string, description string, startedAt time.Time) error
	DeleteEpic(epicID string) error
	GetEpicByID(epicID string) (*pb.EpicV1, error)
	GetAllEpics() ([]*pb.EpicV1, error)
}

type EpicServiceImpl struct {
	epicClient pb.EpicServiceV1Client
}

func NewEpicService(epicClient pb.EpicServiceV1Client) EpicService {
	return &EpicServiceImpl{
		epicClient: epicClient,
	}
}

func (s *EpicServiceImpl) CreateEpic(name string, description string, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AddEpicRequestV1{
		Name:        name,
		Description: description,
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.epicClient.AddEpic(ctx, req)
	return err
}

func (s *EpicServiceImpl) UpdateEpic(id string, name string, description string, startedAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateEpicRequestV1{
		Id:          id,
		Name:        name,
		Description: description,
		StartedAt:   timestamppb.New(startedAt),
	}

	_, err := s.epicClient.UpdateEpic(ctx, req)
	return err
}

func (s *EpicServiceImpl) DeleteEpic(epicID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteEpicRequestV1{
		Id: epicID,
	}

	_, err := s.epicClient.DeleteEpic(ctx, req)
	return err
}

func (s *EpicServiceImpl) GetEpicByID(epicID string) (*pb.EpicV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetEpicByIDRequestV1{
		Id: epicID,
	}

	resp, err := s.epicClient.GetEpicByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetEpic(), nil
}

func (s *EpicServiceImpl) GetAllEpics() ([]*pb.EpicV1, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	req := &pb.GetEpicsRequestV1{}
	resp, err := s.epicClient.GetEpics(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetEpics(), nil
}
