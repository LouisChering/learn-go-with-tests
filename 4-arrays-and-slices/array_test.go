package arraysandslices_test

import (
	"reflect"
	"testing"

	"slices"

	arraysandslices "github.com/louischering/learn-go-with-tests/4-arrays-and-slices"
)

func Benchmark(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		arraysandslices.Sum(numbers)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum 5 numbers total to 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := arraysandslices.Sum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("sum 3 numbers total to 6", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := arraysandslices.Sum(numbers)
		expected := 6
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums two arrays", func(t *testing.T) {
		got := arraysandslices.SumAll([]int{1, 2, 3}, []int{3, 4, 5})
		expected := []int{6, 12}
		if !slices.Equal(got, expected) {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := arraysandslices.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := arraysandslices.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	output := arraysandslices.Sum(numbers)
	println(output)
	// Outputs: 15
}

func ExampleSumAll() {
	output := arraysandslices.SumAll([]int{1, 2, 3}, []int{2, 3, 4})
	println(output)
	// Outputs: {6,9}
}

func ExampleSumAllTails() {
	output := arraysandslices.SumAllTails([]int{1, 2, 3})
	println(output)
	// Outputs: {5}
}
