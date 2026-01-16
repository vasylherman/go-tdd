package arrays

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			fmt.Printf("want %d, got %d, given %#v", want, got, numbers)
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
