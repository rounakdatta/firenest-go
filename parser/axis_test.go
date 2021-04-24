package parser

import (
	"testing"

	"github.com/rounakdatta/firenest/utils"
	"github.com/stretchr/testify/assert"
)

func TestAxisCreditMessage(t *testing.T) {
	axisAccount := AxisAccount{}
	message := utils.ReadMessageFromFile("axis.credit")
	Process(&axisAccount, message)

	assert.Equal(t, axisAccount.Name, "Axis Bank")
	assert.Equal(t, axisAccount.TransactionDetails.Type, CREDIT)
	assert.Equal(t, axisAccount.TransactionDetails.Amount, float64(1998))
	assert.Equal(t, axisAccount.TransactionDetails.Date, "2021-12-10")
}

func TestAxisDebitMessage(t *testing.T) {
	axisAccount := AxisAccount{}
	message := utils.ReadMessageFromFile("axis.debit")
	Process(&axisAccount, message)

	assert.Equal(t, axisAccount.Name, "Axis Bank")
	assert.Equal(t, axisAccount.TransactionDetails.Type, DEBIT)
	// slight gotcha in the golang float handling, but fortunately firefly takes care in rounding off
	assert.Equal(t, axisAccount.TransactionDetails.Amount, float64(10280.349609375))
	assert.Equal(t, axisAccount.TransactionDetails.Date, "2021-02-22")
}
