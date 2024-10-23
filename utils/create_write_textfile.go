package utils

import (
	"fmt"
	"os"
	"strings"
)

// CreateTextFile creates a text file at the specified path with the given content.
func CreateTextFile(filePath string, content string) error {
	if err := validateFilePath(filePath); err != nil {
		return err
	}

	if err := writeToFile(filePath, content); err != nil {
		return err
	}

	return nil
}

// validateFilePath checks if the file path has a .txt suffix.
func validateFilePath(filePath string) error {
	if !strings.HasSuffix(filePath, ".txt") {
		return fmt.Errorf("invalid file extension: expected .txt")
	}
	return nil
}

// writeToFile creates a file and writes the content to it.
func writeToFile(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("error writing content: %w", err)
	}

	return nil
}
