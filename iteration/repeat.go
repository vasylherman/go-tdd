package iteration

// Repeat returns a string consisting of the input char repeated five times.
func Repeat(char string) (repeated string) {
	for range 5 {
		repeated = repeated + char
	}
	return repeated
}
