package handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jeeo/pack-management/internal/api/handlers"
	"github.com/jeeo/pack-management/internal/api/handlers/mocks"
	"github.com/jeeo/pack-management/internal/model"
	helpers "github.com/jeeo/pack-management/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPackHandler_FindAll(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)
	packs := handlers.FindAllPackResponse{}
	mockApp.On("FindAll", mock.Anything).Return([]model.Pack{{ID: "1", Amount: 10}}, nil)

	recorder, request := helpers.CreateRequest(t, http.MethodGet, "/packs", nil)
	handler.FindAll(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	helpers.UnmarshalResponse(t, recorder, &packs)
	assert.Len(t, packs.Packs, 1)
	assert.Equal(t, "1", packs.Packs[0].ID)

	mockApp.AssertExpectations(t)
}

func TestPackHandler_FindAll_Error(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)

	mockApp.On("FindAll", mock.Anything).Return(nil, errors.New("some error"))
	recorder, request := helpers.CreateRequest(t, http.MethodGet, "/packs", nil)

	handler.FindAll(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Create(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)

	mockApp.On("Create", mock.Anything, mock.AnythingOfType("int")).Return(nil)
	recorder, request := helpers.CreateRequest(t, http.MethodPost, "/packs", handlers.CreatePackRequest{Amount: 5})

	handler.Create(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Create_InvalidRequest(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)

	recorder, request := helpers.CreateRequest(t, http.MethodPost, "/packs", struct{}{})

	handler.Create(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Update(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)
	packId := uuid.New().String()
	mockApp.On("Update", mock.Anything, mock.AnythingOfType("model.Pack")).Return(nil)
	recorder, request := helpers.CreateRequest(t, http.MethodPut, "/package/"+packId, handlers.UpdatePackRequest{Amount: 5})
	muxVars := map[string]string{
		"id": packId,
	}
	request = mux.SetURLVars(request, muxVars)
	handler.Update(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Update_InvalidRequest(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)

	recorder, request := helpers.CreateRequest(t, http.MethodPut, "/packs/1", struct{}{})

	handler.Update(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Delete(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)
	packId := uuid.New().String()
	mockApp.On("Delete", mock.Anything, mock.Anything).Return(nil)
	recorder, request := helpers.CreateRequest(t, http.MethodDelete, "/package/"+packId, nil)
	muxVars := map[string]string{
		"id": packId,
	}
	request = mux.SetURLVars(request, muxVars)
	handler.Delete(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockApp.AssertExpectations(t)
}

func TestPackHandler_Delete_InvalidRequest(t *testing.T) {
	mockApp := &mocks.PackApplication{}
	handler := handlers.NewPackHandler(mockApp)

	recorder, request := helpers.CreateRequest(t, http.MethodDelete, "/packs/1", nil)

	handler.Delete(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockApp.AssertExpectations(t)
}
