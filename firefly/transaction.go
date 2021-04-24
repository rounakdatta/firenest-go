package firefly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/rounakdatta/firenest/database"
	"github.com/rounakdatta/firenest/parser"
)

type Transaction struct {
	Type              string `json:"type"`
	Date              string `json:"date"`
	Amount            string `json:"amount"`
	Description       string `json:"description"`
	SourceId          int    `json:"source_id"`
	SourceName        string `json:"source_name"`
	DestinationId     int    `json:"destination_id"`
	DestinationName   string `json:"destination_name"`
	CategoryName      string `json:"category_name" default:""`
	InterestDate      string `json:"interest_date" default:""`
	BookDate          string `json:"book_date" default:""`
	ProcessDate       string `json:"process_date" default:""`
	DueDate           string `json:"due_date" default:""`
	PaymentDate       string `json:"payment_date" default:""`
	InvoiceDate       string `json:"invoice_date" default:""`
	InternalReference string `json:"internal_reference" default:""`
	Notes             string `json:"notes" default:""`
}

func transformType(direction parser.Direction) string {
	switch direction {
	case parser.CREDIT:
		return "deposit"
	case parser.DEBIT:
		return "withdrawal"
	}

	return ""
}

func getDefaultAccount() (int, string) {
	defaultAccountName := "(no name)"

	db, _ := database.GetConnection()
	return database.GetAccountIdFromName(db, defaultAccountName), defaultAccountName
}

func getSource(account parser.AssetAccount) (int, string) {
	if account.TransactionDetails.Type == parser.DEBIT {
		return account.Id, account.Name
	}

	accountId, accountName := getDefaultAccount()
	return accountId, accountName
}

func getDestination(account parser.AssetAccount) (int, string) {
	if account.TransactionDetails.Type == parser.CREDIT {
		return account.Id, account.Name
	}

	accountId, accountName := getDefaultAccount()
	return accountId, accountName
}

func ParseMessage(message string, sender string) parser.AssetAccount {
	hdfcRegex := regexp.MustCompile(`(?i)HDFCBK`)
	axisRegex := regexp.MustCompile(`(?i)AxisBk`)

	var processor parser.AssetAccountParser
	if hdfcRegex.MatchString(sender) {
		processor = &parser.HDFCAccount{}
	} else if axisRegex.MatchString(sender) {
		processor = &parser.AxisAccount{}
	} else {
		return parser.AssetAccount{}
	}

	return parser.Process(processor, message)
}

func CreateTransaction(account parser.AssetAccount) Transaction {
	sourceAccountId, sourceAccountName := getSource(account)
	destinationAccountId, destinationAccountName := getDestination(account)

	return Transaction{
		Type:            transformType(account.TransactionDetails.Type),
		Date:            account.TransactionDetails.Date,
		Amount:          fmt.Sprintf("%f", account.TransactionDetails.Amount),
		Description:     account.TransactionDetails.Description,
		SourceId:        sourceAccountId,
		SourceName:      sourceAccountName,
		DestinationId:   destinationAccountId,
		DestinationName: destinationAccountName,
		Notes:           account.TransactionDetails.Description,
	}
}

func SendFireflyRequest(method string, url string, transaction Transaction, headers map[string]string) error {
	payload, err := json.Marshal(transaction)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	for header_key, header_value := range headers {
		req.Header.Add(header_key, header_value)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
