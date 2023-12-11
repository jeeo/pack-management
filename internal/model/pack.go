package model

// Pack describe a item package
// for simplicity it does not contain a reference to an item
// which is strongly applyable.
type Pack struct {
	ID     string
	Amount int
}

func PacksToMap(packs []Pack) map[int]struct{} {
	result := map[int]struct{}{}
	for _, pack := range packs {
		result[pack.Amount] = struct{}{}
	}
	return result
}

func PacksMapToSlice(storedPacks []Pack, orderPacks map[int]struct{}) []Pack {
	result := make([]Pack, len(orderPacks))
	for amount := range orderPacks {
		for _, pack := range storedPacks {
			if pack.Amount == amount {
				result = append(result, pack)
			}
		}
	}
	return result
}
