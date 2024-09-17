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

func TestEpicServiceCreateEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockEpicStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicService := NewEpicService(mockStorage, l)

	ctx := context.Background()
	epic := model.NewEpic("1", "user1", "Test Epic", "This is a test epic", time.Now(), []model.Subtask{})

	mockStorage.EXPECT().Create(ctx, epic).Return(nil).Times(1)

	err = epicService.CreateEpic(ctx, epic)
	assert.NoError(t, err, "expected no error when creating an epic")
}

func TestEpicServiceUpdateEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockEpicStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicService := NewEpicService(mockStorage, l)

	ctx := context.Background()
	epic := model.NewEpic("1", "user1", "Updated Epic", "Updated description", time.Now(), []model.Subtask{})

	mockStorage.EXPECT().Update(ctx, epic).Return(nil).Times(1)

	err = epicService.UpdateEpic(ctx, epic)
	assert.NoError(t, err, "expected no error when updating an epic")
}

func TestEpicServiceDeleteEpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockEpicStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicService := NewEpicService(mockStorage, l)

	ctx := context.Background()
	epicID := "1"
	userID := "user1"

	mockStorage.EXPECT().GetByID(ctx, epicID).Return(model.NewEpic(epicID, userID, "Test Epic", "Test description",
		time.Now(), []model.Subtask{}), nil).Times(1)
	mockStorage.EXPECT().Delete(ctx, epicID).Return(nil).Times(1)

	err = epicService.DeleteEpic(ctx, userID, epicID)
	assert.NoError(t, err, "expected no error when deleting an epic")
}

func TestEpicServiceGetEpicByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockEpicStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicService := NewEpicService(mockStorage, l)

	ctx := context.Background()
	epicID := "1"
	userID := "user1"
	expectedEpic := model.NewEpic(epicID, userID, "Test Epic", "Test description", time.Now(), []model.Subtask{})

	mockStorage.EXPECT().GetByID(ctx, epicID).Return(expectedEpic, nil).Times(1)

	epic, err := epicService.GetEpicByID(ctx, userID, epicID)
	assert.NoError(t, err, "expected no error when getting an epic by ID")
	assert.Equal(t, expectedEpic, epic, "expected epics to be equal")
}

func TestEpicServiceGetAllEpics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockEpicStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error initializing logger")

	epicService := NewEpicService(mockStorage, l)

	ctx := context.Background()
	userID := "user1"
	expectedEpics := []model.Epic{
		model.NewEpic("1", userID, "Epic 1", "Epic description 1", time.Now(), []model.Subtask{}),
		model.NewEpic("2", userID, "Epic 2", "Epic description 2", time.Now(), []model.Subtask{}),
	}

	mockStorage.EXPECT().GetAll(ctx, userID).Return(expectedEpics, nil).Times(1)

	epics, err := epicService.GetAllEpics(ctx, userID)
	assert.NoError(t, err, "expected no error when getting all epics")
	assert.Equal(t, expectedEpics, epics, "expected epic lists to be equal")
}
