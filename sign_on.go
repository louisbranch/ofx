package ofx

type SignOn struct {
	SignOnResponse SignOnResponse `xml:"SONRS"`
}

type SignOnResponse struct {
	Status         Status   `xml:"STATUS"`
	DateTimeServer DateTime `xml:"DTSERVER"`
	Language       Language `xml:"LANGUAGE"`
}
