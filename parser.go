package ofx

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func Parse(in io.Reader) (*OFX, error) {
	txt, err := endTags(in)
	if err != nil {
		return nil, err
	}

	ofx := &OFX{}
	err = xml.Unmarshal(txt, &ofx)
	return ofx, err
}

func endTags(in io.Reader) ([]byte, error) {
	var buf bytes.Buffer

	in, err := encode(in)
	if err != nil {
		return nil, err
	}

	dec := xml.NewDecoder(in)

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
				xml.Escape(&buf, []byte(s))
			}
			prev.content = s
		case xml.EndElement:
			n := t.Name.Local
			if prev.content != "" && prev.name != n {
				fmt.Fprintf(&buf, `</%s>`, prev.name)
				prev.content = ""
			}
			fmt.Fprintf(&buf, `</%s>`, n)
		}
	}

	return buf.Bytes(), nil
}

func encode(in io.Reader) (io.Reader, error) {
	tmp, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	t := charmap.Windows1252

	tmp, _, err = transform.Bytes(t.NewDecoder(), tmp)
	if err != nil {
		return nil, err
	}

	// sanitize replaces malformated XML with unescaped & => &amp;
	re := regexp.MustCompile(`&[^&]{0,4}`)
	tmp = re.ReplaceAllFunc(tmp, func(m []byte) []byte {
		if bytes.Equal(m, []byte("&amp;")) {
			return m
		}
		return bytes.Replace(m, []byte("&"), []byte("&amp;"), -1)
	})

	return bytes.NewReader(tmp), nil
}
