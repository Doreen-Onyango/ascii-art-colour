package utils

import (
	"errors"
	"strconv"
	"strings"
)

// converts an hsl(h, s, l) string to an RGB array.
func HSLStringToRGB(hslStr string) ([3]int, error) {
	h, s, l, err := parseHSL(hslStr)
	if err != nil {
		return [3]int{}, err
	}
	return HSLToRGB(h, s, l), nil
}

// extracts and converts HSL values from the hsl string.
func parseHSL(hslStr string) (int, int, int, error) {
	hslStr = strings.TrimPrefix(hslStr, "hsl(")
	hslStr = strings.TrimSuffix(hslStr, ")")
	parts := strings.Split(hslStr, ",")

	if len(parts) != 3 {
		return 0, 0, 0, errors.New("invalid hsl color")
	}

	h, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, 0, errors.New("invalid hsl color")
	}

	s, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[1], "%")))
	if err != nil {
		return 0, 0, 0, errors.New("invalid hsl color")
	}

	l, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[2], "%")))
	if err != nil {
		return 0, 0, 0, errors.New("invalid hsl color")
	}

	return h, s, l, nil
}

// converts HSL values to RGB values.
func HSLToRGB(h, s, l int) [3]int {
	h = h % 360
	s = s % 101
	l = l % 101

	c := (1.0 - float64(abs(2*l-100))/100.0) * float64(s) / 100.0
	x := c * (1.0 - float64(abs(h/60%2-1)))
	m := float64(l)/100.0 - c/2.0

	r, g, b := calculateRGBComponents(float64(h), c, x)

	return [3]int{
		int((r + m) * 255.0),
		int((g + m) * 255.0),
		int((b + m) * 255.0),
	}
}

// determines RGB values based on HSL components.
func calculateRGBComponents(h, c, x float64) (float64, float64, float64) {
	switch {
	case h < 60:
		return c, x, 0
	case h < 120:
		return x, c, 0
	case h < 180:
		return 0, c, x
	case h < 240:
		return 0, x, c
	case h < 300:
		return x, 0, c
	default:
		return c, 0, x
	}
}

// returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
