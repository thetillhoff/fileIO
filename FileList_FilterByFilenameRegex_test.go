package fileIO

import (
	"strconv"
	"testing"
)

// Test if only the paths are returned that match a regex
func TestFilterByFilenameRegex(t *testing.T) {
	expectedValues := []string{"match.file", "folder/with/subfolder/match.file"}

	fileList := FileList{Files: []string{
		"fail.file",
		"match.file",
		"folder/with/subfolder/match.file",
		"folder/with/subfolder/fail.file",
	}}

	fileList = fileList.FilterByFilenameRegex("match.file$")

	if len(fileList.Files) != len(expectedValues) {
		t.Fatal("expected length of fileList.Files is", len(expectedValues), "got", len(fileList.Files), ":\n", fileList.Files)
	}

	for index, actualValue := range fileList.Files {
		if expectedValues[index] != actualValue {
			t.Fatal("expected value of fileList.Files["+strconv.Itoa(index)+"] is", expectedValues[index], "got", actualValue)
		}
	}
}
