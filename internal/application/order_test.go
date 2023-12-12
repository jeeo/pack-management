package application_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jeeo/pack-management/internal/application"
	"github.com/jeeo/pack-management/internal/application/mocks"
	"github.com/jeeo/pack-management/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculate(t *testing.T) {
	// Arrange
	packRepoMock := mocks.PackRepository{}
	app := application.NewOrderApplication(&packRepoMock)
	testCases := []struct {
		OrderQuantity int
		Expected      []model.OrderPack
	}{
		{
			OrderQuantity: 1,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 250,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 251,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 500},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 501,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
				{
					Pack:     model.Pack{Amount: 500},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 12001,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
				{
					Pack:     model.Pack{Amount: 2000},
					Quantity: 1,
				},
				{
					Pack:     model.Pack{Amount: 5000},
					Quantity: 2,
				},
			},
		},
		{
			OrderQuantity: 750,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
				{
					Pack:     model.Pack{Amount: 500},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 751,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 1000},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 900,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 1000},
					Quantity: 1,
				},
			},
		},
	}
	packRepoMock.On("FindAll", mock.Anything).Return([]model.Pack{
		{Amount: 250},
		{Amount: 500},
		{Amount: 1000},
		{Amount: 2000},
		{Amount: 5000},
	}, nil)

	// Act + Assert
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("OrderQuantity: %d", testCase.OrderQuantity), func(t *testing.T) {

			resultOrder, err := app.CalculateOrderPack(context.TODO(), testCase.OrderQuantity)
			if err != nil {
				t.Error(err)
			}

			assert.ElementsMatch(t, resultOrder, testCase.Expected)
		})
	}
}

func TestCalculate2(t *testing.T) {
	// Arrange
	packRepoMock := mocks.PackRepository{}
	app := application.NewOrderApplication(&packRepoMock)
	testCases := []struct {
		OrderQuantity int
		Expected      []model.OrderPack
	}{
		{
			OrderQuantity: 500000,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 53},
					Quantity: 9434,
				},
			},
		},
	}
	packRepoMock.On("FindAll", mock.Anything).Return([]model.Pack{
		{Amount: 23},
		{Amount: 31},
		{Amount: 53},
	}, nil)

	// Act + Assert
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("OrderQuantity: %d", testCase.OrderQuantity), func(t *testing.T) {

			resultOrder, err := app.CalculateOrderPack(context.TODO(), testCase.OrderQuantity)
			if err != nil {
				t.Error(err)
			}

			assert.ElementsMatch(t, resultOrder, testCase.Expected)
		})
	}
}
