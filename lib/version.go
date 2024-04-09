package version

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const Ver = "VERSION"

// attrib -a path\*
func IsA(path string, clear bool) (bool, error) {
	info, err := os.Stat(path)
	A := false
	if err != nil {
		return A, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return A, err
	}
	if !info.IsDir() {
		return IsFileA(absPath, clear)
	}
	matches := []string{}
	err = filepath.WalkDir(absPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return A, err
	}
	log.Println(matches)
	for _, file := range matches {
		a, err := IsFileA(file, clear)
		if err != nil {
			return false, err
		}
		A = A || a
		if !clear && A {
			return A, err
		}
	}
	return A, nil
}
