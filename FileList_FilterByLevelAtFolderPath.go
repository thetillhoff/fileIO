package fileIO

import (
	"strings"
)

// Retains the filepaths below the provided path, which are the specified amount of path-levels below it
// Requires a trailing slash. If there is none, it will be automatically added temporarily.
// The trailing slash is to ensure that not multiple folder names match (e.g. `a/b/` vs `a/bc/` which might both match `a/b`)
func (fileList FileList) FilterByLevelAtFolderPath(path string, level int) FileList {
	path = ensureTrailingSlash(path)
	return fileList.FilterByFolderPath(path).Filter(
		func(s string) bool {
			trimmed := strings.TrimPrefix(s, path)
			return (strings.Count(trimmed, "/") == level)
		})
}
