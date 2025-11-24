package add_test

import (
	"testing"

	"github.com/owlcode3/jma/add"
)

func TestAddNumbers(t *testing.T) {
	got := add.Add(200, 150)
	want := 350

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
