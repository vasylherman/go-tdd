package arrays

import (
	"fmt"
	"testing"
)

func ExampleSumArray() {
	numbers := [3]int{1, 2, 3}
	sum := SumArray(numbers)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkSum(b *testing.B) {
	numbers := [3]int{1, 2, 3}
	for b.Loop() {
		SumArray(numbers)
	}
}
