package basket_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rob-lowcock/goconf-2017/basket"
)

var _ = Describe("Basket", func() {
	It("returns a total of items", func() {
		output := basket.GetTotal([]basket.Item{
			{
				Sku:   1,
				Title: "Milk",
				Price: 50,
			},
			{
				Sku:   2,
				Title: "Bread",
				Price: 100,
			},
			{
				Sku:   3,
				Title: "Crisps",
				Price: 30,
			},
			{
				Sku:   4,
				Title: "Banana",
				Price: 20,
			},
		})

		Expect(output).To(Equal(200))
	})

	It("adds a cheaper substituted item for the cheaper price", func() {
		output := basket.GetTotal([]basket.Item{
			{
				Sku:   2,
				Title: "Bread",
				Price: 100,
			},
			{
				Sku:         2,
				Title:       "Bread",
				Price:       100,
				SubPrice:    50,
				Substituted: true,
			},
		})

		Expect(output).To(Equal(150))
	})

	It("adds a more expensive substituted item for the original price", func() {
		output := basket.GetTotal([]basket.Item{
			{
				Sku:   2,
				Title: "Bread",
				Price: 100,
			},
			{
				Sku:         1,
				Title:       "Milk",
				Price:       50,
				SubPrice:    100,
				Substituted: true,
			},
		})

		Expect(output).To(Equal(150))
	})

	It("can cope with a combination of scenarios in a single basket", func() {
		output := basket.GetTotal([]basket.Item{
			{
				Sku:   2,
				Title: "Bread",
				Price: 100,
			},
			{
				Sku:         1,
				Title:       "Milk",
				Price:       50,
				SubPrice:    100,
				Substituted: true,
			},
			{
				Sku:         5,
				Title:       "Chocolate",
				Price:       150,
				SubPrice:    75,
				Substituted: true,
			},
			{
				Sku:         3,
				Title:       "Banana",
				Price:       20,
				SubPrice:    30,
				Substituted: true,
			},
		})

		Expect(output).To(Equal(245))
	})
})
