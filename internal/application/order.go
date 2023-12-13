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
	calcTable := make([]int, orderQuantity+1)

	if orderQuantity < packs[0].Amount {
		order.Packs = map[int]int{packs[0].Amount: 1}
		return order
	}

	for i := range calcTable {
		calcTable[i] = orderQuantity + 1
	}
	calcTable[0] = 0

	for _, pack := range packs {
		for i := pack.Amount; i <= orderQuantity; i++ {
			calcTable[i] = min(calcTable[i], calcTable[i-pack.Amount]+1)
		}
	}

	remainingQuantity := orderQuantity
	if calcTable[orderQuantity] > orderQuantity {
		for remainingQuantity > 0 && calcTable[remainingQuantity] > orderQuantity {
			order.Packs[packs[0].Amount]++
			remainingQuantity -= packs[0].Amount
		}
	}
	for remainingQuantity > 0 {
		for _, pack := range packs {
			if pack.Amount <= remainingQuantity && calcTable[remainingQuantity] == calcTable[remainingQuantity-pack.Amount]+1 {
				order.Packs[pack.Amount]++
				remainingQuantity -= pack.Amount
				break
			}
		}
	}
	return order
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// optimizeOrder tries to optimize the order by replacing the packs with the bigger ones
func (p OrderApplication) optimizeOrder(packs []model.Pack, order *model.Order) {
	packsMap := model.PacksToMap(packs)
	shouldRecheck := true
	for shouldRecheck {
		shouldRecheck = false
		for packAmount, packQuantity := range order.Packs {
			for packQuantity > 1 {
				if _, ok := packsMap[packAmount*packQuantity]; ok {
					order.Packs[packAmount*packQuantity]++
					if order.Packs[packAmount]-packQuantity <= 0 {
						delete(order.Packs, packAmount)
					} else {
						order.Packs[packAmount] -= packQuantity
						if order.Packs[packAmount] > 1 {
							shouldRecheck = true
						}
					}
					break
				} else {
					bestPackAmount := 0
					sum := sumPackMap(order.Packs)
					for _, pack := range packs {
						if pack.Amount == packAmount {
							continue
						}
						if pack.Amount < packAmount*packQuantity && pack.Amount > bestPackAmount && pack.Amount < sum && pack.Amount >= order.Quantity {
							bestPackAmount = pack.Amount
						}
					}
					if bestPackAmount != 0 {
						order.Packs[bestPackAmount]++
						order.Packs[packAmount] -= packQuantity
						if order.Packs[packAmount] == 0 {
							delete(order.Packs, packAmount)
						}
						if order.Packs[packAmount] > 1 {
							shouldRecheck = true
						}
						break
					}
				}
				packQuantity--
			}
		}
	}
}

func sumPackMap(packs map[int]int) int {
	sum := 0
	for packAmount, packQuantity := range packs {
		sum += packAmount * packQuantity
	}
	return sum
}
