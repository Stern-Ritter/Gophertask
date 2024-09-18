package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

type EpicStorage interface {
	Create(ctx context.Context, epic model.Epic) error
	Update(ctx context.Context, epic model.Epic) error
	Delete(ctx context.Context, epicID string) error
	GetByID(ctx context.Context, epicID string) (model.Epic, error)
	GetAll(ctx context.Context, userID string) ([]model.Epic, error)
}

type EpicStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

func NewEpicStorage(db PgxIface, logger *logger.ServerLogger) EpicStorage {
	return &EpicStorageImpl{
		db:     db,
		Logger: logger,
	}
}

func (s *EpicStorageImpl) Create(ctx context.Context, epic model.Epic) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	var epicID string
	err = tx.QueryRow(ctx, `
	INSERT INTO gophertask.tasks (name, user_id, description, status, duration, started_at)
	VALUES (@name, @user_id, @description, @status, @duration, @started_at)
	RETURNING id
`, pgx.NamedArgs{
		"name":        epic.Name(),
		"user_id":     epic.UserID(),
		"description": epic.Description(),
		"status":      epic.Status(),
		"duration":    epic.Duration(),
		"started_at":  epic.StartedAt(),
	}).Scan(&epicID)

	if err != nil {
		s.Logger.Debug("Failed to insert epic", zap.String("event", "add epic"),
			zap.String("user id", epic.UserID()), zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx, `
	INSERT INTO gophertask.epics (id)
	VALUES (@id)
`, pgx.NamedArgs{
		"id": epicID,
	})

	if err != nil {
		s.Logger.Debug("Failed to insert epic", zap.String("event", "add epic"),
			zap.String("user id", epic.UserID()), zap.Error(err))
		return err
	}

	return tx.Commit(ctx)
}

func (s *EpicStorageImpl) Update(ctx context.Context, epic model.Epic) error {
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
		"name":        epic.Name(),
		"user_id":     epic.UserID(),
		"description": epic.Description(),
		"status":      string(epic.Status()),
		"duration":    epic.Duration(),
		"started_at":  epic.StartedAt(),
		"id":          epic.ID(),
	})

	if err != nil {
		s.Logger.Debug("Failed to update epic", zap.String("event", "update epic"),
			zap.String("user id", epic.UserID()), zap.Error(err))
		return err
	}

	return nil
}

func (s *EpicStorageImpl) Delete(ctx context.Context, epicID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM gophertask.tasks WHERE id = @id
`, pgx.NamedArgs{
		"id": epicID,
	})

	if err != nil {
		s.Logger.Debug("Failed to delete epic", zap.String("event", "delete epic"),
			zap.String("epic id", epicID), zap.Error(err))
		return err
	}

	return nil
}

func (s *EpicStorageImpl) GetByID(ctx context.Context, epicID string) (model.Epic, error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return model.Epic{}, err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	row := tx.QueryRow(ctx, `
		SELECT id, user_id, name, description, started_at
		FROM gophertask.tasks
		WHERE id = @id
	`, pgx.NamedArgs{
		"id": epicID,
	})

	var id, userID, name, description string
	var startedAt time.Time
	err = row.Scan(&id, &userID, &name, &description, &startedAt)

	if err != nil {
		s.Logger.Debug("Failed to get epic", zap.String("event", "get epic"),
			zap.String("epic id", epicID), zap.Error(err))
		return model.Epic{}, err
	}

	rows, err := tx.Query(ctx, `
		SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		INNER JOIN gophertask.subtasks as s
		ON t.id = s.id
		WHERE s.epic_id = @id
	`, pgx.NamedArgs{
		"id": epicID,
	})

	if err != nil {
		s.Logger.Debug("Failed to get epic", zap.String("event", "get epic"),
			zap.String("epic id", epicID), zap.Error(err))
		return model.Epic{}, err
	}
	defer rows.Close()

	subtasks := make([]model.Subtask, 0)

	for rows.Next() {
		var subtaskID, subtaskUserID, subtaskEpicID, subtaskName, subtaskDescription, subtaskStatus string
		var subtaskDuration time.Duration
		var subtaskStartedAt time.Time
		err = rows.Scan(&subtaskID, &subtaskUserID, &subtaskEpicID, &subtaskName, &subtaskDescription, &subtaskStatus,
			&subtaskDuration, &subtaskStartedAt)

		if err != nil {
			s.Logger.Debug("Failed to get epic", zap.String("event", "get epic"),
				zap.String("epic id", epicID), zap.Error(err))
			return model.Epic{}, err
		}

		subtask := model.NewSubtask(subtaskID, subtaskUserID, subtaskEpicID, subtaskName, subtaskDescription,
			model.TaskStatus(subtaskStatus), subtaskDuration, subtaskStartedAt)

		subtasks = append(subtasks, subtask)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to get epic", zap.String("event", "get epic"),
			zap.String("epic id", epicID), zap.Error(err))
		return model.Epic{}, err
	}

	epic := model.NewEpic(id, userID, name, description, startedAt, subtasks)

	if err = tx.Commit(ctx); err != nil {
		s.Logger.Debug("Failed to get epic", zap.String("event", "get epic"),
			zap.String("epic id", epicID), zap.Error(err))
		return model.Epic{}, err
	}

	return epic, nil
}

func (s *EpicStorageImpl) GetAll(ctx context.Context, userID string) ([]model.Epic, error) {
	rows, err := s.db.Query(ctx, `
        SELECT e.id, e.user_id, e.name, e.description, e.started_at,
               s.id, s.user_id, st.epic_id, s.name, s.description, s.status, s.duration, s.started_at
        FROM gophertask.tasks as e
		INNER JOIN gophertask.epics as et ON e.id = et.id
        LEFT JOIN gophertask.subtasks as st ON e.id = st.epic_id
		LEFT JOIN gophertask.tasks as s ON st.id = s.id
        WHERE e.user_id = @user_id
    `, pgx.NamedArgs{
		"user_id": userID,
	})

	if err != nil {
		s.Logger.Debug("Failed to get epics", zap.String("event", "get epics"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	epicsMap := make(map[string]*model.Epic)

	for rows.Next() {
		var (
			epicID, epicUserID, epicName, epicDescription                                           string
			epicStartedAt                                                                           time.Time
			subtaskID, subtaskUserID, subtaskEpicID, subtaskName, subtaskDescription, subtaskStatus sql.NullString
			subtaskDuration                                                                         pgtype.Interval
			subtaskStartedAt                                                                        sql.NullTime
		)

		err = rows.Scan(
			&epicID, &epicUserID, &epicName, &epicDescription, &epicStartedAt,
			&subtaskID, &subtaskUserID, &subtaskEpicID, &subtaskName, &subtaskDescription, &subtaskStatus, &subtaskDuration,
			&subtaskStartedAt,
		)

		if err != nil {
			s.Logger.Debug("Failed to get epics", zap.String("event", "get epics"),
				zap.String("user id", userID), zap.Error(err))
			return nil, err
		}

		epic, exists := epicsMap[epicID]
		if !exists {
			subtasks := make([]model.Subtask, 0)
			e := model.NewEpic(epicID, epicUserID, epicName, epicDescription, epicStartedAt, subtasks)
			epic = &e
			epicsMap[epicID] = epic
		}

		if subtaskID.Valid {
			var duration time.Duration
			if subtaskDuration.Valid {
				duration = time.Duration(subtaskDuration.Microseconds) * time.Microsecond
			}
			subtask := model.NewSubtask(subtaskID.String, subtaskUserID.String, subtaskEpicID.String, subtaskName.String,
				subtaskDescription.String, model.TaskStatus(subtaskStatus.String), duration, subtaskStartedAt.Time)
			epic.AddSubtask(subtask)
		}
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to get epics", zap.String("event", "get epics"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}

	epics := make([]model.Epic, 0)
	for _, epic := range epicsMap {
		epics = append(epics, *epic)
	}

	return epics, nil
}
