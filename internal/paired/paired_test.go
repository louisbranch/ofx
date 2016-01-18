package paired

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestEndTags(t *testing.T) {
	egs := []bool{true, true}

	for i, ok := range egs {
		var buf bytes.Buffer
		in, _ := os.Open(path.Join("examples", fmt.Sprintf("in_%d.ofx", i)))
		out, _ := ioutil.ReadFile(path.Join("examples", fmt.Sprintf("out_%d.ofx", i)))

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
