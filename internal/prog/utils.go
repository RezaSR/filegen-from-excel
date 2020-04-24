package prog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NormalizePath(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	return filepath.Clean(path)
}

func FileExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("Path %q is not a file.", path)
	}

	return nil
}

func DirExists(path string, create bool) error {
	info, err := os.Stat(path)
	if err != nil {
		if create {
			err = os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else if !info.IsDir() {
		return fmt.Errorf("Path %q is not a directory.", path)
	}

	return nil
}
