package fileIO

import (
	"regexp"
)

// Retain those filePaths that match the specified regular expression
func (fileList FileList) FilterByFilenameRegex(regex string) FileList {
	return fileList.Filter(
		func(s string) bool {
			exp := regexp.MustCompile(regex)
			return exp.MatchString(s)
		})
}
