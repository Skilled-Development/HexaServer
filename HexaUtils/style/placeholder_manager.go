package style

import "strings"

func ReplaceAllPlaceholders(s string) string {
	s = strings.ReplaceAll(s, "&", "§")
	return s
}
