package parser

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadAxisMessageFromFile(fileName string) string {
	filePath := fmt.Sprintf("../resources/test/%s", fileName)
	message, err := ioutil.ReadFile(filePath)
	if err == nil {
		return string(message)
	}

	return ""
}

func TestAxisCreditMessage(t *testing.T) {
	axisAccount := AxisAccount{}
	message := ReadAxisMessageFromFile("axis.credit")
	Process(&axisAccount, message)

	assert.Equal(t, axisAccount.Name, "Axis Bank")
	assert.Equal(t, axisAccount.TransactionDetails.Type, CREDIT)
	assert.Equal(t, axisAccount.TransactionDetails.Amount, float64(1998))
	assert.Equal(t, axisAccount.TransactionDetails.Date, "2021-12-10")
}

func TestAxisDebitMessage(t *testing.T) {
	axisAccount := AxisAccount{}
	message := ReadAxisMessageFromFile("axis.debit")
	Process(&axisAccount, message)

	assert.Equal(t, axisAccount.Name, "Axis Bank")
	assert.Equal(t, axisAccount.TransactionDetails.Type, DEBIT)
	// slight gotcha in the golang float handling, but fortunately firefly takes care in rounding off
	assert.Equal(t, axisAccount.TransactionDetails.Amount, float64(10280.349609375))
	assert.Equal(t, axisAccount.TransactionDetails.Date, "2021-02-22")
}
