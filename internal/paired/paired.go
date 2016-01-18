package paired

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func EndTags(in io.Reader) ([]byte, error) {
	dec := xml.NewDecoder(in)
	var buf bytes.Buffer
	var prev struct {
		name    string
		content string
	}
	for {
		t, err := dec.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch t := t.(type) {
		case xml.StartElement:
			if prev.name != "" && prev.content != "" {
				fmt.Fprintf(&buf, `</%s>`, prev.name)
			}
			n := t.Name.Local
			fmt.Fprintf(&buf, "<%s>", n)
			prev.name = n
			prev.content = ""
		case xml.CharData:
			s := strings.TrimSpace(string(t))
			if s != "" {
				fmt.Fprintf(&buf, "%s", s)
			}
			prev.content = s
		case xml.EndElement:
			n := t.Name.Local
			if prev.content != "" && prev.name != n {
				fmt.Fprintf(&buf, `</%s>`, prev.name)
			}
			fmt.Fprintf(&buf, `</%s>`, n)
		}
	}
	return buf.Bytes(), nil
}
