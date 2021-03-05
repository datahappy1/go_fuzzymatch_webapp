package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/datahappy1/go_fuzzymatch_webapp/api/repository"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("development")
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

	if utils.IsValidUUID(m["RequestID"]) == false {
		t.Errorf("Invalid RequestID. Got '%s'", m["RequestID"])
	}

	if m["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m["error"])
	}

	repository.Delete(m["RequestID"])

}

func TestCreateInvalidPostRequestInvalidPayload(t *testing.T) {
	var tests = []struct {
		stringsToMatch, stringsToMatchIn, mode string
		responseStatusCode                     int
	}{
		{stringsToMatch: "apple inc", stringsToMatchIn: "", mode: "simple", responseStatusCode: 406},
		{stringsToMatch: "Apple", stringsToMatchIn: "Apple Inc", mode: "", responseStatusCode: 406},
		{stringsToMatch: "", stringsToMatchIn: "", mode: "simple", responseStatusCode: 406},
		{stringsToMatch: "", stringsToMatchIn: "Apple Corp. GMBH", mode: "", responseStatusCode: 406},
		{stringsToMatch: "Apple Inc.", stringsToMatchIn: "", mode: "", responseStatusCode: 406},
		{stringsToMatch: "", stringsToMatchIn: "", mode: "", responseStatusCode: 406},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%s,%d", tt.stringsToMatch, tt.stringsToMatchIn, tt.mode, tt.responseStatusCode)
		fmt.Println(testname)

		payload := []byte(`{"stringsToMatch":"` + tt.stringsToMatch + `", "stringsToMatchIn":"` + tt.stringsToMatchIn + `", "mode": "` + tt.mode + `"}`)

		req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
		req.RemoteAddr = "0.0.0.0:80"
		response := executeRequest(req)

		checkResponseCode(t, tt.responseStatusCode, response.Code)

		var m map[string]string
		json.Unmarshal(response.Body.Bytes(), &m)

		if m["error"] != "error invalid request" {
			t.Errorf("Expected invalid request error. Got '%s'", m["error"])
		}

		repository.Delete(m["RequestID"])
	}
}

func TestCreateInvalidPostRequestInvalidIP(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "cannot determine IP address" {
		t.Errorf("Expected cannot determine IP address error. Got '%s'", m["error"])
	}

	repository.Delete(m["RequestID"])
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

	if m2["error"] != "too many requests from IP address, collect request "+m1["RequestID"]+" data first" {
		t.Errorf("Expected too many requests from IP address in flight error. Got '%s'", m2["error"])
	}

	repository.Delete(m1["RequestID"])
	repository.Delete(m2["RequestID"])
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

	repository.Delete(m1["RequestID"])
	repository.Delete(m2["RequestID"])
	repository.Delete(m3["RequestID"])
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

	if m2["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m2["error"])
	}

	repository.Delete(m1["RequestID"])

}

//TODO FIXME
func TestCreateValidGetRequestWithReturnedAllRows(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"'teststring1','teststring2','teststring3','teststring4'","stringsToMatchIn":"teststring2", "mode": "simple"}`)

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

	fmt.Println(m2)

	if m2["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m2["error"])
	}

	req3, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	response3 := executeRequest(req3)

	checkResponseCode(t, http.StatusOK, response3.Code)

	var m3 map[string]string
	json.Unmarshal(response3.Body.Bytes(), &m3)

	fmt.Println(m3["ReturnedAllRows"])

	if m3["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m3["error"])
	}

	repository.Delete(m1["RequestID"])

}

func TestCreateInvalidGetRequestInvalidUUID(t *testing.T) {

	requestID := "invalidRequestId"

	req, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "need a valid UUID for request ID" {
		t.Errorf("Expected need a valid UUID for request ID error. Got '%s'", m["error"])
	}

}

func TestCreateInvalidGetRequestInvalidRequestId(t *testing.T) {

	requestID := "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

	req, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "request not found" {
		t.Errorf("Expected request not found error. Got '%s'", m["error"])
	}

}

func TestCreateInvalidGetRequestInvalidURL(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/invalid_requests_path/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "" {
		t.Errorf("Expected no error. Got '%s'", m["error"])
	}

}
