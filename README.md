# `fileIO`

This package provides basic functions for fileIO, but also a FileList object.
It can be used to first recursively retrieve all existing filepaths at a specific location and then filter them based on different rules.

## Installation

```sh
go get github.com/thetillhoff/fileIO@v1.0.0 // Change version if needed
```

## Basic functions

- CopyFile(src string, dst string) error
- ReadFile(filePath string) ([]byte, error)
- ReadFileLineByLine(filePath string) ([]string, error)
- WriteFile(filePath string, []byte content) error

## The cool stuff

- Watch(inclusionPaths []string, exclusionPaths []string, verbose bool, interval time.Duration, actionEvent func(watcher.Event) error) error
  Watch the specified paths for file changes and run the specified actionEvent on every change.
- FileList
  - Generate(fileListPath string, verbose bool) (FileList, error)
    Recursively traverse the tree below the specified path and return a FileList object with all the found filepaths.
    Will only add files, not folders.
  - GenerateWithIgnoreLines
    Recursively traverse the tree below the specified path and return a FileList object with all the found filepaths.
    Allows to ignore paths during the tree traversal. IgnoreLines work like gitignore.
  - FileList.Filter*
    Returns the FilelList after applying the filter. Can be a custom filter (see FileList.Filter) or a predefined one (see FileList.Filter*, for example ByFilename)
  - Add or remove prefixes to all paths. If the whole tree is moved to a differend folder, it is not necessary to reread the whole tree.

## Development

### Releases

All releases follow the semantic versioning schema with a `v` prefix. Example: `v1.2.3`

Releases are made via git-tags.

### Testing

All internal functions and FileList operations are unit tested.
Actual fileIO functions are not yet tested.

### TODO

- Implement test functions for actual fileIO operations.
- Add fileList.FilterByExtension
