package iteration

import "strings"

// Repeats a given character as many times as the given repeatCount
func Repeat(char string, repeatCount int) string {
	var output strings.Builder
	for i := 0; i < repeatCount; i++ {
		output.WriteString(char)
	}
	return output.String()
}
