package ofx

type SignOn struct {
	SignOnResponse SignOnResponse `xml:"SONRS"`
}

type SignOnResponse struct {
	Status               Status               `xml:"STATUS"`
	DateTimeServer       DateTime             `xml:"DTSERVER"`
	Language             Language             `xml:"LANGUAGE"`
	DateAccountUp        DateTime             `xml:"DTACCTUP,omitempty"`
	FinancialInstitution FinancialInstitution `xml:"FI,omitempty"`
}

type FinancialInstitution struct {
	Organization string `xml:"ORG"`
	ID           string `xml:"FID"`
}
