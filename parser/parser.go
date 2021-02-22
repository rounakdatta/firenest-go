package parser

type AssetAccountParser interface {
	setAccountName()
	setAccountId() error
	parseTransactionAmount(string) error
	parseTransactionDate(string) error
	parseTransactionDescription(string) error
}

type Transaction struct {
	Amount float64
	Date string
	Description string
}

type AssetAccount struct {
	Name string
	Id int
	TransactionDetails Transaction
}

func Process(parser AssetAccountParser, message string) {
	parser.setAccountName()
	parser.setAccountId()
	parser.parseTransactionAmount(message)
	parser.parseTransactionDate(message)
	parser.parseTransactionDescription(message)
}
