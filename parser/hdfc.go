package parser

type HDFCAccount struct {
	AssetAccount
}

func (a *HDFCAccount) setAccountName() {
	a.Name = "HDFC"
}

func (a *HDFCAccount) parseTransactionAmount() error {
	a.TransactionDetails.Amount = 10000
	return nil
}

func (a *HDFCAccount) parseTransactionDate() error {
	a.TransactionDetails.Date = "yesterday"
	return nil
}

func (a *HDFCAccount) parseTransactionDescription() error {
	a.TransactionDetails.Description = "amount x has been credited"
	return nil
}
