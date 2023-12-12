package application

import (
	"context"
	"errors"
	"math"

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
		var quocient float32
		var packIndex int
		// if there is any quocient less than 1,
		// we should try to find the best local pack
		if isThereAnyQuocientLessThanOne(packsQuocients) {
			quocient, packIndex = bestLocalPack(packs, remainingQuantity)
		} else {
			// otherwise, we should find the best quocient (which means saving more packs)
			quocient, packIndex = minQuocient(packsQuocients)
		}
		packAmount := packs[packIndex].Amount
		packQuantity := int(quocient)
		if packQuantity == 0 {
			packQuantity = 1
		}
		order.Packs[packAmount] += packQuantity
		remainingQuantity -= packAmount * packQuantity
	}

	return order
}

// optimizeOrder tries to optimize the order by replacing the packs with the bigger ones
func (p OrderApplication) optimizeOrder(packs []model.Pack, order *model.Order) {
	packsMap := model.PacksToMap(packs)
	for packAmount, packQuantity := range order.Packs {
		for packQuantity > 1 {
			if _, ok := packsMap[packAmount*packQuantity]; ok {
				order.Packs[packAmount*packQuantity]++
				if order.Packs[packAmount]-packQuantity <= 0 {
					delete(order.Packs, packAmount)
				} else {
					order.Packs[packAmount] -= packQuantity
				}
				break
			}
			packQuantity--
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

func isThereAnyQuocientLessThanOne(quocients []float32) bool {
	for _, quocient := range quocients {
		if quocient < 1 {
			return true
		}
	}
	return false
}

func bestLocalPack(packs []model.Pack, orderQuantity int) (float32, int) {
	var bestIndex int
	var bestDiff int = math.MaxInt
	var resultQuantity int
	for i, pack := range packs {
		quocient := float32(orderQuantity) / float32(pack.Amount)
		if quocient < 1 {
			diff := pack.Amount - orderQuantity
			if bestDiff > diff {
				bestDiff = diff
				bestIndex = i
				resultQuantity = 1
			}
		} else {
			quantity := int(quocient)
			diff := orderQuantity - (pack.Amount * quantity)
			if bestDiff > diff {
				bestDiff = diff
				bestIndex = i
				resultQuantity = quantity
			}
		}
	}

	return float32(resultQuantity), bestIndex
}
