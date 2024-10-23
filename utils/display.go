package utils

import (
	"fmt"
	"strings"
)

func DisplayText(color string, inputslice []string, contentLines []string) string {
	var tobecolored, input string

	if len(inputslice) == 2 {
		tobecolored = inputslice[0]
		input = inputslice[1]
	} else {
		tobecolored = inputslice[0]
		input = tobecolored
	}

	if input == "" || input == "\\n" || input == "\n" {
		return ""
	}

	input = UnescapeSpecialCharacters(input)
	tobecolored = UnescapeSpecialCharacters(tobecolored)
	wordslice := strings.Split(input, "\\n")
	tbcsplit := strings.Split(tobecolored, "\\n")

	var s strings.Builder

	for i, word := range wordslice {
		if word == "" {
			s.WriteString("\n")
			continue
		}

		if IsASCII(word) && IsASCII(tobecolored) {
			switch {
			case i < len(tbcsplit) && len(tbcsplit) > 1:
				s.WriteString(PrintWord(color, word, tbcsplit[i], contentLines))
			case i >= len(tbcsplit) && len(tbcsplit) > 1:
				s.WriteString(PrintWord(color, word, "", contentLines))
			default:
				s.WriteString(PrintWord(color, word, tobecolored, contentLines))
			}
			if i != len(wordslice)-1 {
				s.WriteString("\n")
			}
		} else {
			fmt.Println("Invalid input:", word, tobecolored)
		}
	}

	return s.String()
}
