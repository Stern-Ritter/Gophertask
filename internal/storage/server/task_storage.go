package server

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

type TaskStorage interface {
	Create(ctx context.Context, task model.Task) error
	Update(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, taskID string) error
	GetByID(ctx context.Context, taskID string) (model.Task, error)
	GetAll(ctx context.Context, userID string) ([]model.Task, error)
}

type TaskStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

func NewTaskStorage(db PgxIface, logger *logger.ServerLogger) TaskStorage {
	return &TaskStorageImpl{
		db:     db,
		Logger: logger,
	}
}

func (s *TaskStorageImpl) Create(ctx context.Context, task model.Task) error {
	_, err := s.db.Exec(ctx, `
	INSERT INTO gophertask.tasks (name, user_id, description, status, duration, started_at)
	VALUES (@name, @user_id, @description, @status, @duration, @started_at)
`, pgx.NamedArgs{
		"name":        task.Name(),
		"user_id":     task.UserID(),
		"description": task.Description(),
		"status":      string(task.Status()),
		"duration":    task.Duration(),
		"started_at":  task.StartedAt(),
	})

	if err != nil {
		s.Logger.Debug("Failed to insert task", zap.String("event", "add task"),
			zap.String("user id", task.UserID()), zap.Error(err))
		return err
	}

	return nil
}

func (s *TaskStorageImpl) Update(ctx context.Context, task model.Task) error {
	_, err := s.db.Exec(ctx, `
	UPDATE gophertask.tasks 
	SET name = @name, 
		user_id = @user_id, 
		description = @description, 
		status = @status, 
		duration = @duration, 
		started_at = @started_at
	WHERE id = @id
`, pgx.NamedArgs{
		"name":        task.Name(),
		"user_id":     task.UserID(),
		"description": task.Description(),
		"status":      string(task.Status()),
		"duration":    task.Duration(),
		"started_at":  task.StartedAt(),
		"id":          task.ID(),
	})

	if err != nil {
		s.Logger.Debug("Failed to update task", zap.String("event", "update task"),
			zap.String("user id", task.UserID()), zap.Error(err))
		return err
	}

	return nil
}

func (s *TaskStorageImpl) Delete(ctx context.Context, taskID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM gophertask.tasks WHERE id = @id
`, pgx.NamedArgs{
		"id": taskID,
	})

	if err != nil {
		s.Logger.Debug("Failed to delete task", zap.String("event", "delete task"),
			zap.String("task id", taskID), zap.Error(err))
		return err
	}

	return nil
}

func (s *TaskStorageImpl) GetByID(ctx context.Context, taskID string) (model.Task, error) {
	row := s.db.QueryRow(ctx, `
	SELECT id, user_id, name, description, status, duration, started_at
	FROM gophertask.tasks
	WHERE id = @id
	`, pgx.NamedArgs{
		"id": taskID,
	})

	var id, userID, name, description, status string
	var duration time.Duration
	var startedAt time.Time
	err := row.Scan(&id, &userID, &name, &description, &status, &duration, &startedAt)

	if err != nil {
		s.Logger.Debug("Failed to get task", zap.String("event", "get task"),
			zap.String("task id", taskID), zap.Error(err))
		return model.Task{}, err
	}

	task := model.NewTask(id, userID, name, description, model.TaskStatus(status), duration, startedAt)

	return task, nil
}

func (s *TaskStorageImpl) GetAll(ctx context.Context, userID string) ([]model.Task, error) {
	rows, err := s.db.Query(ctx, `
		SELECT t.id, t.user_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		LEFT JOIN gophertask.epics as e
		ON t.id = e.id
		LEFT JOIN gophertask.subtasks as s
		ON t.id = s.id
		WHERE t.user_id = @user_id AND
		e.id is NULL AND
		s.id is NULL
	`, pgx.NamedArgs{
		"user_id": userID,
	})

	if err != nil {
		s.Logger.Debug("Failed to get tasks", zap.String("event", "get tasks"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)

	for rows.Next() {
		var id, name, description, status string
		var duration time.Duration
		var startedAt time.Time
		err = rows.Scan(&id, &userID, &name, &description, &status, &duration, &startedAt)

		if err != nil {
			s.Logger.Debug("Failed to get tasks", zap.String("event", "get tasks"),
				zap.String("user id", userID), zap.Error(err))
			return nil, err
		}

		task := model.NewTask(id, userID, name, description, model.TaskStatus(status), duration, startedAt)

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to get tasks", zap.String("event", "get tasks"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}

	return tasks, nil
}
