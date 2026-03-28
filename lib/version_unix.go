//go:build !windows
// +build !windows

package version

import (
	"encoding/binary"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

const xattrName = "user.mtime"

// attrib -a path
func IsFileA(path string, clear bool) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	currentMtime := info.ModTime().UnixNano()

	// Попробовать прочитать сохранённый mtime из xattr
	var storedBuf [8]byte
	_, err = unix.Getxattr(path, xattrName, storedBuf[:])

	var a bool
	if err != nil {
		// xattr не существует → файл ещё не обрабатывался
		a = true
	} else {
		storedMtime := int64(binary.LittleEndian.Uint64(storedBuf[:]))
		a = currentMtime != storedMtime
	}

	if a {
		log.Println("is A for", path)
		if clear {
			log.Println("clear A for", path)
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], uint64(currentMtime))
			err = unix.Setxattr(path, xattrName, buf[:], 0)
			if err != nil {
				return a, err
			}
		}
	}
	return a, nil
}
