package fp

import (
	"os"
	"path/filepath"
)

// WriteFile writes data to the named file, creating it and directories if necessary.
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
func WriteFile(name string, data []byte) error {
	dir := filepath.Dir(name)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, os.ModePerm)
}
