package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	const repeatCount = 5

	t.Run("string match", func(t *testing.T) {
		repeated := Repeat("a", repeatCount)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("character count", func(t *testing.T) {

		repeated := Repeat("a", repeatCount)
		count := strings.Count(repeated, "a")
		if count != repeatCount {
			t.Errorf("expected %d but got %d", count, repeatCount)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
