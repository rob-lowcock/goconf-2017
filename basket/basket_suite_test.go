package basket_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBasket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basket Suite")
}
