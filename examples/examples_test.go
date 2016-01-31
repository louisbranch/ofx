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
		{"ccstmtrs.ofx", CCStmtRs},
		{"desjardins.ofx", Desjardins},
		{"response.ofx", Response},
		{"ing.qfx", Ing},
		{"fortis.ofx", Fortis},
	}

	for _, eg := range egs {
		in, err := os.Open(fmt.Sprintf("%s", eg.file))
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
