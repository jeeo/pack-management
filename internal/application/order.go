package application

import (
	"context"
	"errors"

	"github.com/jeeo/pack-management/internal/model"
)

type OrderApplication struct {
	packRepo packRepository
}

func NewOrderApplication(packRepo packRepository) OrderApplication {
	return OrderApplication{
		packRepo: packRepo,
	}
}

func (p OrderApplication) CalculateOrderPack(ctx context.Context, orderQuantity int) ([]model.OrderPack, error) {
	packs, err := p.packRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	if orderQuantity <= 0 {
		return nil, errors.New("order quantity must be greater than 0")
	}

	if len(packs) == 0 {
		return nil, errors.New("no packs found")
	}

	order := p.buildOrder(orderQuantity, packs)

	p.optimizeOrder(packs, &order)

	return model.OrderPackToPackSlice(packs, order.Packs), nil
}

func (p OrderApplication) buildOrder(orderQuantity int, packs []model.Pack) model.Order {
	order := model.MakeOrder(orderQuantity, nil)
	remainingQuantity := orderQuantity

	for remainingQuantity > 0 {
		if remainingQuantity < packs[0].Amount {
			if _, ok := order.Packs[packs[0].Amount]; !ok {
				order.Packs[packs[0].Amount] = 1
			} else {
				order.Packs[packs[0].Amount]++
			}

			return order
		}
		packsQuocients := calculatePacksQuocients(remainingQuantity, packs)
		filteredQuocients := filterQuocientsLessThanOne(packsQuocients)
		minimumQuocient, packIndex := minQuocient(filteredQuocients)
		if minimumQuocient > 1 {
			packAmount := packs[packIndex].Amount
			packQuantity := int(minimumQuocient)
			order.Packs[packAmount] = int(packQuantity)
			remainingQuantity -= packAmount * packQuantity
		} else if minimumQuocient == 1 || minimumQuocient < 1 && remainingQuantity < packs[packIndex].Amount {
			order.Packs[packs[packIndex].Amount] = 1
			return order
		}
	}

	return order
}

func (p OrderApplication) optimizeOrder(packs []model.Pack, order *model.Order) {
	packsMap := model.PacksToMap(packs)
	shouldRecheck := true
	for shouldRecheck {
		for packAmount, packQuantity := range order.Packs {
			if packQuantity > 1 {
				if _, ok := packsMap[packAmount*packQuantity]; ok {
					delete(order.Packs, packAmount)
					if _, ok := order.Packs[packAmount*packQuantity]; !ok {
						order.Packs[packAmount*packQuantity] = 1
					} else {
						order.Packs[packAmount*packQuantity]++
						shouldRecheck = true
						break
					}
				}
			}
			shouldRecheck = false
		}
	}
}

func calculatePacksQuocients(quantity int, packs []model.Pack) []float32 {
	result := make([]float32, len(packs))

	for i, pack := range packs {
		result[i] = float32(quantity) / float32(pack.Amount)
	}

	return result
}

func minQuocient(quocients []float32) (float32, int) {
	min := quocients[0]
	var index int
	for i, quocient := range quocients {
		if quocient < min {
			min = quocient
			index = i
		}
	}
	return min, index
}

func filterQuocientsLessThanOne(quocients []float32) []float32 {
	result := make([]float32, 0)
	for _, quocient := range quocients {
		if quocient >= 1 {
			result = append(result, quocient)
		}
	}
	return result
}
