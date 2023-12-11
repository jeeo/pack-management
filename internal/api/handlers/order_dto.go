package handlers

import "github.com/jeeo/pack-management/internal/model"

type CalculateOrderPackRequest struct {
	Amount int `json:"amount"`
}

type CalculateOrderPackResponse struct {
	OrderPacks []OrderPackDTO `json:"order_packs"`
}

type OrderPackDTO struct {
	Pack     PackDTO `json:"pack"`
	Quantity int     `json:"quantity"`
}

func ToOrderPackDTOs(orderPacks []model.OrderPack) []OrderPackDTO {
	result := make([]OrderPackDTO, len(orderPacks))
	for i, orderPack := range orderPacks {
		result[i] = OrderPackDTO{
			Pack:     ToPackDTO(orderPack.Pack),
			Quantity: orderPack.Quantity,
		}
	}
	return result
}
