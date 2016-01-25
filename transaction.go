package ofx

type TransactionType string

const (
	Credit           TransactionType = "CREDIT"
	Debit                            = "DEBIT"
	Interest                         = "INT"
	Dividend                         = "DIVIDEND"
	Fee                              = "FEE"
	ServiceCharge                    = "SRVCHG"
	Deposit                          = "DEP"
	ATM                              = "ATM"
	PointOfSale                      = "POS"
	Transfer                         = "XFER"
	Check                            = "CHECK"
	EletronicPayment                 = "PAYMENT"
	CashWithdrawal                   = "CASH"
	DirectDeposit                    = "DIRECTDEP"
	DirectDebit                      = "DIRECTDEBIT"
	RepeatingPayment                 = "REPEATPMT"
	Other                            = "OTHER"
)

type CorrectAction string

const (
	Replace CorrectAction = "REPLACE"
	Delete  CorrectAction = "DELETE"
)

type Transaction struct {
	TransactionType   TransactionType `xml:"TRNTYPE"`
	DatePosted        DateTime        `xml:"DTPOSTED"`
	DateUser          DateTime        `xml:"DTUSER,omitempty"`
	DateAvailable     DateTime        `xml:"DTAVAIL,omitempty"`
	TransactionAmount float64         `xml:"TRNAMT"`
	FITID             FITID           `xml:"FITID"`
	CorrectFITID      FITID           `xml:"CORRECTFITID,omitempty"`
	CorrectAction     CorrectAction   `xml:"CORRECTACTION,omitempty"`
	ServerID          string          `xml:"SRVRID,omitempty"`
	CheckNumber       string          `xml:"CHECKNUM,omitempty"`
	ReferenceNumber   string          `xml:"REFNUM,omitempty"`
	StdIndustrialCode string          `xml:"SIC"`
	PayeeID           string          `xml:"PAYEEID,omitempty"`
	Name              string          `xml:"NAME"`
	Memo              string          `xml:"MEMO,omitempty"`
	/*
		TODO
			PAYEE
			BANKACCTTO
			CCACCTTO
			CURRENCY
			ORIFINALCURRENCY
	*/
}
