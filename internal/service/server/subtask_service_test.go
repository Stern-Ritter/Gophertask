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

func TestCreateSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockSubtaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskService := NewSubtaskService(mockStorage, l)

	ctx := context.Background()
	subtask := model.NewSubtask("1", "user1", "epic1", "Test Subtask", "This is a test subtask", model.New, time.Hour, time.Now())

	mockStorage.EXPECT().Create(ctx, subtask).Return(nil).Times(1)

	err = subtaskService.CreateSubtask(ctx, subtask)
	assert.NoError(t, err, "expected no error when creating a subtask")
}

func TestSubtaskServiceUpdateSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockSubtaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskService := NewSubtaskService(mockStorage, l)

	ctx := context.Background()
	subtask := model.NewSubtask("1", "user1", "epic1", "Updated Subtask", "Updated description", model.InProgress, time.Hour, time.Now())

	mockStorage.EXPECT().Update(ctx, subtask).Return(nil).Times(1)

	err = subtaskService.UpdateSubtask(ctx, subtask)
	assert.NoError(t, err, "expected no error when updating a subtask")
}

func TestSubtaskServiceDeleteSubtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockSubtaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskService := NewSubtaskService(mockStorage, l)

	ctx := context.Background()
	subtaskID := "1"
	userID := "user1"

	mockStorage.EXPECT().GetByID(ctx, subtaskID).Return(model.NewSubtask(subtaskID, userID, "epic1", "Test Subtask", "Test description", model.New, time.Hour, time.Now()), nil).Times(1)
	mockStorage.EXPECT().Delete(ctx, subtaskID).Return(nil).Times(1)

	err = subtaskService.DeleteSubtask(ctx, userID, subtaskID)
	assert.NoError(t, err, "expected no error when deleting a subtask")
}

func TestSubtaskServiceGetSubtaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockSubtaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskService := NewSubtaskService(mockStorage, l)

	ctx := context.Background()
	subtaskID := "1"
	userID := "user1"
	expectedSubtask := model.NewSubtask(subtaskID, userID, "epic1", "Test Subtask", "Test description", model.New, time.Hour, time.Now())

	mockStorage.EXPECT().GetByID(ctx, subtaskID).Return(expectedSubtask, nil).Times(1)

	subtask, err := subtaskService.GetSubtaskByID(ctx, userID, subtaskID)
	assert.NoError(t, err, "expected no error when getting a subtask by ID")
	assert.Equal(t, expectedSubtask, subtask, "expected subtasks to be equal")
}

func TestSubtaskServiceGetAllSubtasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockSubtaskStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	subtaskService := NewSubtaskService(mockStorage, l)

	ctx := context.Background()
	userID := "user1"
	expectedSubtasks := []model.Subtask{
		model.NewSubtask("1", userID, "epic1", "Subtask 1", "Subtask description 1", model.New, time.Hour, time.Now()),
		model.NewSubtask("2", userID, "epic1", "Subtask 2", "Subtask description 2", model.InProgress, time.Hour, time.Now()),
	}

	mockStorage.EXPECT().GetAll(ctx, userID).Return(expectedSubtasks, nil).Times(1)

	subtasks, err := subtaskService.GetAllSubtasks(ctx, userID)
	assert.NoError(t, err, "expected no error when getting all subtasks")
	assert.Equal(t, expectedSubtasks, subtasks, "expected subtask lists to be equal")
}
