package ofx

import "encoding/xml"

type Banking struct {
	TransactionResponseResponse struct {
		XMLName        xml.Name `xml:"STMTTRNRS"`
		TransactionUID `xml:"TRNUID"`
		ClientCookie   string `xml:"CLTCOOKIE,omitempty"`
		Status         Status
		BankingAccount BankingAccount
	}
}

type BankStatementResponse struct {
	XMLName              xml.Name       `xml:"STMTRS"`
	CurrencyDefault      CurrencySymbol `xml:"CURDEF"`
	BankingAccount       BankingAccount
	BankTransactionsList BankTransactionsList `xml:"BANKTRANLIST,omitempty"`
	LedgerBalance        Balance              `xml:"LEDGERBAL"`
	AvailableBalance     Balance              `xml:"AVAILBAL,omitempty"`
	MarketingInfo        string               `xml:"MKTGINFO,omitempty"`
}

type BankingAccount struct {
	XMLName     xml.Name `xml:"BANKACCTFROM"`
	BankID      string   `xml:"BANKID"`
	BranchID    string   `xml:"BRANCHID,omitempty"`
	ID          string   `xml:"ACCTID"`
	AccountType `xml:"ACCTTYPE"`
	Key         string `xml:"ACCTKEY,omitempty"`
}
