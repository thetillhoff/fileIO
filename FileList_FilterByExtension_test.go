package fileIO

import (
	"strconv"
	"testing"
)

// Test if only the paths are returned that point to a matching extension
func TestFilterByExtension(t *testing.T) {
	expectedValues := []string{"matching.valid", "folder/with/subfolder/matching.valid"}

	fileList := FileList{Files: []string{
		"not-matching.invalid",
		"matching.valid",
		"folder/with/subfolder/matching.valid",
		"folder/with/subfolder/not-matching.invalid",
	}}

	fileList = fileList.FilterByExtension(".valid")

	if len(fileList.Files) != len(expectedValues) {
		t.Fatal("expected length of fileList.Files is", len(expectedValues), "got", len(fileList.Files), ":\n", fileList.Files)
	}

	for index, actualValue := range fileList.Files {
		if expectedValues[index] != actualValue {
			t.Fatal("expected value of fileList.Files["+strconv.Itoa(index)+"] is", expectedValues[index], "got", actualValue)
		}
	}
}
