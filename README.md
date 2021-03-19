# go_fuzzymatch_webapp

## How to run locally the application back-end
1) Make sure you have `Golang 1.15+` installed
2) set environment variable name `environment` to `development` in your IDE
3) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go mod vendor`
4) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go run main.go`
- The api by default runs at http://localhost:8080/api/v1/requests/
- API Documentation Markdown file is located [here](https://github.com/datahappy1/go_fuzzymatch_webapp/blob/main/ui/dist/api_documentation.md)
- For back-end configuration, check files in [here](https://github.com/datahappy1/go_fuzzymatch_webapp/tree/main/api/config)
- To run tests, just run `go test` at the root directory level

# How to run locally the application front-end
1) Install `Webpack`
2) For local development bundle frontend code using Webpack command `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:development`
- For production release bundle frontend code using Webpack command `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:production`
- The application by default runs at http://localhost:8080/