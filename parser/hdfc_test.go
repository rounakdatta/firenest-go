package parser

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadHDFCMessageFromFile(fileName string) string {
	filePath := fmt.Sprintf("../resources/test/%s", fileName)
	message, err := ioutil.ReadFile(filePath)
	if err == nil {
		return string(message)
	}

	return ""
}

func TestHDFCCreditMessage(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	message := ReadHDFCMessageFromFile("hdfc.credit")
	Process(&hdfcAccount, message)

	assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	assert.Equal(t, hdfcAccount.TransactionDetails.Type, CREDIT)
	assert.Equal(t, hdfcAccount.TransactionDetails.Amount, float64(10567))
	assert.Equal(t, hdfcAccount.TransactionDetails.Date, "2021-01-10")
}

func TestHDFCDebitMessage(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	message := ReadHDFCMessageFromFile("hdfc.debit")
	Process(&hdfcAccount, message)

	assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	assert.Equal(t, hdfcAccount.TransactionDetails.Type, DEBIT)
	assert.Equal(t, hdfcAccount.TransactionDetails.Amount, float64(10000))
	assert.Equal(t, hdfcAccount.TransactionDetails.Date, "2021-05-12")
}
