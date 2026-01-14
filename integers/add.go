package integers

// Add takes two integers and returns the sum of them.
// (sum int) - is a named return value
//
// There is also ExampleAdd func that is composed of Example + func name Add to show example below
// ExampleAdd Output is evaluated at go test.
// See https://go.dev/blog/examples
//
// Notice the special format of the comment, Output: 6. While the example will always be compiled, adding this comment means the example will also be executed.
func Add(a, b int) (sum int) {
	sum = a + b
	return
}
