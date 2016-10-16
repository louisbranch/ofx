package ofx

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type Date string //YYYYMMDD
type DateTime string
type Language string
type TransactionUID string
type Currency string
type FITID string
type Amount float64

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
	Code     int      `xml:"CODE"`
	Severity Severity `xml:"SEVERITY"`
	Message  string   `xml:"MESSAGE,omitempty"`
}

type Balance struct {
	Amount        Amount   `xml:"BALAMT"`
	EffectiveDate DateTime `xml:"DTASOF"`
}

type BankTransactionsList struct {
	DateStart    Date          `xml:"DTSTART"`
	DateEnd      Date          `xml:"DTEND"`
	Transactions []Transaction `xml:"STMTTRN"`
}

// UnmarshalXML tries to unmarshal the amount using a decimal point or comma
// separator
func (a *Amount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*a = Amount(v)
		return nil
	}
	s = strings.Replace(s, ",", ".", 1)
	v, nerr := strconv.ParseFloat(s, 64)
	if nerr == nil {
		*a = Amount(v)
		return nil
	}
	return err
}

func (d DateTime) Time() time.Time {
	t, _ := time.Parse("20060102150505", string(d))
	return t
}
