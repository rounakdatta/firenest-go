package parser

import (
	"testing"

	"github.com/rounakdatta/firenest/utils"
	"github.com/stretchr/testify/assert"
)

func TestHDFCCreditMessage(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	message := utils.ReadMessageFromFile("hdfc.credit")
	Process(&hdfcAccount, message)

	assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	assert.Equal(t, hdfcAccount.TransactionDetails.Type, utils.CREDIT)
	assert.Equal(t, hdfcAccount.TransactionDetails.Amount, float64(10567))
	assert.Equal(t, hdfcAccount.TransactionDetails.Date, "2021-01-10")
}

func TestHDFCDebitMessage(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	message := utils.ReadMessageFromFile("hdfc.debit")
	Process(&hdfcAccount, message)

	assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	assert.Equal(t, hdfcAccount.TransactionDetails.Type, utils.DEBIT)
	assert.Equal(t, hdfcAccount.TransactionDetails.Amount, float64(10000))
	assert.Equal(t, hdfcAccount.TransactionDetails.Date, "2021-05-12")
}
