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

type SubtaskService interface {
	CreateSubtask(ctx context.Context, subtask model.Subtask) error
	UpdateSubtask(ctx context.Context, subtask model.Subtask) error
	DeleteSubtask(ctx context.Context, userID string, subtaskID string) error
	GetSubtaskByID(ctx context.Context, userID string, subtaskID string) (model.Subtask, error)
	GetAllSubtasks(ctx context.Context, userID string) ([]model.Subtask, error)
}

type SubtaskServiceImpl struct {
	subtaskStorage storage.SubtaskStorage
	logger         *logger.ServerLogger
}

func NewSubtaskService(subtaskStorage storage.SubtaskStorage, logger *logger.ServerLogger) SubtaskService {
	return &SubtaskServiceImpl{
		subtaskStorage: subtaskStorage,
		logger:         logger,
	}
}

func (s *SubtaskServiceImpl) CreateSubtask(ctx context.Context, subtask model.Subtask) error {
	err := s.subtaskStorage.Create(ctx, subtask)
	if err != nil {
		s.logger.Error("Failed save subtask", zap.String("event", "save subtask"),
			zap.Error(err))
		return fmt.Errorf("failed save subtask: %w", err)
	}

	return nil
}

func (s *SubtaskServiceImpl) UpdateSubtask(ctx context.Context, subtask model.Subtask) error {
	err := s.subtaskStorage.Update(ctx, subtask)
	if err != nil {
		s.logger.Error("Failed update subtask", zap.String("event", "update subtask"),
			zap.Error(err))
		return fmt.Errorf("failed update subtask: %w", err)
	}

	return nil
}

func (s *SubtaskServiceImpl) DeleteSubtask(ctx context.Context, userID string, subtaskID string) error {
	subtask, err := s.subtaskStorage.GetByID(ctx, subtaskID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Subtask does not exists", zap.String("event", "delete subtask"),
				zap.String("subtask id", subtaskID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("subtask with id:%s does not exist", subtaskID), err)
		default:
			s.logger.Error("Failed to get subtask by id", zap.String("event", "delete subtask"),
				zap.String("subtask id", subtaskID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get subtask by id: %w", err)
		}
	}

	if subtask.UserID() != userID {
		s.logger.Warn("User attempted to access subtask belonging to another user",
			zap.String("event", "delete subtask"), zap.String("subtask id", subtaskID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.subtaskStorage.Delete(ctx, subtaskID)
	if err != nil {
		s.logger.Error("Failed to delete subtask", zap.String("event", "delete subtask"),
			zap.String("subtask id", subtaskID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

func (s *SubtaskServiceImpl) GetSubtaskByID(ctx context.Context, userID, subtaskID string) (model.Subtask, error) {
	subtask, err := s.subtaskStorage.GetByID(ctx, subtaskID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Subtask does not exists", zap.String("event", "get subtask by id"),
				zap.String("subtask id", subtaskID), zap.String("user id", userID), zap.Error(err))
			return model.Subtask{}, er.NewNotFoundError(fmt.Sprintf("subtask with id:%s does not exist", subtaskID), err)
		default:
			s.logger.Error("Failed to get subtask by id", zap.String("event", "get subtask by id"),
				zap.String("subtask id", subtaskID), zap.String("user id", userID), zap.Error(err))
			return model.Subtask{}, fmt.Errorf("failed get subtask by id: %w", err)
		}
	}

	if subtask.UserID() != userID {
		s.logger.Warn("User attempted to access subtask belonging to another user",
			zap.String("event", "get subtask by id"), zap.String("subtask id", subtaskID),
			zap.String("user id", userID), zap.Error(err))
		return model.Subtask{}, er.NewForbiddenError("access denied", err)
	}

	return subtask, nil
}

func (s *SubtaskServiceImpl) GetAllSubtasks(ctx context.Context, userID string) ([]model.Subtask, error) {
	subtasks, err := s.subtaskStorage.GetAll(ctx, userID)
	if err != nil {
		s.logger.Error("Failed get subtasks", zap.String("event", "get subtasks"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Subtask{}, fmt.Errorf("failed get subtasks: %w", err)
	}

	return subtasks, nil
}
