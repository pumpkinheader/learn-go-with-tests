package iteration // Package名はいったん気にしないことにする

import (
	"testing"
	"fmt"
)

func TestHasPrefix (t *testing.T) {
	got := HasPrefix("gopher", "go")
	want := true

	if got != want {
		t.Errorf("expected %t but got %t", want, got)
	}
}

func ExampleHasPrefix() {
	fmt.Println(HasPrefix("＼(・ω・＼)SAN値!(／・ω・)／ピンチ!", "＼(・ω・＼)"))
	fmt.Println(HasPrefix("＼(・ω・＼)SAN値!(／・ω・)／ピンチ!", "(」・ω・)」"))
	// Output:
	// true
	// false
}