package utils

import (
	"errors"
	"strconv"
	"strings"
)

// converts a hex color string to an RGB array.
func HexToRGB(hex string) ([3]int, error) {
	var rgb [3]int
	hex = normalizeHex(hex)
	if !isValidHex(hex) {
		return rgb, errors.New("invalid hex color")
	}
	return convertHexToRGB(hex)
}

// removes the leading '#' and expands short hex codes.
func normalizeHex(hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		return expandShortHex(hex)
	}
	return hex
}

// expands a 3-character hex code to 6 characters.
func expandShortHex(hex string) string {
	var expanded strings.Builder
	for _, c := range hex {
		expanded.WriteString(string(c) + string(c))
	}
	return expanded.String()
}

// checks if the hex string is valid (must be 6 characters).
func isValidHex(hex string) bool {
	return len(hex) == 6
}

// converts a valid 6-character hex string to an RGB array.
func convertHexToRGB(hex string) ([3]int, error) {
	var rgb [3]int
	for i := 0; i < 3; i++ {
		val, err := strconv.ParseInt(hex[i*2:i*2+2], 16, 64)
		if err != nil {
			return rgb, errors.New("invalid hex color")
		}
		rgb[i] = int(val)
	}
	return rgb, nil
}
