package basket

type Item struct {
	Sku         int
	Title       string
	Price       int
	Substituted bool
	SubPrice    int
}

func GetTotal(items []Item) int {
	return 0
}
