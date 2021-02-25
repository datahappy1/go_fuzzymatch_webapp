package main

//https://github.com/kelvins/GoApiTutorial/blob/master/main_test.go

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	//"github.com/datahappy1/go_fuzzymatch_webapp/api/controller"
	//"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("development")
	//a.Initialize("root", "", "rest_api_example")
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

// func TestCreateValidPostRequest(t *testing.T) {

// 	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

// 	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if controller.IsValidUUID(m["RequestID"]) == false {
// 		t.Errorf("Invalid RequestID. Got '%s'", m["RequestID"])
// 	}

// 	if m["error"] != "" {
// 		t.Errorf("Expected no error. Got '%s'", m["error"])
// 	}

// 	model.DeleteFuzzyMatchDAOInRequestsData(m["RequestID"])

// }

//here parametrize error cases for invalid post request
// func TestCreateInvalidPostRequestInvalidPayload(t *testing.T) {

// 	payload := []byte(`{"stringsToMatch":"teststring1", "mode": "Invalid"}`)

// 	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusNotAcceptable, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["error"] != "" {
// 		t.Errorf("Expected no error. Got '%s'", m["error"])
// 	}

// 	model.DeleteFuzzyMatchDAOInRequestsData(m["RequestID"])

// }

// func TestCreateInvalidPostRequestTooManyRequestsFromSameIP(t *testing.T) {

// 	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

// 	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m1 map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m1)

// 	firstRequestID := m1["requestID"]

// 	req, _ = http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusTooManyRequests, response.Code)

// 	var m2 map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m2)

// 	if m2["error"] != "too many requests from IP address in flight, collect previous request "+firstRequestID+" data first" {
// 		t.Errorf("Expected too many requests from IP address in flight error. Got '%s'", m2["error"])
// 	}

// 	model.DeleteFuzzyMatchDAOInRequestsData(m1["RequestID"])
// 	model.DeleteFuzzyMatchDAOInRequestsData(m2["RequestID"])

// }

func TestCreateInvalidPostRequestTooManyOverallRequests(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "too Many requests" {
		t.Errorf("Expected no error. Got '%s'", m["error"])
	}

}

// func TestCreateValidGetRequest(t *testing.T) {

// 	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

// 	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))

// 	requestID := model.RequestsData[0].RequestID

// 	req, _ = http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	fmt.Println(m)
// 	if m["error"] != "" {
// 		t.Errorf("Expected no error. Got '%s'", m["error"])
// 	}
// }

// func TestCreateInvalidGetRequest(t *testing.T) {

// 	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

// 	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))

// 	requestID := model.RequestsData[0].RequestID

// 	req, _ = http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
// 	req.RemoteAddr = "0.0.0.0:80"
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	fmt.Println(m)
// 	if m["error"] != "" {
// 		t.Errorf("Expected no error. Got '%s'", m["error"])
// 	}
// }

//https://github.com/kelvins/GoApiTutorial/blob/master/main_test.go
//https://webpack.js.org/guides/getting-started/
