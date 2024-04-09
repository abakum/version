//go:build !windows
// +build !windows

package version

import (
	"log"
	"os"
	"time"
)

// attrib -a path
func IsFileA(path string, clear bool) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	a := info.ModTime().Unix() > 0

	if a {
		log.Println("is A for", path)
		if clear {
			log.Println("clear A for", path)
			os.Chtimes(path, time.Unix(0, 0), time.Unix(0, 0))
		}
	}
	return a, nil
}
