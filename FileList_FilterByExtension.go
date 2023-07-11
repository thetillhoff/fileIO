package fileIO

import (
	"path"
)

// Retain those filePaths that have the specified extension
// Format must be `.ext`
func (fileList FileList) FilterByExtension(extension string) FileList {
	return fileList.Filter(
		func(s string) bool {
			return (path.Ext(s) == extension)
		})
}
