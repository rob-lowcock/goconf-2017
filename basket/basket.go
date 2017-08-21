package basket

import (
	"math"
)

type Item struct {
	Sku         int
	Title       string
	Price       int
	Substituted bool
	SubPrice    int
}

func GetTotal(items []Item) int {
	var total = 0
	for _, item := range items {
		if item.Substituted {
			total += int(math.Min(float64(item.Price), float64(item.SubPrice)))
		} else {
			if item.SubPrice < item.Price {
				total += item.Price
			} else {
				total += item.Price
			}
		}
	}
	return total
}
