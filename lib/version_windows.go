//go:build windows
// +build windows

package version

import (
	"log"
	"syscall"
)

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
