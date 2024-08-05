package main

import (
	"fmt"
	"os"
	"strings"

	"color/utils"
)

func main() {
	// check enough command line arguments
	usage := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something"
	if len(os.Args[1:]) < 1 || len(os.Args[1:]) > 6 || !utils.ValidateArguments(os.Args[1:]) {
		fmt.Println(usage)
		return
	}
	inputslice := utils.ParseFlags()
	if len(inputslice) == 0 {
		fmt.Println(usage)
		return
	}
	color := *utils.ColorPtr
	output := *utils.OutputPtr
	asciiArt := ""
	var filename string
	inputslice, filename = utils.FindBannerFile(inputslice)
	if utils.ValidateFileSize(filename) {

		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("invalid text file")
			return
		}
		contentLines := utils.SplitText(string(content))
		if color != "" {
			_, errcolor := utils.FindColor(strings.ToLower(color))
			if errcolor != nil {
				fmt.Println(errcolor, color)
				return
			}
		}
		asciiArt = utils.DisplayText(color, inputslice, contentLines)
	} else {
		fmt.Println("invalid text file")
	}
	if output != "" {
		err := utils.CreateTextFile(output, asciiArt)
		if err != nil {
			fmt.Println(usage)
			return
		}
	}
	fmt.Println(asciiArt)
}
