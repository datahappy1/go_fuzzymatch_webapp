package main

//https://github.com/kelvins/GoApiTutorial/blob/master/main_test.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/datahappy1/go_fuzzymatch_webapp/api/controller"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
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

func TestCreateValidPostRequest(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if controller.IsValidUUID(m["RequestID"]) == false {
		t.Errorf("Invalid RequestID. Got '%s'", m["RequestID"])
	}

	if m["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m["error"])
	}

	model.DeleteFuzzyMatchDAOInRequestsData(m["RequestID"])

}

//here parametrize error cases for invalid post request
func TestCreateInvalidPostRequestInvalidPayload(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1", "mode": "Invalid"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotAcceptable, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "error invalid request" {
		t.Errorf("Expected no error. Got '%s'", m["error"])
	}

	model.DeleteFuzzyMatchDAOInRequestsData(m["RequestID"])

}

func TestCreateInvalidPostRequestTooManyRequestsFromSameIP(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	checkResponseCode(t, http.StatusOK, response1.Code)

	var m1 map[string]string
	json.Unmarshal(response1.Body.Bytes(), &m1)

	req2, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req2.RemoteAddr = "0.0.0.0:80"
	response2 := executeRequest(req2)

	checkResponseCode(t, http.StatusTooManyRequests, response2.Code)

	var m2 map[string]string
	json.Unmarshal(response2.Body.Bytes(), &m2)

	if m2["error"] != "too many requests from IP address in flight, collect previous request "+m1["RequestID"]+" data first" {
		t.Errorf("Expected too many requests from IP address in flight error. Got '%s'", m2["error"])
	}

	model.DeleteFuzzyMatchDAOInRequestsData(m1["RequestID"])
	model.DeleteFuzzyMatchDAOInRequestsData(m2["RequestID"])
}

func TestCreateInvalidPostRequestTooManyOverallRequests(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	req2, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req2.RemoteAddr = "0.0.0.1:80"
	response2 := executeRequest(req2)

	req3, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req3.RemoteAddr = "0.0.0.3:80"
	response3 := executeRequest(req3)

	checkResponseCode(t, http.StatusOK, response1.Code)

	var m1 map[string]string
	json.Unmarshal(response1.Body.Bytes(), &m1)

	checkResponseCode(t, http.StatusOK, response2.Code)

	var m2 map[string]string
	json.Unmarshal(response2.Body.Bytes(), &m2)

	checkResponseCode(t, http.StatusTooManyRequests, response3.Code)

	var m3 map[string]string
	json.Unmarshal(response3.Body.Bytes(), &m3)

	if m3["error"] != "too many overall requests in flight, try later" {
		t.Errorf("Expected too many overall requests error. Got '%s'", m3["error"])
	}

	model.DeleteFuzzyMatchDAOInRequestsData(m1["RequestID"])
	model.DeleteFuzzyMatchDAOInRequestsData(m2["RequestID"])
	model.DeleteFuzzyMatchDAOInRequestsData(m3["RequestID"])
}

func TestCreateValidGetRequest(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	checkResponseCode(t, http.StatusOK, response1.Code)

	var m1 map[string]string
	json.Unmarshal(response1.Body.Bytes(), &m1)

	requestID := m1["RequestID"]

	req2, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	response2 := executeRequest(req2)

	checkResponseCode(t, http.StatusOK, response2.Code)

	var m2 map[string]string
	json.Unmarshal(response2.Body.Bytes(), &m2)

	fmt.Println(m2["error"])
	if m2["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m2["error"])
	}
}

func TestCreateInvalidGetRequest(t *testing.T) {
	var tests = []struct {
		stringsToMatch, stringsToMatchIn, mode string
		responseStatusCode                     int
	}{
		{stringsToMatch: "aplle", stringsToMatchIn: "tree", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "apple inc", stringsToMatchIn: "apple inc", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "apple inc", stringsToMatchIn: "Apple Inc.", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "Apple", stringsToMatchIn: "Apple Inc", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "aplle", stringsToMatchIn: "Apple", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "Apple Corp.", stringsToMatchIn: "Apple Corp. GMBH", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "Apple Inc.", stringsToMatchIn: "GMBH Apple Corp", mode: "simple", responseStatusCode: 201},
		{stringsToMatch: "Aplle Inc.", stringsToMatchIn: "GMBH Apple Corp", mode: "simple", responseStatusCode: 201},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%s,%d", tt.stringsToMatch, tt.stringsToMatchIn, tt.mode, tt.responseStatusCode)
		fmt.Println(testname)

		payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

		req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
		req1.RemoteAddr = "0.0.0.0:80"
		response1 := executeRequest(req1)

		checkResponseCode(t, http.StatusOK, response1.Code)

		var m1 map[string]string
		json.Unmarshal(response1.Body.Bytes(), &m1)

		requestID := m1["RequestID"]

		req2, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
		req2.RemoteAddr = "0.0.0.0:80"
		response := executeRequest(req2)

		//checkResponseCode(t, http.StatusOK, response.Code)
		checkResponseCode(t, tt.responseStatusCode, response.Code)

		var m2 map[string]string
		json.Unmarshal(response.Body.Bytes(), &m2)

		fmt.Println(m2)
		if m2["error"] != "" {
			t.Errorf("Expected no error. Got '%s'", m2["error"])
		}
	}
}
