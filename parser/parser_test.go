package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractClass(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	Process(&hdfcAccount, "Just a blank SMS")

	assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	assert.Equal(t, hdfcAccount.TransactionDetails.Type, DEBIT)
}
