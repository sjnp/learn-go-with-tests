package iteration

import "strings"

func Repeat(char string, times int) string {
	result := ""
	for range times {
		result += char
	}

	return result
}

func FastRepeat(char string, times int) string {
	var result strings.Builder
	for range times {
		result.WriteString(char)
	}

	return result.String()
}
