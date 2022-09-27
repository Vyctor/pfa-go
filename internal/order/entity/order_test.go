package entity_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vyctor/pfa-go/internal/order/entity"
	"testing"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{
		ID: "123",
	}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{
		ID:    "123",
		Price: 10,
	}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.NoError(t, err)

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, "123", order.Price)
	assert.Equal(t, "123", order.Tax)
}

func TestGivenAValidParams_WhenCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.NoError(t, err)

	err = entity.CalculateFinalPrice(order)

	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
