package utils

import (
	"fmt"
	"strings"
)

const Reset = "\033[0m"

// formats the input string with the specified RGB color.
func Color(text, colorName string) string {
	rgbValues, err := FindColor(strings.ToLower(colorName))
	if err != nil {
		return fmt.Sprintf("COLOR NOT FOUND: %s", text)
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s%s", rgbValues[0], rgbValues[1], rgbValues[2], text, Reset)
}
