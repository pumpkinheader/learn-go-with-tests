package iteration // Package名はいったん気にしないことにする

import (
	"testing"
)

func TestHasPrefix (t *testing.T) {
	got := HasPrefix("gopher", "go")
	want := true

	if got != want {
		t.Errorf("expected %t but got %t", want, got)
	}
}