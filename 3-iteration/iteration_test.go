package iteration_test

import (
	"testing"

	iteration "github.com/louischering/learn-go-with-tests/3-iteration"
)

func ExampleRepeat() {
	repeated := iteration.Repeat("a", 5)
	print(repeated)
	// Outputs: "aaaaa"
}

func TestRepeat(t *testing.T) {
	t.Run("repeat string 5 times", func(t *testing.T) {
		repeated := iteration.Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeat string 3 times", func(t *testing.T) {
		repeated := iteration.Repeat("a", 3)
		expected := "aaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeat string 3 times", func(t *testing.T) {
		repeated := iteration.Repeat("a", 0)
		expected := ""
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 5)
	}
}
