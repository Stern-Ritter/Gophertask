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

func TestTaskStorageCreate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	taskStorage := NewTaskStorage(mock, l)

	task := model.NewTask(
		"",
		"1",
		"Test Task",
		"Test Description",
		model.New,
		time.Hour,
		time.Now(),
	)

	mock.ExpectExec(`
		INSERT INTO gophertask.tasks \(name, user_id, description, status, duration, started_at\)
		VALUES \(@name, @user_id, @description, @status, @duration, @started_at\)
	`).WithArgs(
		task.Name(),
		task.UserID(),
		task.Description(),
		string(task.Status()),
		task.Duration(),
		task.StartedAt(),
	).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = taskStorage.Create(context.Background(), task)
	assert.NoError(t, err, "Error creating task")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestTaskStorageUpdate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	taskStorage := NewTaskStorage(mock, l)

	task := model.NewTask(
		"1",
		"1",
		"Updated Task",
		"Updated Description",
		model.InProgress,
		2*time.Hour,
		time.Now(),
	)

	mock.ExpectExec(`
		UPDATE gophertask.tasks 
		SET name = @name, user_id = @user_id, description = @description, status = @status, duration = @duration, started_at = @started_at
		WHERE id = @id
	`).WithArgs(
		task.Name(),
		task.UserID(),
		task.Description(),
		string(task.Status()),
		task.Duration(),
		task.StartedAt(),
		task.ID(),
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = taskStorage.Update(context.Background(), task)
	assert.NoError(t, err, "Error updating task")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestTaskStorageDelete(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	taskStorage := NewTaskStorage(mock, l)

	taskID := "1"

	mock.ExpectExec(`
		DELETE FROM gophertask.tasks WHERE id = @id
	`).WithArgs(taskID).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = taskStorage.Delete(context.Background(), taskID)
	assert.NoError(t, err, "Error deleting task")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestTaskStorageGetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	taskStorage := NewTaskStorage(mock, l)

	taskID := "1"
	expectedTask := model.NewTask(
		taskID,
		"1",
		"Test Task",
		"Test Description",
		model.New,
		time.Hour,
		time.Now(),
	)

	mock.ExpectQuery(`
		SELECT id, user_id, name, description, status, duration, started_at
		FROM gophertask.tasks
		WHERE id = @id
	`).WithArgs(taskID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "name", "description", "status", "duration", "started_at"}).
			AddRow(expectedTask.ID(), expectedTask.UserID(), expectedTask.Name(), expectedTask.Description(), string(expectedTask.Status()), expectedTask.Duration(), expectedTask.StartedAt()))

	task, err := taskStorage.GetByID(context.Background(), taskID)
	assert.NoError(t, err, "Error getting task by ID")
	assert.Equal(t, expectedTask, task, "Returned task does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestTaskStorageGetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error initializing connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	taskStorage := NewTaskStorage(mock, l)

	userID := "1"
	expectedTasks := []model.Task{
		model.NewTask("1", userID, "Test Task 1", "Test Description 1", model.New, time.Hour, time.Now()),
		model.NewTask("2", userID, "Test Task 2", "Test Description 2", model.Done, 2*time.Hour, time.Now()),
	}

	mock.ExpectQuery(`
		SELECT t.id, t.user_id, t.name, t.description, t.status, t.duration, t.started_at
		FROM gophertask.tasks as t
		LEFT JOIN gophertask.epics as e
		ON t.id = e.id
		LEFT JOIN gophertask.subtasks as s
		ON t.id = s.id
		WHERE t.user_id = @user_id AND
		e.id is NULL AND
		s.id is NULL
	`).WithArgs(userID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "name", "description", "status", "duration", "started_at"}).
			AddRow(expectedTasks[0].ID(), expectedTasks[0].UserID(), expectedTasks[0].Name(), expectedTasks[0].Description(), string(expectedTasks[0].Status()), expectedTasks[0].Duration(), expectedTasks[0].StartedAt()).
			AddRow(expectedTasks[1].ID(), expectedTasks[1].UserID(), expectedTasks[1].Name(), expectedTasks[1].Description(), string(expectedTasks[1].Status()), expectedTasks[1].Duration(), expectedTasks[1].StartedAt()))

	tasks, err := taskStorage.GetAll(context.Background(), userID)
	assert.NoError(t, err, "Error getting all tasks")
	assert.Equal(t, expectedTasks, tasks, "Returned tasks do not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}
