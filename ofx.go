package ofx

import (
	"encoding/xml"
	"io"

	"github.com/luizbranco/ofx/internal/paired"
)

type Date string //YYYYMMDD
type DateTime string
type Language string
type TransactionUID string
type CurrencySymbol string
type FITID string

type AccountType string

const (
	Checking     AccountType = "CHECKING"
	Savings                  = "SAVINGS"
	MoneyMarket              = "MONEYMRKT"
	LineOfCredit             = "CREDITLINE"
)

type OFX struct {
	SignOn     SignOn     `xml:"SIGNONMSGSRSV1"`
	Banking    Banking    `xml:"BANKMSGSRSV1,omitempty"`
	CreditCard CreditCard `xml:"CREDITCARDMSGSRSV1,omitempty"`
}

type Status struct {
	Code     StatusCode `xml:"CODE"`
	Severity Severity   `xml:"SEVERITY"`
	Message  string     `xml:"MESSAGE,omitempty"`
}

type Balance struct {
	Amount        float64  `xml:"BALAMT"`
	EffectiveDate DateTime `xml:"DTASOF"`
}

type BankTransactionsList struct {
	DateStart    Date          `xml:"DTSTART"`
	DateEnd      Date          `xml:"DTEND"`
	Transactions []Transaction `xml:"STMTTRN"`
}

func Parse(in io.Reader) (*OFX, error) {
	ofx := &OFX{}
	f, err := paired.EndTags(in)
	if err != nil {
		return ofx, err
	}
	err = xml.Unmarshal(f, &ofx)
	return ofx, err
}
