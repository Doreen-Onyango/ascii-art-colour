package utils

import "strings"

// processes the command line args to determine the appropriate banner file.
func FindBannerFile(args []string) ([]string, string) {
	if len(args) == 0 {
		return nil, ""
	}

	switch len(args) {
	case 1:
		return args, "banners/standard.txt"
	case 2:
		return handleTwoArgs(args)
	case 3:
		return handleThreeArgs(args)
	default:
		return args, "banners/standard.txt"
	}
}

// processes the case where there are two arguments.
func handleTwoArgs(args []string) ([]string, string) {
	bannerName := strings.ToLower(args[1])
	if isValidBanner(bannerName) {
		return []string{args[0]}, "banners/" + bannerName + ".txt"
	}
	return args, "banners/standard.txt"
}

// processes the case where there are three arguments.
func handleThreeArgs(args []string) ([]string, string) {
	return args[:2], "banners/" + strings.ToLower(args[2]) + ".txt"
}

// checks if the provided banner name is valid.
func isValidBanner(name string) bool {
	validBanners := map[string]bool{
		"standard":   true,
		"thinkertoy": true,
		"shadow":     true,
	}
	return validBanners[name]
}
