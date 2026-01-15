package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}

// Run go test ./iteration -bench=. to execute BenchmarkRepeat
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
