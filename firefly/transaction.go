package firefly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rounakdatta/firenest/utils"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
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

type Payload struct {
	Transactions []Transaction `json:"transactions"`
}

func transformType(direction utils.Direction) string {
	switch direction {
	case utils.CREDIT:
		return "deposit"
	case utils.DEBIT:
		return "withdrawal"
	}

	return ""
}

func getDefaultAccount(direction utils.Direction) (int, string) {
	defaultAccountName := "(no name)"

	db, err := database.GetConnection()
	if err != nil {
		return -1, defaultAccountName
	}
	return database.GetAccountIdFromName(db, defaultAccountName, direction), defaultAccountName
}

func getSource(account parser.AssetAccount) (int, string) {
	if account.TransactionDetails.Type == utils.DEBIT {
		return account.Id, account.Name
	}

	accountId, accountName := getDefaultAccount(account.TransactionDetails.Type)
	return accountId, accountName
}

func getDestination(account parser.AssetAccount) (int, string) {
	if account.TransactionDetails.Type == utils.CREDIT {
		return account.Id, account.Name
	}

	accountId, accountName := getDefaultAccount(account.TransactionDetails.Type)
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

func CreateTransaction(account parser.AssetAccount) Payload {
	sourceAccountId, sourceAccountName := getSource(account)
	destinationAccountId, destinationAccountName := getDestination(account)

	transaction := Transaction{
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

	return Payload{
		Transactions: []Transaction{transaction},
	}
}

func GetEndpoint() string {
	err := godotenv.Load()
	host := "http://localhost"
	if err == nil {
		host = os.Getenv("FIREFLY_ENDPOINT")
	}

	fireflyEndpoint := fmt.Sprintf("%s/money/api/v1/transactions", host)
	return fireflyEndpoint
}

func GetPersonalAccessToken() string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}

	return os.Getenv("FIREFLY_PAT")
}

func SendFireflyRequest(method string, url string, transaction Payload, headers map[string]string) error {
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
