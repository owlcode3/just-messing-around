package times_test

import (
	"testing"

	"github.com/owlcode3/jma/times"
)

func TestTimesNumbers(t *testing.T) {
	got := times.Times(200, 150)
	want := 300000

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
