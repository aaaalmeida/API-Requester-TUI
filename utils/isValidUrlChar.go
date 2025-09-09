package utils

import "regexp"

var urlCharRegex = regexp.MustCompile(`^[a-zA-Z0-9\-\._:/\?#\[\]@!\$&'\(\)\*\+,;=]$`)

func IsValidUrlChar(s string) bool {
	return urlCharRegex.MatchString(s)
}
