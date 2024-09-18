package server

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

func TestEpicStorageCreate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicStorage := NewEpicStorage(mock, l)

	epic := model.NewEpic(
		"",
		"1",
		"Test Epic",
		"Test Epic Description",
		time.Now(),
		[]model.Subtask{},
	)

	mock.ExpectBegin()

	mock.ExpectQuery(`
		INSERT INTO gophertask.tasks \(name, user_id, description, status, duration, started_at\)
		VALUES \(@name, @user_id, @description, @status, @duration, @started_at\)
		RETURNING id
	`).WithArgs(
		epic.Name(),
		epic.UserID(),
		epic.Description(),
		epic.Status(),
		epic.Duration(),
		epic.StartedAt(),
	).WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow("1"))

	mock.ExpectExec(`
		INSERT INTO gophertask.epics \(id\)
		VALUES \(@id\)
	`).WithArgs("1").WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mock.ExpectCommit()

	err = epicStorage.Create(context.Background(), epic)
	assert.NoError(t, err, "Error creating epic")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestEpicStorageUpdate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicStorage := NewEpicStorage(mock, l)

	epic := model.NewEpic(
		"1",
		"1",
		"Updated Epic",
		"Updated Epic Description",
		time.Now(),
		[]model.Subtask{},
	)

	mock.ExpectExec(`
		UPDATE gophertask.tasks 
		SET name = @name, user_id = @user_id, description = @description, status = @status, duration = @duration, started_at = @started_at
		WHERE id = @id
	`).WithArgs(
		epic.Name(),
		epic.UserID(),
		epic.Description(),
		string(epic.Status()),
		epic.Duration(),
		epic.StartedAt(),
		epic.ID(),
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = epicStorage.Update(context.Background(), epic)
	assert.NoError(t, err, "Error updating epic")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestEpicStorageDelete(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicStorage := NewEpicStorage(mock, l)

	epicID := "1"

	mock.ExpectExec(`
		DELETE FROM gophertask.tasks WHERE id = @id
	`).WithArgs(epicID).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = epicStorage.Delete(context.Background(), epicID)
	assert.NoError(t, err, "Error deleting epic")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestEpicStorageGetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicStorage := NewEpicStorage(mock, l)

	epicID := "1"
	now := time.Now()
	subtask1 := model.NewSubtask("2", "1", epicID, "Subtask 1", "Subtask Description 1",
		model.New, time.Hour, now)
	subtask2 := model.NewSubtask("3", "1", epicID, "Subtask 2", "Subtask Description 2",
		model.InProgress, 2*time.Hour, now)

	expectedEpic := model.NewEpic(
		epicID,
		"1",
		"Test Epic",
		"Test Epic Description",
		now,
		[]model.Subtask{subtask1, subtask2},
	)

	mock.ExpectBeginTx(pgx.TxOptions{})
	mock.ExpectQuery(`
		SELECT id, user_id, name, description, started_at
		FROM gophertask.tasks
		WHERE id = @id
	`).WithArgs(epicID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "name", "description", "started_at"}).
			AddRow(expectedEpic.ID(), expectedEpic.UserID(), expectedEpic.Name(), expectedEpic.Description(), expectedEpic.StartedAt()))

	mock.ExpectQuery(`
		SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		INNER JOIN gophertask.subtasks as s
		ON t.id = s.id
		WHERE s.epic_id = @id
	`).WithArgs(epicID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "epic_id", "name", "description", "status", "duration", "started_at"}).
			AddRow(subtask1.ID(), subtask1.UserID(), subtask1.EpicID(), subtask1.Name(), subtask1.Description(),
				string(subtask1.Status()), subtask1.Duration(), subtask1.StartedAt()).
			AddRow(subtask2.ID(), subtask2.UserID(), subtask2.EpicID(), subtask2.Name(), subtask2.Description(),
				string(subtask2.Status()), subtask2.Duration(), subtask2.StartedAt()))

	mock.ExpectCommit()

	epic, err := epicStorage.GetByID(context.Background(), epicID)
	assert.NoError(t, err, "Error getting epic by ID")
	assert.Equal(t, expectedEpic, epic, "Returned epic does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestEpicStorageGetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicStorage := NewEpicStorage(mock, l)

	userID := "1"
	now := time.Now()

	rows := pgxmock.NewRows([]string{
		"id", "user_id", "name", "description", "started_at",
		"id", "user_id", "epic_id", "name", "description", "status", "duration", "started_at",
	}).
		AddRow("1", userID, "Test Epic 1", "Test Epic Description 1", now,
			"101", userID, "1", "Subtask 1", "Subtask Description 1", "IN PROGRESS", "1 hour", now,
		)

	mock.ExpectQuery(`
        SELECT e.id, e.user_id, e.name, e.description, e.started_at,
               s.id, s.user_id, st.epic_id, s.name, s.description, s.status, s.duration, s.started_at
        FROM gophertask.tasks as e
		INNER JOIN gophertask.epics as et ON e.id = et.id
        LEFT JOIN gophertask.subtasks as st ON e.id = st.epic_id
		LEFT JOIN gophertask.tasks as s ON st.id = s.id
        WHERE e.user_id = @user_id
    `).WithArgs(userID).
		WillReturnRows(rows)

	_, err = epicStorage.GetAll(context.Background(), userID)
	assert.NoError(t, err, "Error getting all epics")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}
