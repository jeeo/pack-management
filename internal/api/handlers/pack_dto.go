package handlers

import "github.com/jeeo/pack-management/internal/model"

type CreatePackRequest struct {
	Amount int `json:"amount"`
}

type CreatePackResponse struct {
	Pack PackDTO `json:"pack"`
}

type PackDTO struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

func ToPackDTO(pack model.Pack) PackDTO {
	return PackDTO{
		ID:     pack.ID,
		Amount: pack.Amount,
	}
}

type UpdatePackRequest struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type UpdatePackResponse struct {
	Pack PackDTO `json:"pack"`
}

type FindAllPackResponse struct {
	Packs []PackDTO `json:"packs"`
}
