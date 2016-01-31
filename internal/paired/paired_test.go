package paired

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func TestEndTags(t *testing.T) {
	egs := []bool{true, true, true}

	for i, ok := range egs {
		var buf bytes.Buffer
		in, _ := os.Open(path.Join("examples", fmt.Sprintf("in_%d.ofx", i)))
		f, _ := os.Open(path.Join("examples", fmt.Sprintf("out_%d.ofx", i)))

		tr := charmap.Windows1252
		r := transform.NewReader(f, tr.NewDecoder())
		out, _ := ioutil.ReadAll(r)

		for _, l := range bytes.Split(out, []byte("\n")) {
			buf.Write(bytes.TrimSpace(l))
		}
		out = buf.Bytes()

		res, err := EndTags(in)

		if ok {
			if err != nil {
				t.Errorf("EndTags unexpected error file %d, %s", i, err)
			}
			if !bytes.Equal(out, res) {
				t.Errorf("EndTags expected file %d\n%s\ngot\n%s", i, out, res)
			}
		} else {
			if err == nil {
				t.Errorf("EndTags expected error file %d, got none", i)
			}
		}
	}
}
