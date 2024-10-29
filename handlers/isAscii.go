package utils

// confirms that a character is a printable ASCII.
func IsASCII(s string) bool {
	for _, char := range s {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
