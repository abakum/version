//go:build windows
// +build windows

package version

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"syscall"
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

// attrib -a path
func IsFileA(path string, clear bool) (bool, error) {
	// Appending `\\?\` to the absolute path helps with
	// preventing 'Path Not Specified Error' when accessing
	// long paths and filenames
	// https://docs.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation?tabs=cmd
	pointer, err := syscall.UTF16PtrFromString(`\\?\` + path)
	a := false
	if err != nil {
		return a, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return a, err
	}

	a = attributes&syscall.FILE_ATTRIBUTE_ARCHIVE != 0
	if a {
		log.Println("is A for", path)
		if clear {
			log.Println("clear A for", path)
			err = syscall.SetFileAttributes(pointer, attributes^syscall.FILE_ATTRIBUTE_ARCHIVE)
			if err != nil {
				return a, err
			}
		}
	}
	return a, nil
}
