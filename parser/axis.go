package parser

type AxisAccount struct {
	AssetAccount
}

func (a *AxisAccount) setAccountName() {
	a.Name = "Axis"
}

func (a *AxisAccount) parseTransactionAmount() error {
	a.TransactionDetails.Amount = 500
	return nil
}

func (a *AxisAccount) parseTransactionDate() error {
	a.TransactionDetails.Date = "today"
	return nil
}

func (a *AxisAccount) parseTransactionDescription() error {
	a.TransactionDetails.Description = "amount y has been debited"
	return nil
}
