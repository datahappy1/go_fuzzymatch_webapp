package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	fm "github.com/datahappy1/go_fuzzymatch"
	"github.com/datahappy1/go_fuzzymatch_webapp/controller"
	"github.com/datahappy1/go_fuzzymatch_webapp/model"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// here will be served homepage html
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	fuzzyMatchRequest := controller.CreateFuzzyMatchRequest(
		controller.SplitFormStringValueToSliceOfStrings(r.FormValue("stringsToMatch")),
		controller.SplitFormStringValueToSliceOfStrings(r.FormValue("stringsToMatchIn")),
		r.FormValue("mode"))
	// curl sample request:
	// curl -d "stringsToMatch='apple gmbh corp', 'bear'&stringsToMatchIn='apple inc', 'apple gmbh', 'hair'&mode=deepDive" -X POST http://localhost:8080/api/v1/requests/

	fuzzyMatchDAO := model.CreateFuzzyMatchDAO(fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch, fuzzyMatchRequest.StringsToMatchIn, fuzzyMatchRequest.Mode)

	model.RequestsData = append(model.RequestsData, fuzzyMatchDAO)

	fuzzyMatchRequestResponse := controller.CreateFuzzyMatchResponse(fuzzyMatchRequest.RequestID)
	//w.Write([]byte(fmt.Sprintf(`{"result": %s }`, fmresponse)))
	fmt.Println(model.RequestsData)
	//w.Write([]byte(fmt.Sprint(fmresponse.RequestID)))
	fmt.Fprintf(w, "%+v", fuzzyMatchRequestResponse)
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
			fuzzyMatchDAO = model.CreateFuzzyMatchDAO(requestID, model.RequestsData[i].StringsToMatch, model.RequestsData[i].StringsToMatchIn, model.RequestsData[i].Mode)
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
	} else {
		returnedRowsUpperBound = fuzzyMatchDAO.ReturnedRows + fuzzyMatchDAO.BatchSize
		returnedAllRows = false
	}

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

	fmt.Fprintf(w, "%+v", fuzzyMatchResultsResponse)
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("/requests/{requestID}/", getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", post).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
