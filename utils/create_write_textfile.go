package utils

import (
	"fmt"
	"os"
	"strings"
)

// Create a file at the specified path and writes the provided content to it.
// The function ensures the file has a .txt extension and
// returns an error if any issues occur during file creation or writing.

func CreateTextFile(filePath string, content string) error {
	if !strings.HasSuffix(filePath, ".txt") {
		return fmt.Errorf("invalid file extension: expected .txt")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing content: %w", err)
	}

	return nil
}
