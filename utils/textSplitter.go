package utils

import (
	"strings"
)

// Splits the input string into lines based on the presence of specific line break chars.
// If string contains the char 'o', it splits using the "\r\n" .
// Otherwise, it splits using the "\n" .

func SplitText(text string) []string {
	if strings.Contains(text, "o") {
		return strings.Split(text, "\r\n")
	}
	return strings.Split(text, "\n")
}
