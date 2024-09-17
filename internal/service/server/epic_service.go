package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	er "github.com/Stern-Ritter/gophertask/internal/errors"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
	storage "github.com/Stern-Ritter/gophertask/internal/storage/server"
)

type EpicService interface {
	CreateEpic(ctx context.Context, epic model.Epic) error
	UpdateEpic(ctx context.Context, epic model.Epic) error
	DeleteEpic(ctx context.Context, userID string, epicID string) error
	GetEpicByID(ctx context.Context, userID string, epicID string) (model.Epic, error)
	GetAllEpics(ctx context.Context, userID string) ([]model.Epic, error)
}

type EpicServiceImpl struct {
	epicStorage storage.EpicStorage
	logger      *logger.ServerLogger
}

func NewEpicService(epicStorage storage.EpicStorage, logger *logger.ServerLogger) EpicService {
	return &EpicServiceImpl{
		epicStorage: epicStorage,
		logger:      logger,
	}
}

func (s *EpicServiceImpl) CreateEpic(ctx context.Context, epic model.Epic) error {
	err := s.epicStorage.Create(ctx, epic)
	if err != nil {
		s.logger.Error("Failed save epic", zap.String("event", "save epic"),
			zap.Error(err))
		return fmt.Errorf("failed save epic: %w", err)
	}

	return nil
}

func (s *EpicServiceImpl) UpdateEpic(ctx context.Context, epic model.Epic) error {
	err := s.epicStorage.Update(ctx, epic)
	if err != nil {
		s.logger.Error("Failed update epic", zap.String("event", "update epic"),
			zap.Error(err))
		return fmt.Errorf("failed update epic: %w", err)
	}

	return nil
}

func (s *EpicServiceImpl) DeleteEpic(ctx context.Context, userID string, epicID string) error {
	epic, err := s.epicStorage.GetByID(ctx, epicID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Epic does not exists", zap.String("event", "delete epic"),
				zap.String("epic id", epicID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("epic with id:%s does not exist", epicID), err)
		default:
			s.logger.Error("Failed to get epic by id", zap.String("event", "delete epic"),
				zap.String("epic id", epicID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get epic by id: %w", err)
		}
	}

	if epic.UserID() != userID {
		s.logger.Warn("User attempted to access epic belonging to another user",
			zap.String("event", "delete epic"), zap.String("epic id", epicID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.epicStorage.Delete(ctx, epicID)
	if err != nil {
		s.logger.Error("Failed to delete epic", zap.String("event", "delete epic"),
			zap.String("epic id", epicID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

func (s *EpicServiceImpl) GetEpicByID(ctx context.Context, userID string, epicID string) (model.Epic, error) {
	epic, err := s.epicStorage.GetByID(ctx, epicID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Epic does not exists", zap.String("event", "get epic by id"),
				zap.String("epic id", epicID), zap.String("user id", userID), zap.Error(err))
			return model.Epic{}, er.NewNotFoundError(fmt.Sprintf("epic with id:%s does not exist", epicID), err)
		default:
			s.logger.Error("Failed to get epic by id", zap.String("event", "get epic by id"),
				zap.String("epic id", epicID), zap.String("user id", userID), zap.Error(err))
			return model.Epic{}, fmt.Errorf("failed get epic by id: %w", err)
		}
	}

	if epic.UserID() != userID {
		s.logger.Warn("User attempted to access epic belonging to another user",
			zap.String("event", "get epic by id"), zap.String("epic id", epicID),
			zap.String("user id", userID), zap.Error(err))
		return model.Epic{}, er.NewForbiddenError("access denied", err)
	}

	return epic, nil
}

func (s *EpicServiceImpl) GetAllEpics(ctx context.Context, userID string) ([]model.Epic, error) {
	epics, err := s.epicStorage.GetAll(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to get epics", zap.String("event", "get epics"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Epic{}, fmt.Errorf("failed to get epics: %w", err)
	}

	return epics, nil
}
