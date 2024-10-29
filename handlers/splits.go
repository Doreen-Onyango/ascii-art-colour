package utils

import (
	"strings"
)

// splits strings using carriage return and new line
func SplitText(text string) []string {
	if strings.Contains(text, "o") {
		return strings.Split(text, "\r\n")
	}
	return strings.Split(text, "\n")
}
