package ofx

import (
	"reflect"
	"testing"
	"time"
)

func TestDateTime_Time(t *testing.T) {
	egs := []struct {
		in  DateTime
		out string
	}{
		{"20151112120000", "12 Nov 15 12:00 UTC"},
	}

	for _, eg := range egs {
		want, err := time.Parse(time.RFC822, eg.out)
		if err != nil {
			t.Errorf("error %s", err)
		}
		got := eg.in.Time()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	}

}
