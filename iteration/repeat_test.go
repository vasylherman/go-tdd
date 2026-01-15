package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat(`ç`, 10)
	fmt.Println(repeated)
	// Output: çççççççççç
}

// Run go test ./iteration -bench=. to execute BenchmarkRepeat
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("√", 50)
	}
}

func BenchmarkRepeatRune(b *testing.B) {
	for b.Loop() {
		RepeatRune('√', 50)
	}
}
