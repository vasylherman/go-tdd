package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("multiple collections of any size", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !slices.Equal(want, got) {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("want %d, got %d, given %#v", want, got, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 6}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 9
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 6}
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{3, 3, 3})
	}
}
