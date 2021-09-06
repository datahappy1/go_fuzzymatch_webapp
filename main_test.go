package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/datahappy1/go_fuzzymatch_webapp/api"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/repository"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var a api.App

var shortTimeOutInSeconds = 10

func TestMain(m *testing.M) {
	a = api.App{}

	a.InitializeAPI("testing")
	a.InitializeDB()

	go a.ClearAppRequestData()

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

type failureResponse struct {
	Error string `json:"error"`
}

type successPostRequestResponse struct {
	RequestID string `json:"RequestID"`
}

type result struct {
	StringToMatch string `json:"StringToMatch"`
	StringMatched string `json:"StringMatched"`
	Result        int    `json:"Result"`
}
type successGetResultsResponse struct {
	RequestID       string   `json:"RequestID"`
	Mode            string   `json:"Mode"`
	RequestedOn     string   `json:"RequestedOn"`
	ReturnedAllRows bool     `json:"ReturnedAllRows"`
	Results         []result `json:"Results"`
}

func TestCreateValidPostRequest(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	SuccessResponse := successPostRequestResponse{}
	json.Unmarshal(response.Body.Bytes(), &SuccessResponse)

	if utils.IsValidUUID(SuccessResponse.RequestID) == false {
		t.Errorf("Invalid RequestID. Got '%s'", SuccessResponse.RequestID)
	}

	repository.Delete(a.Repo.DB, SuccessResponse.RequestID)

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

		FailureResponse := failureResponse{}
		json.Unmarshal(response.Body.Bytes(), &FailureResponse)

		if FailureResponse.Error != "error invalid request" {
			t.Errorf("Expected invalid request error. Got '%s'", FailureResponse.Error)
		}

	}
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

	SuccessResponse1 := successPostRequestResponse{}
	json.Unmarshal(response1.Body.Bytes(), &SuccessResponse1)

	checkResponseCode(t, http.StatusOK, response2.Code)

	SuccessResponse2 := successPostRequestResponse{}
	json.Unmarshal(response2.Body.Bytes(), &SuccessResponse2)

	checkResponseCode(t, http.StatusTooManyRequests, response3.Code)

	FailureResponse := failureResponse{}
	json.Unmarshal(response3.Body.Bytes(), &FailureResponse)

	if FailureResponse.Error != "too many overall requests in flight, try later" {
		t.Errorf("Expected too many overall requests error. Got '%s'", FailureResponse.Error)
	}

	repository.Delete(a.Repo.DB, SuccessResponse1.RequestID)
	repository.Delete(a.Repo.DB, SuccessResponse2.RequestID)

}

func TestCreateValidGetRequest(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	checkResponseCode(t, http.StatusOK, response1.Code)

	SuccessResponse1 := successPostRequestResponse{}
	json.Unmarshal(response1.Body.Bytes(), &SuccessResponse1)

	req2, _ := http.NewRequest("GET", "/api/v1/requests/"+SuccessResponse1.RequestID+"/", nil)
	response2 := executeRequest(req2)

	checkResponseCode(t, http.StatusOK, response2.Code)

	SuccessResultsResponse := successGetResultsResponse{}
	json.Unmarshal(response2.Body.Bytes(), &SuccessResultsResponse)

	if SuccessResultsResponse.Mode != "simple" {
		t.Errorf("Invalid mode value, expected simple. Got '%s'", SuccessResultsResponse.Mode)
	}

	if SuccessResultsResponse.ReturnedAllRows != true {
		t.Errorf("Invalid ReturnedAllRows value, expected true. Got '%v'", SuccessResultsResponse.ReturnedAllRows)
	}

	if utils.IsValidUUID(SuccessResultsResponse.RequestID) == false {
		t.Errorf("Invalid RequestID value, expected valid UUID. Got '%s'", SuccessResponse1.RequestID)
	}

	if len(SuccessResultsResponse.Results) != 1 {
		t.Errorf("Invalid Results array item count value, expected 1. Got '%d'", len(SuccessResultsResponse.Results))
	}

	_, err := time.Parse("2006-01-02T15:04:05", SuccessResultsResponse.RequestedOn)
	if err != nil {
		t.Errorf("Invalid RequestedOn value, cannot parse to layout 2006-01-02T15:04:05. Got '%s'", SuccessResultsResponse.RequestedOn)
	}

	if SuccessResultsResponse.Results[0].StringToMatch != "teststring1" {
		t.Errorf("Invalid StringToMatch value, expected teststring1. Got '%s'", SuccessResultsResponse.Results[0].StringToMatch)
	}

	if SuccessResultsResponse.Results[0].StringMatched != "teststring2" {
		t.Errorf("Invalid StringMatched value, expected teststring2. Got '%s'", SuccessResultsResponse.Results[0].StringMatched)
	}

	if SuccessResultsResponse.Results[0].Result != 90 {
		t.Errorf("Invalid Result value, expected 90. Got '%d'", SuccessResultsResponse.Results[0].Result)
	}

	repository.Delete(a.Repo.DB, SuccessResponse1.RequestID)

}

func TestCreateValidGetRequestWithReturnedAllRows(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"'teststring1','teststring2','teststring3','teststring4'","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	checkResponseCode(t, http.StatusOK, response1.Code)

	SuccessRequestResponse := successPostRequestResponse{}
	json.Unmarshal(response1.Body.Bytes(), &SuccessRequestResponse)

	req2, _ := http.NewRequest("GET", "/api/v1/requests/"+SuccessRequestResponse.RequestID+"/", nil)
	response2 := executeRequest(req2)

	checkResponseCode(t, http.StatusOK, response2.Code)

	SuccessResultsResponse := successGetResultsResponse{}
	json.Unmarshal(response2.Body.Bytes(), &SuccessResultsResponse)

	if SuccessResultsResponse.ReturnedAllRows != false {
		t.Errorf("Expected ReturnedAllRows false. Got '%t'", SuccessResultsResponse.ReturnedAllRows)
	}

	req3, _ := http.NewRequest("GET", "/api/v1/requests/"+SuccessRequestResponse.RequestID+"/", nil)
	response3 := executeRequest(req3)

	checkResponseCode(t, http.StatusOK, response3.Code)

	SuccessResultsResponse = successGetResultsResponse{}
	json.Unmarshal(response3.Body.Bytes(), &SuccessResultsResponse)

	if SuccessResultsResponse.ReturnedAllRows != true {
		t.Errorf("Expected ReturnedAllRows true. Got '%t'", SuccessResultsResponse.ReturnedAllRows)
	}

	repository.Delete(a.Repo.DB, SuccessRequestResponse.RequestID)

}

func TestCreateInvalidGetRequestInvalidUUID(t *testing.T) {

	requestID := "invalidRequestId"

	req, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotAcceptable, response.Code)

	FailureResponse := failureResponse{}
	json.Unmarshal(response.Body.Bytes(), &FailureResponse)

	if FailureResponse.Error != "need a valid UUID for request ID" {
		t.Errorf("Expected need a valid UUID for request ID error. Got '%s'", FailureResponse.Error)
	}

}

func TestCreateInvalidGetRequestInvalidRequestId(t *testing.T) {

	requestID := "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

	req, _ := http.NewRequest("GET", "/api/v1/requests/"+requestID+"/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	FailureResponse := failureResponse{}
	json.Unmarshal(response.Body.Bytes(), &FailureResponse)

	if FailureResponse.Error != "request not found" {
		t.Errorf("Expected request not found error. Got '%s'", FailureResponse.Error)
	}

}

func TestCreateInvalidGetRequestInvalidURL(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/invalid_requests_path/", nil)
	req.RemoteAddr = "0.0.0.0:80"
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	FailureResponse := failureResponse{}
	json.Unmarshal(response.Body.Bytes(), &FailureResponse)

	if FailureResponse.Error != "" {
		t.Errorf("Expected no error. Got '%s'", FailureResponse.Error)
	}

}

func TestCreateValidPostRequestTimeout(t *testing.T) {

	payload := []byte(`{"stringsToMatch":"teststring1","stringsToMatchIn":"teststring2", "mode": "simple"}`)

	req1, _ := http.NewRequest("POST", "/api/v1/requests/", bytes.NewBuffer(payload))
	req1.RemoteAddr = "0.0.0.0:80"
	response1 := executeRequest(req1)

	checkResponseCode(t, http.StatusOK, response1.Code)

	SuccessResponse1 := successPostRequestResponse{}
	json.Unmarshal(response1.Body.Bytes(), &SuccessResponse1)

	if utils.IsValidUUID(SuccessResponse1.RequestID) == false {
		t.Errorf("Invalid RequestID. Got '%s'", SuccessResponse1.RequestID)
	}

	time.Sleep(time.Duration(a.Conf.RequestTTLInMinutes) * time.Minute)
	time.Sleep(time.Duration(shortTimeOutInSeconds) * time.Second)

	req2, _ := http.NewRequest("GET", "/api/v1/requests/"+SuccessResponse1.RequestID+"/", nil)
	response2 := executeRequest(req2)

	checkResponseCode(t, http.StatusNotFound, response2.Code)

	FailureResponse := failureResponse{}
	json.Unmarshal(response2.Body.Bytes(), &FailureResponse)

	if FailureResponse.Error != "request not found" {
		t.Errorf("Expected request not found. Got '%s'", FailureResponse.Error)
	}

}
