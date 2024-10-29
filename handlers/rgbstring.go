package utils

import (
	"errors"
	"strconv"
	"strings"
)

// converts an rgb(r, g, b) string to an RGB array
func RGBStringToRGB(rgbStr string) ([3]int, error) {
	var rgb [3]int
	rgbStr = strings.Trim(rgbStr, "rgb()")
	parts := strings.Split(rgbStr, ",")

	if len(parts) != 3 {
		return rgb, errors.New("invalid rgb color")
	}

	// Convert string parts to integers
	for i, part := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return rgb, err
		}
		rgb[i] = val
	}

	return rgb, nil
}
