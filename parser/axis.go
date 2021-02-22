package parser

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
	a.TransactionDetails.Amount = 500
	return nil
}

func (a *AxisAccount) parseTransactionDate(message string) error {
	a.TransactionDetails.Date = "today"
	return nil
}

func (a *AxisAccount) parseTransactionDescription(message string) error {
	a.TransactionDetails.Description = "amount y has been debited"
	return nil
}
