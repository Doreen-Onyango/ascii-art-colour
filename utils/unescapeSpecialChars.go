package utils

import "strings"

// Replaces escaped sequences in the input string with their actual character equivalents.

func UnescapeSpecialCharacters(input string) string {
	input = strings.ReplaceAll(input, "\\b", "\b")   // Replace escaped backspace with actual backspace character
	input = strings.ReplaceAll(input, "\\a", "\a")   // Replace escaped alert (bell) with actual alert character
	input = strings.ReplaceAll(input, "\\v", "\v")   // Replace escaped vertical tab with actual vertical tab character
	input = strings.ReplaceAll(input, "\n", "\\n")   // Replace newline with escaped newline
	input = strings.ReplaceAll(input, "\\f", "\f")   // Replace escaped form feed with actual form feed character
	input = strings.ReplaceAll(input, "\\t", "    ") // Replace escaped tab with four spaces
	input = strings.ReplaceAll(input, "\\r", "\r")   // Replace escaped carriage return with actual carriage return character

	return input
}
