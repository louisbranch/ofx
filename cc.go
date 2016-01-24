package ofx

import "encoding/xml"

type CreditCard struct {
	TransactionResponseResponse struct {
		XMLName                     xml.Name `xml:"CCSTMTTRNRS"`
		TransactionUID              `xml:"TRNUID"`
		ClientCookie                string `xml:"CLTCOOKIE,omitempty"`
		Status                      Status
		CreditCardStatementResponse CreditCardStatementResponse
	}
}

type CreditCardStatementResponse struct {
	XMLName              xml.Name       `xml:"CCSTMTRS"`
	CurrencyDefault      CurrencySymbol `xml:"CURDEF"`
	CreditCardAccount    CreditCardAccount
	BankTransactionsList BankTransactionsList `xml:"BANKTRANLIST,omitempty"`
	LedgerBalance        Balance              `xml:"LEDGERBAL"`
	AvailableBalance     Balance              `xml:"AVAILBAL,omitempty"`
	MarketingInfo        string               `xml:"MKTGINFO,omitempty"`
}

type CreditCardAccount struct {
	XMLName xml.Name `xml:"CCACCTFROM"`
	ID      string   `xml:"ACCTID"`
	Key     string   `xml:"ACCTKEY,omitempty"`
}
