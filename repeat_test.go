package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatv2(t *testing.T) {
	repeated := Repeatv2("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// run benchmarks with `go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeatv2("a", 10)
	}
}

func ExampleRepeatv2() {
	repeat := Repeatv2("y", 16)
	fmt.Println(repeat)
	// Output: yyyyyyyyyyyyyyyy
}
