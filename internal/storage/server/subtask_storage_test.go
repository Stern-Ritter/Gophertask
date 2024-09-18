package server

import (
	"context"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

func TestSubtaskStorageCreate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskStorage := NewSubtaskStorage(mock, l)

	subtask := model.NewSubtask(
		"",
		"1",
		"2",
		"Test Subtask",
		"Test Description",
		model.New,
		time.Hour,
		time.Now(),
	)

	mock.ExpectBegin()

	mock.ExpectQuery(`
		INSERT INTO gophertask.tasks \(name, user_id, description, status, duration, started_at\)
		VALUES \(@name, @user_id, @description, @status, @duration, @started_at\)
		RETURNING id
	`).WithArgs(
		subtask.Name(),
		subtask.UserID(),
		subtask.Description(),
		subtask.Status(),
		subtask.Duration(),
		subtask.StartedAt(),
	).WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow("1"))

	mock.ExpectExec(`
		INSERT INTO gophertask.subtasks \(id, epic_id\)
		VALUES \(@id, @epic_id\)
	`).WithArgs("1", subtask.EpicID()).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mock.ExpectCommit()

	err = subtaskStorage.Create(context.Background(), subtask)
	assert.NoError(t, err, "Error creating subtask")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestSubtaskStorageUpdate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskStorage := NewSubtaskStorage(mock, l)

	subtask := model.NewSubtask(
		"1",
		"1",
		"2",
		"Updated Subtask",
		"Updated Description",
		model.InProgress,
		2*time.Hour,
		time.Now(),
	)

	mock.ExpectBegin()

	mock.ExpectExec(`
		UPDATE gophertask.tasks 
		SET name = @name, user_id = @user_id, description = @description, status = @status, duration = @duration, started_at = @started_at
		WHERE id = @id
	`).WithArgs(
		subtask.Name(),
		subtask.UserID(),
		subtask.Description(),
		string(subtask.Status()),
		subtask.Duration(),
		subtask.StartedAt(),
		subtask.ID(),
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(`
		UPDATE gophertask.subtasks 
		SET epic_id = @epic_id
		WHERE id = @id
	`).WithArgs(subtask.EpicID(), subtask.ID()).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectCommit()

	err = subtaskStorage.Update(context.Background(), subtask)
	assert.NoError(t, err, "Error updating subtask")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestSubtaskStorageDelete(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskStorage := NewSubtaskStorage(mock, l)

	subtaskID := "1"

	mock.ExpectExec(`
		DELETE FROM gophertask.tasks WHERE id = @id
	`).WithArgs(subtaskID).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = subtaskStorage.Delete(context.Background(), subtaskID)
	assert.NoError(t, err, "Error deleting subtask")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestSubtaskStorageGetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskStorage := NewSubtaskStorage(mock, l)

	subtaskID := "1"
	expectedSubtask := model.NewSubtask(
		subtaskID,
		"1",
		"2",
		"Test Subtask",
		"Test Description",
		model.New,
		time.Hour,
		time.Now(),
	)

	mock.ExpectQuery(`
		SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		INNER JOIN gophertask.subtasks as s 
		ON t.id = s.id
		WHERE s.id = @id
	`).WithArgs(subtaskID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "epic_id", "name", "description", "status", "duration", "started_at"}).
			AddRow(expectedSubtask.ID(), expectedSubtask.UserID(), expectedSubtask.EpicID(), expectedSubtask.Name(), expectedSubtask.Description(), string(expectedSubtask.Status()), expectedSubtask.Duration(), expectedSubtask.StartedAt()))

	subtask, err := subtaskStorage.GetByID(context.Background(), subtaskID)
	assert.NoError(t, err, "Error getting subtask by ID")
	assert.Equal(t, expectedSubtask, subtask, "Returned subtask does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestSubtaskStorageGetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskStorage := NewSubtaskStorage(mock, l)

	userID := "1"
	expectedSubtasks := []model.Subtask{
		model.NewSubtask("1", userID, "2", "Test Subtask 1", "Test Description 1", model.New, time.Hour, time.Now()),
		model.NewSubtask("2", userID, "3", "Test Subtask 2", "Test Description 2", model.Done, 2*time.Hour, time.Now()),
	}

	mock.ExpectQuery(`
		SELECT t.id, t.user_id, s.epic_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		INNER JOIN gophertask.subtasks as s 
		ON t.id = s.id
		WHERE t.userID = @userID
	`).WithArgs(userID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "epic_id", "name", "description", "status", "duration", "started_at"}).
			AddRow(expectedSubtasks[0].ID(), expectedSubtasks[0].UserID(), expectedSubtasks[0].EpicID(), expectedSubtasks[0].Name(), expectedSubtasks[0].Description(), string(expectedSubtasks[0].Status()), expectedSubtasks[0].Duration(), expectedSubtasks[0].StartedAt()).
			AddRow(expectedSubtasks[1].ID(), expectedSubtasks[1].UserID(), expectedSubtasks[1].EpicID(), expectedSubtasks[1].Name(), expectedSubtasks[1].Description(), string(expectedSubtasks[1].Status()), expectedSubtasks[1].Duration(), expectedSubtasks[1].StartedAt()))

	subtasks, err := subtaskStorage.GetAll(context.Background(), userID)
	assert.NoError(t, err, "Error getting all subtasks")
	assert.Equal(t, expectedSubtasks, subtasks, "Returned subtasks do not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}
