//go:build windows
// +build windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

//go:generate go run .

func main() {
	log.SetFlags(log.Lshortfile)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	ver := "VERSION"
	if len(os.Args) > 1 {
		ver = os.Args[1]
	}
	ver = filepath.Join(wd, ver)
	todo, err := isA(ver, true)
	if err != nil {
		log.Fatalln(err)
	}
	if todo {
		data, err := os.ReadFile(ver)
		if err != nil {
			log.Fatalln(err)
		}

		// git tag v%VERSION%-lw
		cmd := exec.Command("git",
			"tag",
			fmt.Sprintf("v%s-lw", strings.TrimSpace(string(data))),
		)
		data, err = cmd.Output()
		log.Println(cmd.Args, err, string(data))
		if err != nil {
			return
		}

		// git push origin --tags
		cmd = exec.Command("git",
			"push",
			"origin",
			"--tags",
		)
		data, err = cmd.Output()
		log.Println(cmd.Args, err, string(data))
	}
}

func isA(path string, clear bool) (bool, error) {

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}

	// Appending `\\?\` to the absolute path helps with
	// preventing 'Path Not Specified Error' when accessing
	// long paths and filenames
	// https://docs.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation?tabs=cmd
	pointer, err := syscall.UTF16PtrFromString(`\\?\` + absPath)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}

	a := attributes&syscall.FILE_ATTRIBUTE_ARCHIVE != 0
	if a {
		log.Println("is A for", path)
		if clear {
			log.Println("clear A for", path)
			err = syscall.SetFileAttributes(pointer, attributes^syscall.FILE_ATTRIBUTE_ARCHIVE)
			if err != nil {
				return true, err
			}
		}
	}
	return a, nil
}
