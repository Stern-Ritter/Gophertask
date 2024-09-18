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

type TaskService interface {
	CreateTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, task model.Task) error
	DeleteTask(ctx context.Context, userID string, taskID string) error
	GetTaskByID(ctx context.Context, userID string, taskID string) (model.Task, error)
	GetAllTasks(ctx context.Context, userID string) ([]model.Task, error)
}

type TaskServiceImpl struct {
	taskStorage storage.TaskStorage
	logger      *logger.ServerLogger
}

func NewTaskService(storage storage.TaskStorage, logger *logger.ServerLogger) TaskService {
	return &TaskServiceImpl{
		taskStorage: storage,
		logger:      logger,
	}
}

func (s *TaskServiceImpl) CreateTask(ctx context.Context, task model.Task) error {
	err := s.taskStorage.Create(ctx, task)
	if err != nil {
		s.logger.Error("Failed save task", zap.String("event", "save task"),
			zap.Error(err))
		return fmt.Errorf("failed save task: %w", err)
	}

	return nil
}

func (s *TaskServiceImpl) UpdateTask(ctx context.Context, task model.Task) error {
	err := s.taskStorage.Update(ctx, task)
	if err != nil {
		s.logger.Error("Failed update task", zap.String("event", "update task"),
			zap.Error(err))
		return fmt.Errorf("failed update task: %w", err)
	}

	return nil
}

func (s *TaskServiceImpl) DeleteTask(ctx context.Context, userID string, taskID string) error {
	task, err := s.taskStorage.GetByID(ctx, taskID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Task does not exists", zap.String("event", "delete task"),
				zap.String("task id", taskID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("task with id:%s does not exist", taskID), err)
		default:
			s.logger.Error("Failed to get task by id", zap.String("event", "delete task"),
				zap.String("task id", taskID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get task by id: %w", err)
		}
	}

	if task.UserID() != userID {
		s.logger.Warn("User attempted to access task belonging to another user",
			zap.String("event", "delete task"), zap.String("task id", taskID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.taskStorage.Delete(ctx, taskID)
	if err != nil {
		s.logger.Error("Failed to delete task", zap.String("event", "delete task"),
			zap.String("task id", taskID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

func (s *TaskServiceImpl) GetTaskByID(ctx context.Context, userID string, taskID string) (model.Task, error) {
	task, err := s.taskStorage.GetByID(ctx, taskID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Task does not exists", zap.String("event", "get task by id"),
				zap.String("task id", taskID), zap.String("user id", userID), zap.Error(err))
			return model.Task{}, er.NewNotFoundError(fmt.Sprintf("task with id:%s does not exist", taskID), err)
		default:
			s.logger.Error("Failed to get task by id", zap.String("event", "get task by id"),
				zap.String("task id", taskID), zap.String("user id", userID), zap.Error(err))
			return model.Task{}, fmt.Errorf("failed get task by id: %w", err)
		}
	}

	if task.UserID() != userID {
		s.logger.Warn("User attempted to access task belonging to another user",
			zap.String("event", "get task by id"), zap.String("task id", taskID),
			zap.String("user id", userID), zap.Error(err))
		return model.Task{}, er.NewForbiddenError("access denied", err)
	}

	return task, nil
}

func (s *TaskServiceImpl) GetAllTasks(ctx context.Context, userID string) ([]model.Task, error) {
	tasks, err := s.taskStorage.GetAll(ctx, userID)
	if err != nil {
		s.logger.Error("Failed get tasks", zap.String("event", "get tasks"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Task{}, fmt.Errorf("failed get tasks: %w", err)
	}

	return tasks, nil
}
