// Add `//go:generate version` or `//go:generate github.com/abakum/version` to `main.go` so that changes in the `VERSION` file
// and for `Windows` in the 'winres' directory affect the result of 'go build'. After the changes and before `go build`, run `go generate`.

// Добавь `//go:generate version` или `//go:generate github.com/abakum/version` в `main.go` чтоб изменения в файле `VERSION`
// а для `Windows` и в каталоге `winres` учитывались при `go build`. После изменений и перед `go build` запускай `go generate`.
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	version "github.com/abakum/version/lib"
)

//go:generate go run .

//go:embed VERSION
var VERSION string

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println(VERSION)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	ver := version.Ver
	if len(os.Args) > 1 {
		ver = os.Args[1]
	}
	ver = filepath.Join(wd, ver)
	// attrib -a VERSION
	todo, err := version.IsA(ver, true)
	if err != nil {
		log.Fatalln(err)
	}
	if todo {
		data, err := os.ReadFile(ver)
		if err != nil || bytes.Count(data, []byte(".")) != 2 {
			log.Fatalln(err)
		}
		// set /p VERSION=<VERSION
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

	if runtime.GOOS != "windows" {
		return
	}

	winres := filepath.Join(wd, "winres")
	// attrib -a winres\*
	todo, err = version.IsA(winres, true)
	if err != nil || !todo {
		return
	}

	// https://github.com/tc-hib/go-winres
	// go install github.com/tc-hib/go-winres@latest
	// go-winres init
	// go-winres make --product-version=git-tag --file-version=git-tag
	cmd := exec.Command("go-winres",
		"make",
		"--product-version",
		"git-tag",
		"--file-version",
		"git-tag",
	)
	data, err := cmd.Output()
	log.Println(cmd.Args, err, string(data))
}
