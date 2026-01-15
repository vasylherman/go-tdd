package iteration

import "strings"

// Repeat returns a string consisting of the input letter repeated five times.
func Repeat(letter string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(letter)
	}
	return repeated.String()
}

func RepeatRune(rune rune, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteRune(rune)
	}
	return repeated.String()
}
