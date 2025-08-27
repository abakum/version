# version
// Add `//go:generate go run github.com/abakum/version` to `main.go` so that changes in the `VERSION` file
// and for `Windows` in the 'winres' directory affect the result of 'go build'. After the changes and before `go build`, run `go generate`.

// Добавь `//go:generate go run github.com/abakum/version` в `main.go` чтоб изменения в файле `VERSION`
// а для `Windows` и в каталоге `winres` учитывались при `go build`. После изменений и перед `go build` запускай `go generate`.

# [example](example/main.go)