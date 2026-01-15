package iteration

import "strings"

// Repeat returns a string consisting of the input char repeated five times.
func Repeat(char string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(char)
	}
	return repeated.String()
}

func RepeatRune(char rune, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteRune(char)
	}
	return repeated.String()
}
