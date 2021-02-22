package parser

import (
	"regexp"
	"strconv"
	"strings"
)

type HDFCAccount struct {
	AssetAccount
}

func (a *HDFCAccount) setAccountId() error {
	// logic to fetch ID from db lies here
	a.Id = -1
	return nil
}

func (a *HDFCAccount) setAccountName() {
	a.Name = "HDFC Bank"
}

func (a *HDFCAccount) parseTransactionAmount(message string) error {
	amountRegex := regexp.MustCompile("(?:(?:RS|Rs|INR|MRP)\\.?\\s?)(\\d+(:?\\,\\d+)?(\\,\\d+)?(\\.\\d{1,2})?)")
	amountStringified := amountRegex.FindString(message)

	if len(amountStringified) != 0 {
		amountSanitised := strings.ReplaceAll(amountStringified, ",", "")
		currencyValueSplit := strings.Split(amountSanitised, " ")
		_ = currencyValueSplit[0]
		amount, err := strconv.ParseFloat(currencyValueSplit[1], 32)
		if err == nil {
			a.TransactionDetails.Amount = amount
			return nil
		}
	}

	a.TransactionDetails.Amount = 0
	return nil
}

func (a *HDFCAccount) parseTransactionDate(message string) error {
	a.TransactionDetails.Date = "yesterday"
	return nil
}

func (a *HDFCAccount) parseTransactionDescription(message string) error {
	a.TransactionDetails.Description = "amount x has been credited"
	return nil
}
