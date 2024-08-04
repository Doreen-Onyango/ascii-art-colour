package utils

import "strings"

// FindBannerFile processes the command line args to determine the appropriate banner file.
// It returns the modified argument slice and the path to the banner file.
func FindBannerFile(args []string) ([]string, string) {
	if len(args) == 0 {
		return nil, ""
	}

	switch len(args) {
	case 1:
		return args, "banners/standard.txt"
	case 2:
		// Two args, check if the second one is a valid banner name
		bannerName := strings.ToLower(args[1])
		if bannerName == "standard" || bannerName == "thinkertoy" || bannerName == "shadow" {
			return []string{args[0]}, "banners/" + bannerName + ".txt"
		}
		return args, "banners/standard.txt"
	case 3:
		// Three args, use the third argument as the banner file name
		return args[:2], "banners/" + strings.ToLower(args[2]) + ".txt"
	default:
		return args, "banners/standard.txt"
	}
}
