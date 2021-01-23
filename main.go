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

	fmrequest := controller.CreateFuzzyMatchRequest(
		controller.SplitFormStringValueToSliceOfStrings(r.FormValue("stringsToMatch")),
		controller.SplitFormStringValueToSliceOfStrings(r.FormValue("stringsToMatchIn")),
		r.FormValue("mode"))
	// curl sample request:
	// curl -d "stringsToMatch='apple gmbh corp', 'bear'&stringsToMatchIn='apple inc', 'apple gmbh', 'hair'&mode=deepDive" -X POST http://localhost:8080/api/v1/requests/

	fmDAO := model.CreateFuzzyMatchDAO(fmrequest.RequestID, fmrequest.StringsToMatch, fmrequest.StringsToMatchIn, fmrequest.Mode)

	model.RequestsData = append(model.RequestsData, fmDAO)

	fmresponse := controller.CreateFuzzyMatchResponse(fmrequest.RequestID)
	//w.Write([]byte(fmt.Sprintf(`{"result": %s }`, fmresponse)))
	fmt.Println(model.RequestsData)
	//w.Write([]byte(fmt.Sprint(fmresponse.RequestID)))
	fmt.Fprintf(w, "%+v", fmresponse)
}

func getLazy(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	// curl sample request curl -X GET http://localhost:8080/api/v1/requests/66e3a79e-a05e-4d63-856f-5a12ed673965/

	requestID := ""
	//var err error
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		fmt.Println(requestID)
		// requestID, err = strconv.Atoi(val)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(`{"message": "need a number"}`))
		// 	return
		// }
	}

	var fuzzyMatchResultsResponse controller.FuzzyMatchResultsResponse
	var returnedRowsUpperBound int
	var returnedAllRows bool

	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {

			fuzzyMatchResultsResponse = controller.FuzzyMatchResultsResponse{
				RequestID:   requestID,
				Mode:        model.RequestsData[i].Mode,
				RequestedOn: model.RequestsData[i].RequestedOn}

			if model.RequestsData[i].ReturnedRows+model.RequestsData[i].BatchSize >= model.RequestsData[i].StringsToMatchLength {
				returnedRowsUpperBound = model.RequestsData[i].StringsToMatchLength
				returnedAllRows = true
			} else {
				returnedRowsUpperBound = model.RequestsData[i].ReturnedRows + model.RequestsData[i].BatchSize
				returnedAllRows = false
			}

			for stringToMatch := model.RequestsData[i].ReturnedRows; stringToMatch < returnedRowsUpperBound; stringToMatch++ {
				var auxiliaryMatchResults []controller.AuxiliaryMatchResult

				for stringToMatchIn := 0; stringToMatchIn < model.RequestsData[i].StringsToMatchInLength; stringToMatchIn++ {
					auxiliaryMatchResult := controller.AuxiliaryMatchResult{
						StringMatched: model.RequestsData[i].StringsToMatchIn[stringToMatchIn],
						Result: fm.FuzzyMatch(
							model.RequestsData[i].StringsToMatch[stringToMatch],
							model.RequestsData[i].StringsToMatchIn[stringToMatchIn],
							model.RequestsData[i].Mode)}

					auxiliaryMatchResults = append(auxiliaryMatchResults, auxiliaryMatchResult)
				}

				sort.SliceStable(auxiliaryMatchResults, func(i, j int) bool {
					return auxiliaryMatchResults[i].Result > auxiliaryMatchResults[j].Result
				})

				fuzzyMatchResult := controller.FuzzyMatchResult{
					StringToMatch: model.RequestsData[i].StringsToMatch[stringToMatch],
					StringMatched: auxiliaryMatchResults[0].StringMatched,
					Result:        auxiliaryMatchResults[0].Result}

				fuzzyMatchResultsResponse.Results = append(fuzzyMatchResultsResponse.Results, fuzzyMatchResult)
				fuzzyMatchResultsResponse.ReturnedAllRows = returnedAllRows
			}

			// TODO convert DAO to Response with CreateFuzzyMatchResultsResponse transf.function
			if returnedAllRows == true {
				model.DeleteFuzzyMatchDAOInRequestsData(requestID)
			} else {
				model.UpdateFuzzyMatchDAOInRequestsData(requestID, returnedRowsUpperBound)
			}
		}
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
