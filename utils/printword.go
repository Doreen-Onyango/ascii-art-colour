package utils

import (
	"strings"
)

// prints a word using predefined lines from the content file.
func PrintWord(color, word, tobecolored string, contentLines []string) string {
	linesOfSlice := make([]string, 8)

	if tobecolored == "" || color == "" {
		return buildLines(word, contentLines, linesOfSlice, false)
	}
	return buildLines(word, contentLines, linesOfSlice, true, color, tobecolored)
}

// constructs lines for the word, optionally applying color.
func buildLines(word string, contentLines []string, linesOfSlice []string, colored bool, args ...string) string {
	start := 0
	var color string
	var tobecolored string

	if colored {
		color = args[0]
		tobecolored = args[1]
	}

	for start < len(word) {
		if colored && strings.HasPrefix(word[start:], tobecolored) {
			appendColoredLines(tobecolored, color, contentLines, linesOfSlice)
			start += len(tobecolored)
		} else {
			appendNormalLines(word[start], contentLines, linesOfSlice)
			start++
		}
	}

	return strings.Join(linesOfSlice, "\n")
}

// appends colored lines for the substring.
func appendColoredLines(tobecolored, color string, contentLines []string, linesOfSlice []string) {
	for i := 1; i < 9; i++ {
		for _, char := range tobecolored {
			linesOfSlice[i-1] += Color(contentLines[int(char-32)*9+i], color)
		}
	}
}

// appends normal lines for a character.
func appendNormalLines(char byte, contentLines []string, linesOfSlice []string) {
	for i := 1; i < 9; i++ {
		linesOfSlice[i-1] += contentLines[int(char-32)*9+i]
	}
}
