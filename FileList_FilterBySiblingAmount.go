package fileIO

import (
	"path"
)

// Retains all filePaths that have the exact specified amount of sibling files.
func (fileList FileList) FilterBySiblingAmount(comparisonOperator ComparisonOperator, amount int) FileList {
	var (
		siblingCount = map[string]int{}
	)

	for _, filePath := range fileList.Files { // For all filepaths
		if _, ok := siblingCount[path.Dir(filePath)]; ok { // If this has a sibling
			siblingCount[path.Dir(filePath)] = siblingCount[path.Dir(filePath)] + 1 // Increment counter for this folder by 1
		} else { // If this doesn't have a sibling (yet)
			siblingCount[path.Dir(filePath)] = 1 // Set counter to 1
		}
	}

	return fileList.Filter(
		func(s string) bool {
			return comparisonOperator(siblingCount[path.Dir(s)], amount)
		})
}
