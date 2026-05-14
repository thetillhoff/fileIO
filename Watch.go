package fileIO

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Watch watches inclusionPaths recursively for changes and calls actionEvent with
// the path of the changed file. exclusionPaths are skipped. interval is used as a
// debounce window to coalesce rapid changes into a single call.
func Watch(inclusionPaths []string, exclusionPaths []string, verbose bool, interval time.Duration, actionEvent func(path string) error) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	excluded := map[string]bool{}
	for _, p := range exclusionPaths {
		if abs, err := filepath.Abs(p); err == nil {
			excluded[abs] = true
		}
	}

	isExcluded := func(p string) bool {
		abs, _ := filepath.Abs(p)
		return excluded[abs]
	}

	addRecursive := func(root string) error {
		return filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil // skip inaccessible paths
			}
			if isExcluded(p) {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
			if d.IsDir() {
				if verbose {
					log.Println("Watching:", p)
				}
				return w.Add(p)
			}
			return nil
		})
	}

	for _, p := range inclusionPaths {
		if err := addRecursive(p); err != nil {
			return err
		}
	}

	var (
		debounce = time.NewTimer(0)
		lastPath string
	)
	<-debounce.C // drain initial tick

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return nil
			}
			lastPath = event.Name
			// Automatically watch newly created subdirectories
			if event.Has(fsnotify.Create) {
				if info, err := os.Stat(event.Name); err == nil && info.IsDir() && !isExcluded(event.Name) {
					if err := addRecursive(event.Name); err != nil {
						log.Println("Failed to watch new directory:", err)
					}
				}
			}
			debounce.Reset(interval)
		case <-debounce.C:
			if err := actionEvent(lastPath); err != nil {
				log.Println("Watch action error:", err)
			}
		case err, ok := <-w.Errors:
			if !ok {
				return nil
			}
			log.Println("Watcher error:", err)
		}
	}
}
