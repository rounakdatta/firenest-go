package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"
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
	dateRegex := regexp.MustCompile("(?:(?:on)\\.?\\s?)(\\d+?-(\\d+?|\\w{3})-\\d+)")
	bankDateFormat := "02-Jan-06"

	extractedDate := dateRegex.FindString(message)
	extractedDateSplit := strings.Split(extractedDate, " ")
	if len(extractedDateSplit) > 1 {
		extractedDate = extractedDateSplit[1]
	} else {
		extractedDate = extractedDateSplit[0]
	}
	standardDateFormat, err := time.Parse(bankDateFormat, extractedDate)
	if err != nil {
		return err
	}

	expectedDateFormat := standardDateFormat.Format("2006-01-02")
	a.TransactionDetails.Date = expectedDateFormat
	return nil
}

func (a *HDFCAccount) parseTransactionDescription(message string) error {
	var descriptionRegex *regexp.Regexp
	if a.TransactionDetails.Type == DEBIT {
		descriptionRegex = regexp.MustCompile(`(?:(?:to)\.?\s?)(.+?(?P<desc>\.))`)
	} else {
		descriptionRegex = regexp.MustCompile(`(?:(?:to\ [^a|A]|for)\s?)(.+?(?P<desc>\.))`)
	}

	a.TransactionDetails.Description = descriptionRegex.FindString(message)
	return nil
}
