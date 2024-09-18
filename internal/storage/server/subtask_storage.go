package server

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

type SubtaskStorage interface {
	Create(ctx context.Context, subtask model.Subtask) error
	Update(ctx context.Context, subtask model.Subtask) error
	Delete(ctx context.Context, subtaskID string) error
	GetByID(ctx context.Context, subtaskID string) (model.Subtask, error)
	GetAll(ctx context.Context, userID string) ([]model.Subtask, error)
}

type SubtaskStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

func NewSubtaskStorage(db PgxIface, logger *logger.ServerLogger) SubtaskStorage {
	return &SubtaskStorageImpl{
		db:     db,
		Logger: logger,
	}
}

func (s *SubtaskStorageImpl) Create(ctx context.Context, subtask model.Subtask) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	var subtaskID string
	err = tx.QueryRow(ctx, `
	INSERT INTO gophertask.tasks (name, user_id, description, status, duration, started_at)
	VALUES (@name, @user_id, @description, @status, @duration, @started_at)
	RETURNING id
`, pgx.NamedArgs{
		"name":        subtask.Name(),
		"user_id":     subtask.UserID(),
		"description": subtask.Description(),
		"status":      subtask.Status(),
		"duration":    subtask.Duration(),
		"started_at":  subtask.StartedAt(),
	}).Scan(&subtaskID)

	if err != nil {
		s.Logger.Debug("Failed to insert subtask", zap.String("event", "add subtask"),
			zap.String("user id", subtask.UserID()), zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx, `
	INSERT INTO gophertask.subtasks (id, epic_id)
	VALUES (@id, @epic_id)
`, pgx.NamedArgs{
		"id":      subtaskID,
		"epic_id": subtask.EpicID(),
	})

	if err != nil {
		s.Logger.Debug("Failed to insert subtask", zap.String("event", "add subtask"),
			zap.String("user id", subtask.UserID()), zap.Error(err))
		return err
	}

	return tx.Commit(ctx)
}

func (s *SubtaskStorageImpl) Update(ctx context.Context, subtask model.Subtask) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	_, err = tx.Exec(ctx, `
	UPDATE gophertask.tasks 
	SET name = @name, 
		user_id = @user_id, 
		description = @description, 
		status = @status, 
		duration = @duration, 
		started_at = @started_at
	WHERE id = @id
`, pgx.NamedArgs{
		"name":        subtask.Name(),
		"user_id":     subtask.UserID(),
		"description": subtask.Description(),
		"status":      string(subtask.Status()),
		"duration":    subtask.Duration(),
		"started_at":  subtask.StartedAt(),
		"id":          subtask.ID(),
	})

	if err != nil {
		s.Logger.Debug("Failed to update subtask", zap.String("event", "update subtask"),
			zap.String("user id", subtask.UserID()), zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx, `
	UPDATE gophertask.subtasks 
	SET epic_id = @epic_id
	WHERE id = @id
`, pgx.NamedArgs{
		"epic_id": subtask.EpicID(),
		"id":      subtask.ID(),
	})

	if err != nil {
		s.Logger.Debug("Failed to update subtask", zap.String("event", "update subtask"),
			zap.String("user id", subtask.UserID()), zap.Error(err))
		return err
	}

	return tx.Commit(ctx)
}

func (s *SubtaskStorageImpl) Delete(ctx context.Context, subtaskID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM gophertask.tasks WHERE id = @id
`, pgx.NamedArgs{
		"id": subtaskID,
	})

	if err != nil {
		s.Logger.Debug("Failed to delete subtask", zap.String("event", "delete subtask"),
			zap.String("subtask id", subtaskID), zap.Error(err))
		return err
	}

	return nil
}

func (s *SubtaskStorageImpl) GetByID(ctx context.Context, subtaskID string) (model.Subtask, error) {
	row := s.db.QueryRow(ctx, `
	SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
	FROM gophertask.tasks as t
	INNER JOIN gophertask.subtasks as s 
	ON t.id = s.id
	WHERE s.id = @id
	`, pgx.NamedArgs{
		"id": subtaskID,
	})

	var id, userID, epicID, name, description, status string
	var duration time.Duration
	var startedAt time.Time
	err := row.Scan(&id, &userID, &epicID, &name, &description, &status, &duration, &startedAt)

	if err != nil {
		s.Logger.Debug("Failed to get subtask", zap.String("event", "get subtask"),
			zap.String("subtask id", subtaskID), zap.Error(err))
		return model.Subtask{}, err
	}

	subtask := model.NewSubtask(id, userID, epicID, name, description, model.TaskStatus(status), duration, startedAt)

	return subtask, nil
}

func (s *SubtaskStorageImpl) GetAll(ctx context.Context, userID string) ([]model.Subtask, error) {
	rows, err := s.db.Query(ctx, `
		SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		INNER JOIN gophertask.subtasks as s 
		ON t.id = s.id
		WHERE t.userID = @userID
	`, pgx.NamedArgs{
		"userID": userID,
	})

	if err != nil {
		s.Logger.Debug("Failed to get subtasks", zap.String("event", "get subtasks"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	subtasks := make([]model.Subtask, 0)

	for rows.Next() {
		var id, epicID, name, description, status string
		var duration time.Duration
		var startedAt time.Time
		err = rows.Scan(&id, &userID, &epicID, &name, &description, &status, &duration, &startedAt)

		if err != nil {
			s.Logger.Debug("Failed to get subtasks", zap.String("event", "get subtasks"),
				zap.String("user id", userID), zap.Error(err))
			return nil, err
		}

		subtask := model.NewSubtask(id, userID, epicID, name, description, model.TaskStatus(status), duration, startedAt)

		subtasks = append(subtasks, subtask)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to get subtasks", zap.String("event", "get subtasks"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}

	return subtasks, nil
}
