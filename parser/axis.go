package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AxisAccount struct {
	AssetAccount
}

func (a *AxisAccount) setAccountId() error {
	// logic to fetch ID from db lies here
	a.Id = -1
	return nil
}

func (a *AxisAccount) setAccountName() {
	a.Name = "Axis Bank"
}

func (a *AxisAccount) parseTransactionAmount(message string) error {
	amountRegex := regexp.MustCompile(`(?:(?:RS|Rs|INR|MRP)\.?\s?)(\d+(:?\,\d+)?(\,\d+)?(\.\d{1,2})?)`)
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

func (a *AxisAccount) parseTransactionDate(message string) error {
	dateRegex := regexp.MustCompile(`(?:(?:on)\.?\s?)(\d+?-(\d+?|\w{3})-\d+)`)
	bankDateFormat1 := "02-01-06"
	bankDateFormat2 := "02-01-2006"

	extractedDate := dateRegex.FindString(message)
	extractedDateSplit := strings.Split(extractedDate, " ")
	if len(extractedDateSplit) > 1 {
		extractedDate = extractedDateSplit[1]
	} else {
		extractedDate = extractedDateSplit[0]
	}
	standardDateFormat, err := time.Parse(bankDateFormat1, extractedDate)
	if err != nil {
		standardDateFormat, err = time.Parse(bankDateFormat2, extractedDate)
	}
	if err != nil {
		return err
	}

	expectedDateFormat := standardDateFormat.Format("2006-01-02")
	a.TransactionDetails.Date = expectedDateFormat
	return nil
}

func (a *AxisAccount) parseTransactionDescription(message string) error {
	var descriptionRegex *regexp.Regexp
	if a.TransactionDetails.Type == DEBIT {
		descriptionRegex = regexp.MustCompile(`(?:(?:at)\.?\s?)(.+?(?P<desc>\.))`)
	} else {
		descriptionRegex = regexp.MustCompile(`(?:(?:to\ [^a|A]|for|Info)\s?)(.+?(?P<desc>\.))`)
	}

	a.TransactionDetails.Description = descriptionRegex.FindString(message)
	return nil
}
