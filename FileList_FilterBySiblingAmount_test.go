package fileIO

import (
	"strconv"
	"testing"
)

// Test if only the paths are returned that match a specific level, starting from a specific path
func TestFilterBySiblingAmount(t *testing.T) {
	expectedValues := []string{"folderA/match1.file", "folderA/match2.file"}

	fileList := FileList{Files: []string{
		"folderA/match1.file",
		"folderA/match2.file",
		"folderB/fail.file",
		"folderC/fail1.file",
		"folderC/fail2.file",
		"folderC/fail3.file",
	}}

	fileList = fileList.FilterBySiblingAmount(Eq, 2)

	if len(fileList.Files) != len(expectedValues) {
		t.Fatal("expected length of fileList.Files is", len(expectedValues), "got", len(fileList.Files), ":\n", fileList.Files)
	}

	for index, actualValue := range fileList.Files {
		if expectedValues[index] != actualValue {
			t.Fatal("expected value of fileList.Files["+strconv.Itoa(index)+"] is", expectedValues[index], "got", actualValue)
		}
	}
}
