package application_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/jeeo/pack-management/internal/application"
	"github.com/jeeo/pack-management/internal/application/mocks"
	domainerrors "github.com/jeeo/pack-management/internal/errors"
	"github.com/jeeo/pack-management/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPackApplication_FindAll(t *testing.T) {
	mockRepo := &mocks.PackRepository{}
	app := application.NewPackApplication(mockRepo)

	// Mocking the FindAll method of packRepository
	mockRepo.On("FindAll", mock.Anything).Return([]model.Pack{{ID: "1", Amount: 10}}, nil)

	packs, err := app.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, packs, 1)
	assert.Equal(t, "1", packs[0].ID)

	mockRepo.AssertExpectations(t)
}

func TestPackApplication_Create(t *testing.T) {
	mockRepo := &mocks.PackRepository{}
	app := application.NewPackApplication(mockRepo)

	// Mocking the Create method of packRepository
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("model.Pack")).Return(nil)

	err := app.Create(context.Background(), 5)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestPackApplication_Update(t *testing.T) {
	mockRepo := &mocks.PackRepository{}
	app := application.NewPackApplication(mockRepo)

	// Mocking the FindByID and Update methods of packRepository
	mockRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.Pack{ID: "1", Amount: 10}, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("model.Pack")).Return(nil)

	err := app.Update(context.Background(), model.Pack{ID: "1", Amount: 20})

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestPackApplication_Delete(t *testing.T) {
	mockRepo := &mocks.PackRepository{}
	app := application.NewPackApplication(mockRepo)

	// Mocking the FindByID and Delete methods of packRepository
	mockRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.Pack{ID: "1", Amount: 10}, nil)
	mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

	err := app.Delete(context.Background(), "1")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestPackApplication_Delete_PackNotFound(t *testing.T) {
	mockRepo := &mocks.PackRepository{}
	app := application.NewPackApplication(mockRepo)
	refErr := domainerrors.ErrPackNotFound{}

	// Mocking the FindByID method of packRepository to simulate ErrPackNotFound
	mockRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.Pack{}, sql.ErrNoRows)

	err := app.Delete(context.Background(), "1")

	assert.Error(t, err)
	assert.True(t, errors.As(err, &refErr))

	mockRepo.AssertExpectations(t)
}
