package utils

// Check if the character's ASCII value is within the range of printable ASCII characters (32-126)

func IsASCII(s string) bool {
	for _, char := range s {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
