package main

//https://github.com/kelvins/GoApiTutorial/blob/master/main_test.go

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateInvalidRequest(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"testt","stringsToMatchInd":"test", "mode": "simple"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

//https://github.com/kelvins/GoApiTutorial/blob/master/main_test.go
//https://webpack.js.org/guides/getting-started/
