package ofx

type CreditCard struct {
	CreditCardResponse []CreditCardResponse `xml:"CCSTMTTRNRS"`
}

type CreditCardResponse struct {
	TransactionUID              TransactionUID              `xml:"TRNUID"`
	ClientCookie                string                      `xml:"CLTCOOKIE,omitempty"`
	Status                      Status                      `xml:"STATUS"`
	CreditCardStatementResponse CreditCardStatementResponse `xml:"CCSTMTRS"`
}

type CreditCardStatementResponse struct {
	CurrencyDefault      CurrencySymbol       `xml:"CURDEF"`
	CreditCardAccount    CreditCardAccount    `xml:"CCACCTFROM"`
	BankTransactionsList BankTransactionsList `xml:"BANKTRANLIST,omitempty"`
	LedgerBalance        Balance              `xml:"LEDGERBAL"`
	AvailableBalance     Balance              `xml:"AVAILBAL,omitempty"`
	MarketingInfo        string               `xml:"MKTGINFO,omitempty"`
}

type CreditCardAccount struct {
	ID  string `xml:"ACCTID"`
	Key string `xml:"ACCTKEY,omitempty"`
}
