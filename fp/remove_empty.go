package fp

import (
	"os"
	"path/filepath"
)

// RemoveEmpty walks a file system hierarchy and removes all directories with no files.
func RemoveEmpty(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			removeEmpty(filepath.Join(dir, file.Name()))
		}
	}
}

func removeEmpty(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	var count int
	for _, file := range files {
		if file.IsDir() {
			if removeEmpty(filepath.Join(dir, file.Name())) {
				count++
			}
		}
	}
	if count == len(files) {
		return os.Remove(dir) == nil
	}
	return false
}
