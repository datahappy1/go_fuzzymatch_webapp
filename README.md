# go_fuzzymatch_webapp

## How to run the application back-end
1) Install `Golang 1.15+`
2) set environment variable name `environment` to `development` in your IDE
3) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go mod vendor`
4) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go run main.go`
5) The api run by default runs at http://localhost:8080/api/v1/requests/
6) API Documentation is located here: [here](https://www.github.com/datahappy1/go_fuzzymatch_webapp/blob/master/ui/api_documentation.md)
7) For back-end configuration, check files in the /api/config folder

# How to run the application front-end
1) Install `NPM` & `Webpack`
2) For local development bundle frontend code using Webpack `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:development`
3) For production release bundle frontend code using Webpack `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:production`
4) The application by default runs at http://localhost:8080/