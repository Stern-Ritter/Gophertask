package server

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

func TestTaskServiceCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockTaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	taskService := NewTaskService(mockStorage, l)

	ctx := context.Background()
	task := model.NewTask("1", "user1", "Test Task", "This is a test task", model.New, time.Hour, time.Now())

	mockStorage.EXPECT().Create(ctx, task).Return(nil).Times(1)

	err = taskService.CreateTask(ctx, task)
	assert.NoError(t, err, "expected no error when creating a task")
}

func TestTaskServiceUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockTaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	taskService := NewTaskService(mockStorage, l)

	ctx := context.Background()
	task := model.NewTask("1", "user1", "Updated Task", "Updated description", model.InProgress, time.Hour, time.Now())

	mockStorage.EXPECT().Update(ctx, task).Return(nil).Times(1)

	err = taskService.UpdateTask(ctx, task)
	assert.NoError(t, err, "expected no error when updating a task")
}

func TestTaskServiceDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockTaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	taskService := NewTaskService(mockStorage, l)

	ctx := context.Background()
	taskID := "1"
	userID := "user1"

	mockStorage.EXPECT().GetByID(ctx, taskID).Return(model.NewTask(taskID, userID, "Test Task", "Test description", model.New, time.Hour, time.Now()), nil).Times(1)
	mockStorage.EXPECT().Delete(ctx, taskID).Return(nil).Times(1)

	err = taskService.DeleteTask(ctx, userID, taskID)
	assert.NoError(t, err, "expected no error when deleting a task")
}

func TestTaskServiceGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockTaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	taskService := NewTaskService(mockStorage, l)

	ctx := context.Background()
	taskID := "1"
	userID := "user1"
	expectedTask := model.NewTask(taskID, userID, "Test Task", "Test description", model.New, time.Hour, time.Now())

	mockStorage.EXPECT().GetByID(ctx, taskID).Return(expectedTask, nil).Times(1)

	task, err := taskService.GetTaskByID(ctx, userID, taskID)
	assert.NoError(t, err, "expected no error when getting a task by ID")
	assert.Equal(t, expectedTask, task, "expected tasks to be equal")
}

func TestTaskServiceGetAllTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockTaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	taskService := NewTaskService(mockStorage, l)

	ctx := context.Background()
	userID := "user1"
	expectedTasks := []model.Task{
		model.NewTask("1", userID, "Task 1", "Task description 1", model.New, time.Hour, time.Now()),
		model.NewTask("2", userID, "Task 2", "Task description 2", model.InProgress, time.Hour, time.Now()),
	}

	mockStorage.EXPECT().GetAll(ctx, userID).Return(expectedTasks, nil).Times(1)

	tasks, err := taskService.GetAllTasks(ctx, userID)
	assert.NoError(t, err, "expected no error when getting all tasks")
	assert.Equal(t, expectedTasks, tasks, "expected task lists to be equal")
}
