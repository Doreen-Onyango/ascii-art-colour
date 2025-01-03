package utils

import (
	"errors"
	"strings"
)

var rgbColors = map[string][3]int{
	"red":             {255, 0, 0},
	"green":           {0, 255, 0},
	"blue":            {0, 0, 255},
	"yellow":          {255, 255, 0},
	"cyan":            {0, 255, 255},
	"purple":          {128, 0, 128},
	"white":           {255, 255, 255},
	"black":           {0, 0, 0},
	"orange":          {255, 165, 0},
	"pink":            {255, 192, 203},
	"brown":           {165, 42, 42},
	"gray":            {128, 128, 128},
	"navy":            {0, 0, 128},
	"lime":            {0, 255, 0},
	"maroon":          {128, 0, 0},
	"olive":           {128, 128, 0},
	"silver":          {192, 192, 192},
	"teal":            {0, 128, 128},
	"magenta":         {255, 0, 255},
	"lightgray":       {211, 211, 211},
	"darkgray":        {169, 169, 169},
	"lightblue":       {173, 216, 230},
	"lightgreen":      {144, 238, 144},
	"lightcyan":       {224, 255, 255},
	"lightyellow":     {255, 255, 224},
	"lightpurple":     {216, 191, 216},
	"thistle":         {216, 191, 216},
	"darkred":         {139, 0, 0},
	"darkgreen":       {0, 100, 0},
	"darkblue":        {0, 0, 139},
	"darkcyan":        {0, 139, 139},
	"darkmagenta":     {139, 0, 139},
	"darkyellow":      {204, 204, 0}, // Olive
	"beige":           {245, 245, 220},
	"gold":            {255, 215, 0},
	"bronze":          {205, 127, 50},
	"violet":          {238, 130, 238},
	"indigo":          {75, 0, 130},
	"chartreuse":      {127, 255, 0},
	"aquamarine":      {127, 255, 212},
	"coral":           {255, 127, 80},
	"salmon":          {250, 128, 114},
	"turquoise":       {64, 224, 208},
	"khaki":           {240, 230, 140},
	"orchid":          {218, 112, 214},
	"plum":            {221, 160, 221},
	"tan":             {210, 180, 140},
	"peach":           {255, 218, 185},
	"mint":            {189, 252, 201},
	"lavender":        {230, 230, 250},
	"rose":            {255, 0, 127},
	"crimson":         {220, 20, 60},
	"tomato":          {255, 99, 71},
	"firebrick":       {178, 34, 34},
	"indianred":       {205, 92, 92},
	"lightcoral":      {240, 128, 128},
	"hotpink":         {255, 105, 180},
	"deeppink":        {255, 20, 147},
	"fuchsia":         {255, 0, 255},
	"palevioletred":   {219, 112, 147},
	"mediumvioletred": {199, 21, 133},
	"springgreen":     {0, 255, 127},
	"lawngreen":       {124, 252, 0},
	"limegreen":       {50, 205, 50},
	"forestgreen":     {34, 139, 34},
	"seagreen":        {46, 139, 87},
	"mediumseagreen":  {60, 179, 113},
	"darkslategray":   {47, 79, 79},
	"darkslategrey":   {47, 79, 79},
	"cadetblue":       {95, 158, 160},
	"steelblue":       {70, 130, 180},
	"royalblue":       {65, 105, 225},
	"midnightblue":    {25, 25, 112},
	"mediumblue":      {0, 0, 205},
	"blueviolet":      {138, 43, 226},
	"darkviolet":      {148, 0, 211},
	"darkorchid":      {153, 50, 204},
	"mediumpurple":    {147, 112, 219},
	"rebeccapurple":   {102, 51, 153},
	"navajowhite":     {255, 222, 173},
	"wheat":           {245, 222, 179},
	"burlywood":       {222, 184, 135},
	"sienna":          {160, 82, 45},
	"chocolate":       {210, 105, 30},
	"saddlebrown":     {139, 69, 19},
	"sandybrown":      {244, 164, 96},
	"peru":            {205, 133, 63},
	"goldenrod":       {218, 165, 32},
	"darkgoldenrod":   {184, 134, 11},
}

//returns the RGB values for a specified color string.
func FindColor(s string) ([3]int, error) {
	if color, found := checkNamedColor(s); found {
		return color, nil
	}
	if rgbColor, err := parseSpecialColor(s); err == nil {
		return rgbColor, nil
	}
	return [3]int{}, errors.New("color not available")
}

// checks if the color is in the predefined map.
func checkNamedColor(s string) ([3]int, bool) {
	value, found := rgbColors[s]
	return value, found
}

// checks for hex, RGB, or HSL formats.
func parseSpecialColor(s string) ([3]int, error) {
	switch {
	case strings.HasPrefix(s, "#"):
		return HexToRGB(s)
	case strings.HasPrefix(s, "rgb("):
		return RGBStringToRGB(s)
	case strings.HasPrefix(s, "hsl("):
		return HSLStringToRGB(s)
	default:
		return [3]int{}, errors.New("invalid color format")
	}
}
