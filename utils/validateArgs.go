package utils

import (
	"strings"
)

// Checks if the input arguments meet the specified requirements.
// It returns true if the arguments are valid based on the number of arguments and their format.
func ValidateArguments(args []string) bool {
	numArgs := len(args)

	switch numArgs {
	case 1:
		if strings.HasPrefix(args[0], "-") {
			return false
		}
	case 2:
		// Two args, the first must be a flag & the second a non-flag, or the first must be a non-flag & the second a valid banner type
		if isFlag(args[0]) && !isFlag(args[1]) {
			return true
		} else if !isFlag(args[0]) {
			bannerType := strings.ToLower(args[1])
			if bannerType == "standard" || bannerType == "thinkertoy" || bannerType == "shadow" {
				return true
			}
			return false
		}
		return false
	case 3:
		// Three args, the first two should be flags
		if isFlag(args[0]) && isFlag(args[1]) {
			return true
		}
	case 4, 5, 6:
		// Four to six args, the initial args must all be flags
		for i := 0; i < numArgs-3; i++ {
			if !isFlag(args[i]) {
				return false
			}
		}
	default:
		// Any other number of arguments, the input is invalid
		return false
	}
	return true
}

// Checks if a string is a flag in the format --flagname=value
func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "--")
}
