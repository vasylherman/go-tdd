package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}

// Run go test ./iteration -bench=. to execute BenchmarkRepeat
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 50)
	}
}
