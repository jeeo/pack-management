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
	defaultStoredPacks := []model.Pack{
		{Amount: 250},
		{Amount: 500},
		{Amount: 1000},
		{Amount: 2000},
		{Amount: 5000},
	}
	testCases := []struct {
		OrderQuantity int
		StoredPacks   []model.Pack
		Expected      []model.OrderPack
	}{
		{
			OrderQuantity: 1,
			StoredPacks:   defaultStoredPacks,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 250,
			StoredPacks:   defaultStoredPacks,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 250},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 251,
			StoredPacks:   defaultStoredPacks,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 500},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 501,
			StoredPacks:   defaultStoredPacks,
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
			StoredPacks:   defaultStoredPacks,
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
			StoredPacks:   defaultStoredPacks,
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
			StoredPacks:   defaultStoredPacks,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 1000},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 900,
			StoredPacks:   defaultStoredPacks,
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 1000},
					Quantity: 1,
				},
			},
		},
		{
			OrderQuantity: 500000,
			StoredPacks: []model.Pack{
				{Amount: 23},
				{Amount: 31},
				{Amount: 53},
			},
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 53},
					Quantity: 9434,
				},
			},
		},
		{
			OrderQuantity: 270,
			StoredPacks: []model.Pack{
				{Amount: 5},
				{Amount: 100},
				{Amount: 500},
			},
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 5},
					Quantity: 14,
				},
				{
					Pack:     model.Pack{Amount: 100},
					Quantity: 2,
				},
			},
		},
		{
			OrderQuantity: 300,
			StoredPacks: []model.Pack{
				{Amount: 5},
				{Amount: 100},
				{Amount: 500},
			},
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 100},
					Quantity: 3,
				},
			},
		},
		{
			OrderQuantity: 301,
			StoredPacks: []model.Pack{
				{Amount: 5},
				{Amount: 100},
				{Amount: 500},
			},
			Expected: []model.OrderPack{
				{
					Pack:     model.Pack{Amount: 5},
					Quantity: 1,
				},
				{
					Pack:     model.Pack{Amount: 100},
					Quantity: 3,
				},
			},
		},
	}

	// Act + Assert
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("OrderQuantity: %d", testCase.OrderQuantity), func(t *testing.T) {
			packRepoMock.On("FindAll", mock.Anything).Return(testCase.StoredPacks, nil).Once()

			resultOrder, err := app.CalculateOrderPack(context.TODO(), testCase.OrderQuantity)
			if err != nil {
				t.Error(err)
			}

			assert.ElementsMatch(t, resultOrder, testCase.Expected)
		})
	}
}
