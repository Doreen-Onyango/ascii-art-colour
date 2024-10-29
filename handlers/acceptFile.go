package utils

import (
	"os"
)

// IsFileSizeWithinExpectedRanges checks if the size of the specified file matches one of the predefined valid sizes.
func ValidateFileSize(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	// Get the file information
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}

	// Define valid file sizes
	validSizes := map[int64]bool{
		6623: true,
		7463: true,
		5558: true,
	}

	// Check if the file size is in the list of valid sizes
	return validSizes[fileInfo.Size()]
}
