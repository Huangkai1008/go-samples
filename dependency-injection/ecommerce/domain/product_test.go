package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("apply discount when user is preferred customer", func(t *testing.T) {
		productUnitPrice := 100.0
		expectedUnitPrice := 100.0
		product := &Product{UnitPrice: productUnitPrice}

		discount := product.ApplyDiscountFor(&StubUserContext{})

		assert.Equal(t, expectedUnitPrice, discount.UnitPrice)
	})

	t.Run("not apply discount when user is not preferred customer", func(t *testing.T) {
		productUnitPrice := 100.0
		expectedUnitPrice := 95.0
		preferredCustomer := &StubUserContext{[]Role{PreferredCustomer}}
		product := &Product{UnitPrice: productUnitPrice}

		discount := product.ApplyDiscountFor(preferredCustomer)

		assert.Equal(t, expectedUnitPrice, discount.UnitPrice)
	})
}
