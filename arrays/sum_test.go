package arrays

import (
	"fmt"
	"testing"
)

func TestSumArray(t *testing.T) {
	numbers := [...]int{1, 2, 3}
	got := SumArray(numbers)
	want := 6
	if want != got {
		fmt.Printf("want %d, got %d, given %#v", want, got, numbers)
	}
}

func ExampleSumArray() {
	numbers := [...]int{1, 2, 3}
	sum := SumArray(numbers)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	numbers := [...]int{1, 2, 3}
	for b.Loop() {
		SumArray(numbers)
	}
}
