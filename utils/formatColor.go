package utils

import (
	"fmt"
	"strings"
)

const Reset = "\033[0m"

// Format the input string with the specified RGB color.
// Returns an error message with the original string, if color not found.

func Color(text, colorName string) string {
	// Get RGB values for the specified color
	rgbValues, err := FindColor(strings.ToLower(colorName))
	if err != nil {
		return fmt.Sprintf("COLOR NOT FOUND: %s", text)
	}

	// Return the formatted string with the specified color
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s%s", rgbValues[0], rgbValues[1], rgbValues[2], text, Reset)
}
