/*
Add `//go:generate version` or `//go:generate github.com/abakum/version` to `main.go` so that changes in the `VERSION` file
and for `Windows` in the 'winres' directory affect the result of 'go build'. After the changes and before `go build`, run `go generate`.

Добавь `//go:generate version` или `//go:generate github.com/abakum/version` в `main.go` чтоб изменения в файле `VERSION`
а для `Windows` и в каталоге `winres` учитывались при `go build`. После изменений и перед `go build` запускай `go generate`.
*/
package main

import (
	_ "embed"
	"fmt"

	version "github.com/abakum/version/lib"
)

// For `winres` do `go install github.com/tc-hib/go-winres@latest` then `go-winres init`.
// Do `go install github.com/abakum/version` then uncoment next line to use binary `version` for `go generate`.
///go:generate go run version

// Or do not install binary `version` but `go get github.com/abakum/version@latest` then
// append `version "github.com/abakum/version/lib"` to import block and uncomment next 2 line for `go generate`.
var _ = version.Ver

//go:generate go run github.com/abakum/version

//go:embed VERSION
var VERSION string

func main() {
	fmt.Println(VERSION)

}
