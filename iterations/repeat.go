package iterations

import "strings"

func Repeat(char string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(char)
	}

	return result.String()
}

func StandardRepeat(char string, count int) string {
	return strings.Repeat(char, count)
}
