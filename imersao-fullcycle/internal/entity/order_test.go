package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsErrorIfIdIsBlank(t *testing.T) {
	// order, err := NewOrder("", 1, 1)
	// if err != nil {
	// 	t.Error("Error while create order")
	// }

	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfItGetsErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.Validate(), "price")
}

func TestIfItGetsErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "1", Price: 1}
	assert.Error(t, order.Validate(), "tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "1", Price: 10, Tax: 1}
	assert.NoError(t, order.Validate())
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, "1", order.ID)
}
