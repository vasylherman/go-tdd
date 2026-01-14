package integers_test

import (
	"fmt"
	"github.comvasylherman/go-tdd/integers"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}

func ExampleAdd() {
	sum := integers.Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}

func ExampleAdd_second() {
	sum := integers.Add(3, 2)
	fmt.Println(sum)
	// Output:
	// 5
}
