package fileIO

import "strings"

// Returns the string with a trailing slash. If there already is one, it is returned as is
func ensureTrailingSlash(text string) string {

	// Ensure empty strings are returned unchanged
	if len(text) == 0 {
		return ""
	}

	if !strings.HasSuffix(text, "/") {
		return text + "/"
	} else {
		return text
	}
}
