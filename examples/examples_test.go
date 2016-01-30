package examples

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/luizbranco/ofx"
)

func TestParse(t *testing.T) {

	egs := []struct {
		file string
		res  *ofx.OFX
	}{
		{"ccstmtrs", CCStmtRs},
		{"desjardins", Desjardins},
	}

	for _, eg := range egs {
		in, err := os.Open(fmt.Sprintf("%s.ofx", eg.file))
		if err != nil {
			t.Fatal(err)
		}

		res, err := ofx.Parse(in)

		if err != nil {
			t.Errorf("Parse unexpected error file %s, %s", eg.file, err)
		}
		if !reflect.DeepEqual(eg.res, res) {
			t.Errorf("Parse expected file %s\n", eg.file)
			spew.Dump(eg.res, res)
		}
	}
}
