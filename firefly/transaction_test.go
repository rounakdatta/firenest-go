package firefly

import (
	"testing"

	"github.com/rounakdatta/firenest/parser"
	"github.com/rounakdatta/firenest/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseMessage(t *testing.T) {
	message := utils.ReadMessageFromFile("hdfc.credit")
	processor := ParseMessage(message, "VM-HDFCBK")

	assert.Equal(t, processor.Name, "HDFC Bank")
	assert.Equal(t, processor.TransactionDetails.Amount, float64(10567))
	assert.Equal(t, processor.TransactionDetails.Type, parser.CREDIT)
}

func TestCreateTransaction(t *testing.T) {
	message := utils.ReadMessageFromFile("hdfc.credit")
	processor := ParseMessage(message, "VM-HDFCBK")
	transaction := CreateTransaction(processor)

	assert.Equal(t, transaction.Type, "deposit")
	assert.Equal(t, transaction.Date, "2021-01-10")
	assert.Equal(t, transaction.DestinationName, "HDFC Bank")
	assert.Equal(t, transaction.SourceName, "(no name)")
}

func TestGetEndpoint(t *testing.T) {
	defaultEndpoint := GetEndpoint()

	assert.Equal(t, defaultEndpoint, "http://localhost/money/api/v1/transactions")
}
