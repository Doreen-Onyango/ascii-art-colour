package utils

import (
	"fmt"
	"strings"
)

func DisplayText(color string, inputslice []string, contentLines []string) string {
	// initialize words to be colored and input
	tobecolored := ""
	input := ""
	s := ""
	if len(inputslice) == 2 {
		tobecolored = inputslice[0]
		input = inputslice[1]
	} else {
		tobecolored = inputslice[0]
		input = inputslice[0]
	}
	if input == "" {
		return ""
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		return ""
	}
	input = UnescapeSpecialCharacters(input)
	tobecolored = UnescapeSpecialCharacters(tobecolored)
	wordslice := strings.Split(input, "\\n")
	tbcsplit := strings.Split(tobecolored, "\\n")

	for i, word := range wordslice {
		if word == "" {
			s += "\n"
		} else {
			if IsASCII(word) && IsASCII(tobecolored) {
				if i < len(tbcsplit) && len(tbcsplit) > 1 {
					// coloring elements of each word on each line
					s += PrintWord(color, word, tbcsplit[i], contentLines)
				}
				if i >= len(tbcsplit) && len(tbcsplit) > 1 {
					tobecolored = ""
					// unrepresented lines are left uncolored
					s += PrintWord(color, word, tobecolored, contentLines)
				}
				// color each element in every line
				if len(tbcsplit) == 1 {
					s += PrintWord(color, word, tobecolored, contentLines)
				}
				if i != len(wordslice)-1 {
					s += "\n"
				}
			} else {
				fmt.Println("Invalid input:", word, tobecolored)
			}
		}
	}
	return s
}
