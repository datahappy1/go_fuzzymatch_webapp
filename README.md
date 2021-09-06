# go_fuzzymatch_webapp

[![Go Report Card](https://goreportcard.com/badge/github.com/datahappy1/go_fuzzymatch_webapp)](https://goreportcard.com/report/github.com/datahappy1/go_fuzzymatch_webapp)
[![DeepScan grade](https://deepscan.io/api/teams/13362/projects/16364/branches/348988/badge/grade.svg)](https://deepscan.io/dashboard#view=project&tid=13362&pid=16364&bid=348988)

A fuzzymatching web application written using Go and Javascript. 

## How to run locally the application back-end
1) Install `Golang 1.15+`
2) set environment variable name `environment` to `development` in your IDE
3) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go mod vendor`
4) `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go run main.go`
- The API by default runs at http://localhost:8080/api/v1/requests/
- API Documentation is located [here](https://github.com/datahappy1/go_fuzzymatch_webapp/blob/main/ui/src/api_documentation.js)
- For back-end configuration, check files [here](https://github.com/datahappy1/go_fuzzymatch_webapp/tree/main/api/config)
- For running local tests, run `go\src\github.com\datahappy1\go_fuzzymatch_webapp>go test`

## How to run locally the application front-end
1) Install `npm`
2) Install `Webpack` based on [this tutorial](https://webpack.js.org/guides/installation/)
3) For local development, bundle frontend code using Webpack `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:development`
- For production release, bundle frontend code using Webpack `go\src\github.com\datahappy1\go_fuzzymatch_webapp\ui>npm run build:production`
- The application by default runs at http://localhost:8080/

## Screenshot of the app
![image](https://github.com/datahappy1/go_fuzzymatch_webapp/blob/main/app_screenshot.PNG)