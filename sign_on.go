package ofx

import "encoding/xml"

type SignOn struct {
	Response struct {
		XMLName        xml.Name `xml:"SONRS"`
		Status         Status
		DateTimeServer DateTime `xml:"DTSERVER"`
		Language       Language `xml:"LANGUAGE"`
	}
}
