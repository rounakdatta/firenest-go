package parser

type AssetAccountParser interface {
	setAccountName()
	setAccountId() error
	parseTransactionAmount() error
	parseTransactionDate() error
	parseTransactionDescription() error
}

type Transaction struct {
	Amount int
	Date string
	Description string
}

type AssetAccount struct {
	Name string
	Id int
	TransactionDetails Transaction
}

func (a *AssetAccount) setAccountId() error {
	// logic to fetch ID from db lies here
	a.Id = -1
	return nil
}

func Process(parser AssetAccountParser) {
	parser.setAccountName()
	parser.setAccountId()
	parser.parseTransactionAmount()
	parser.parseTransactionDate()
	parser.parseTransactionDescription()
}
