package utils

func StringShortner(s string) string {
	if len(s) > 80 {
		s = s[:80] + "...."
	}
	return s
}
