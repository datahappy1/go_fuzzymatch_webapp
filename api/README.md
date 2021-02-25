# Backend
1) Install godep (for Heroku) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go get -u github.com/tools/godep`
2) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>godep save ./...`
3) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go mod vendor`
4) Local run: `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go run main.go app.go`