package fileIO

import "strings"

// Returns the string without a trailing slash. If there is none, it is returned as is
func ensureNoTrailingSlash(text string) string {

	return strings.TrimSuffix(text, "/")

}
