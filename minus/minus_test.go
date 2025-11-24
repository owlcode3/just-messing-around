package minus_test

import (
	"testing"

	"github.com/owlcode3/jma/minus"
)

func TestMinusNumbers(t *testing.T) {
	got := minus.Minus(800, 700)
	want := 100

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
