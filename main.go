package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	fm "github.com/datahappy1/go_fuzzymatch"
	"github.com/datahappy1/go_fuzzymatch_webapp/controller"
	"github.com/datahappy1/go_fuzzymatch_webapp/model"
	"github.com/gorilla/mux"
)

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var fuzzyMatchExternalRequest controller.FuzzyMatchExternalRequest

	requestBodyString, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	xes := string(requestBodyString[:])
	fmt.Println("xew", xes)
	errx := json.Unmarshal(requestBodyString, &fuzzyMatchExternalRequest)
	if errx != nil {
		log.Printf("error decoding response: %v", errx)
		if e, ok := errx.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", requestBodyString)
	}
	// if err != nil {
	// 	log.Println(err)
	// }

	fmt.Println(fuzzyMatchExternalRequest)
	fuzzyMatchRequest := controller.CreateFuzzyMatchRequest(
		controller.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatch),
		controller.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatchIn),
		fuzzyMatchExternalRequest.Mode)

	// curl sample request:
	// https://stackoverflow.com/questions/11834238/curl-post-command-line-on-windows-restful-service
	// curl -X POST -d "{""stringsToMatch"":""'apple, gmbh','corp'"",""stringsToMatchIn"":""hair"",""mode"":""combined""}" http://localhost:8080/api/v1/requests/

	//	curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
	//	--header 'Content-Type: text/plain' \
	//	--data-raw '{
	//	"stringsToMatch": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.",
	//		"stringsToMatchIn": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.'\''",
	//		"mode": "combined"
	//}'

	model.CreateFuzzyMatchDAOInRequestsData(fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch,
		fuzzyMatchRequest.StringsToMatchIn, fuzzyMatchRequest.Mode)

	fuzzyMatchRequestResponse := controller.CreateFuzzyMatchResponse(fuzzyMatchRequest.RequestID)

	// fmt.Fprintf(w, "%+v", fuzzyMatchRequestResponse)
	jData, err := json.Marshal(fuzzyMatchRequestResponse)
	if err != nil {
		// handle error
	}

	// fmt.Println(model.RequestsData)
	w.Write(jData)

}

func getLazy(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	// curl sample request curl -X GET http://localhost:8080/api/v1/requests/66e3a79e-a05e-4d63-856f-5a12ed673965/

	requestID := ""
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		if controller.IsValidUUID(val) == false {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a valid request UUID"}`))
			return
		}
	}

	var fuzzyMatchDAO model.FuzzyMatchDAO
	var fuzzyMatchResults controller.FuzzyMatchResults
	var fuzzyMatchResultsResponse controller.FuzzyMatchResultsResponse
	var returnedRowsUpperBound int
	var returnedAllRows bool

	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			fuzzyMatchDAO = model.CreateFuzzyMatchDAO(requestID,
				model.RequestsData[i].StringsToMatch,
				model.RequestsData[i].StringsToMatchIn,
				model.RequestsData[i].Mode,
				model.RequestsData[i].ReturnedRows)
			fmt.Println("hah", fuzzyMatchDAO)
			break
		}
	}

	if fuzzyMatchDAO.RequestID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Request not found"}`))
		return
	}

	if fuzzyMatchDAO.ReturnedRows+fuzzyMatchDAO.BatchSize >= fuzzyMatchDAO.StringsToMatchLength {
		returnedRowsUpperBound = fuzzyMatchDAO.StringsToMatchLength
		returnedAllRows = true
		fmt.Println("kuk")
	} else {
		returnedRowsUpperBound = fuzzyMatchDAO.ReturnedRows + fuzzyMatchDAO.BatchSize
		returnedAllRows = false
		fmt.Println("buk")
	}

	fmt.Println(returnedAllRows, returnedRowsUpperBound)
	fmt.Println(model.RequestsData)

	for stringToMatch := fuzzyMatchDAO.ReturnedRows; stringToMatch < returnedRowsUpperBound; stringToMatch++ {
		var auxiliaryMatchResults []controller.AuxiliaryMatchResult

		for stringToMatchIn := 0; stringToMatchIn < fuzzyMatchDAO.StringsToMatchInLength; stringToMatchIn++ {
			auxiliaryMatchResult := controller.AuxiliaryMatchResult{
				StringMatched: fuzzyMatchDAO.StringsToMatchIn[stringToMatchIn],
				Result: fm.FuzzyMatch(
					fuzzyMatchDAO.StringsToMatch[stringToMatch],
					fuzzyMatchDAO.StringsToMatchIn[stringToMatchIn],
					fuzzyMatchDAO.Mode)}

			auxiliaryMatchResults = append(auxiliaryMatchResults, auxiliaryMatchResult)
		}

		sort.SliceStable(auxiliaryMatchResults, func(i, j int) bool {
			return auxiliaryMatchResults[i].Result > auxiliaryMatchResults[j].Result
		})

		fuzzyMatchResult := controller.FuzzyMatchResult{
			StringToMatch: fuzzyMatchDAO.StringsToMatch[stringToMatch],
			StringMatched: auxiliaryMatchResults[0].StringMatched,
			Result:        auxiliaryMatchResults[0].Result}

		fuzzyMatchResults = append(fuzzyMatchResults, fuzzyMatchResult)
	}

	fuzzyMatchResultsResponse = controller.FuzzyMatchResultsResponse{
		RequestID:       requestID,
		Mode:            fuzzyMatchDAO.Mode,
		RequestedOn:     fuzzyMatchDAO.RequestedOn,
		ReturnedAllRows: returnedAllRows,
		Results:         fuzzyMatchResults}

	if returnedAllRows == true {
		model.DeleteFuzzyMatchDAOInRequestsData(requestID)
	} else {
		model.UpdateFuzzyMatchDAOInRequestsData(requestID, returnedRowsUpperBound)
	}

	// fmt.Fprintf(w, "%+v", fuzzyMatchResultsResponse)
	jData, err := json.Marshal(fuzzyMatchResultsResponse)
	if err != nil {
		// handle error
	}

	w.Write(jData)

}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/requests/{requestID}/", getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", post).Methods(http.MethodPost)

	static := r.PathPrefix("").Subrouter()
	fileServer := http.FileServer(http.Dir("./static"))
	static.Handle("/", http.StripPrefix("/", fileServer))

	log.Fatal(http.ListenAndServe(":8080", r))
}
