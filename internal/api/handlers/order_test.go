package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jeeo/pack-management/internal/api/handlers"
	"github.com/jeeo/pack-management/internal/api/handlers/mocks"
	"github.com/jeeo/pack-management/internal/model"
	helpers "github.com/jeeo/pack-management/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderHandler_CalculateOrder(t *testing.T) {
	// Arrange
	orderAppMock := mocks.OrderApplication{}
	handler := handlers.NewOrderHandler(&orderAppMock)
	requestBody := handlers.CalculateOrderPackRequest{
		Amount: 1,
	}
	orderPacks := []model.OrderPack{
		{
			Pack:     model.Pack{Amount: 250},
			Quantity: 1,
		},
	}
	expectedResponseBody := handlers.CalculateOrderPackResponse{
		OrderPacks: handlers.ToOrderPackDTOs(orderPacks),
	}
	orderAppMock.On("CalculateOrderPack", mock.Anything, mock.Anything).Return(orderPacks, nil)
	recorder, request := helpers.CreateRequest(t, http.MethodPost, "/order/calculate", requestBody)

	// Act
	handler.CalculateOrder(recorder, request)
	responseBody := handlers.CalculateOrderPackResponse{}
	helpers.UnmarshalResponse(t, recorder, &responseBody)

	// Assert
	orderAppMock.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Equal(t, expectedResponseBody, responseBody)
}
