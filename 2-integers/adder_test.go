package integers_test

import (
	"fmt"
	"testing"

	integers "github.com/louischering/learn-go-with-tests/2-integers"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
