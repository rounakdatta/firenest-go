package parser

import (
	"errors"
	"regexp"

	"github.com/rounakdatta/firenest/database"
)

type AssetAccountParser interface {
	setAccountName()
	setAccountId() error
	parseTransactionType(string) error
	parseTransactionAmount(string) error
	parseTransactionDate(string) error
	parseTransactionDescription(string) error
	getAssetAccount() AssetAccount
}

type Direction uint32

const (
	DEBIT  Direction = 0
	CREDIT Direction = 1
)

type Transaction struct {
	Type        Direction
	Amount      float64
	Date        string
	Description string
}

type AssetAccount struct {
	Name               string
	Id                 int
	TransactionDetails Transaction
}

func (a *AssetAccount) setAccountId() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}
	a.Id = database.GetAccountIdFromName(db, a.Name)
	return nil
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

func (a *AssetAccount) getAssetAccount() AssetAccount {
	return *a
}

func Process(parser AssetAccountParser, message string) AssetAccount {
	parser.setAccountName()
	parser.setAccountId()
	parser.parseTransactionType(message)
	parser.parseTransactionAmount(message)
	parser.parseTransactionDate(message)
	parser.parseTransactionDescription(message)
	return parser.getAssetAccount()
}
