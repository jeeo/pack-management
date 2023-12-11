package model

type Order struct {
	ID       string
	Quantity int
	Packs    map[int]int
}

type OrderPack struct {
	Pack     Pack
	Quantity int
}

func MakeOrder(orderQuantity int, packs map[int]int) Order {
	if len(packs) == 0 {
		packs = map[int]int{}
	}

	order := Order{
		Quantity: orderQuantity,
		Packs:    packs,
	}

	return order
}

func OrderPackToPackSlice(storedPacks []Pack, orderPacks map[int]int) []OrderPack {
	result := make([]OrderPack, 0, len(orderPacks))
	for amount, quantity := range orderPacks {
		for _, pack := range storedPacks {
			if pack.Amount == amount {
				result = append(result, OrderPack{Pack: pack, Quantity: quantity})
			}
		}
	}
	return result
}
