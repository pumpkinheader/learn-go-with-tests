package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat (b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	got := Repeat("＼(・ω・＼)SAN値!(／・ω・)／ピンチ!", 7)
	fmt.Println(got)
	// Output: ＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!＼(・ω・＼)SAN値!(／・ω・)／ピンチ!

}