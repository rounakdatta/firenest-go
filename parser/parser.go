package parser

import (
	"errors"
	"regexp"
)

type AssetAccountParser interface {
	setAccountName()
	setAccountId() error
	parseTransactionType(string) error
	parseTransactionAmount(string) error
	parseTransactionDate(string) error
	parseTransactionDescription(string) error
}

type Direction uint32
const (
	DEBIT Direction = 0
	CREDIT Direction = 1
)

type Transaction struct {
	Type Direction
	Amount float64
	Date string
	Description string
}

type AssetAccount struct {
	Name string
	Id int
	TransactionDetails Transaction
}

func (a *AssetAccount) parseTransactionType(message string) error {
	debitRegex := regexp.MustCompile(`(?:(?:debited))`)
	creditRegex := regexp.MustCompile(`(?:(?:credited|deposited))`)

	if debitRegex.MatchString(message) {
		a.TransactionDetails.Type = DEBIT
	} else if creditRegex.MatchString(message) {
		a.TransactionDetails.Type = CREDIT
	} else {
		return errors.New("Couldn't understand transaction type")
	}

	return nil
}

func Process(parser AssetAccountParser, message string) {
	parser.setAccountName()
	parser.setAccountId()
	parser.parseTransactionType(message)
	parser.parseTransactionAmount(message)
	parser.parseTransactionDate(message)
	parser.parseTransactionDescription(message)
}
