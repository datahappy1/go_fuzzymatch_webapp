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

	requests = append(requests, fmDAO)

	fmresponse := controller.CreateFuzzyMatchResponse(fmrequest.RequestID)
	//w.Write([]byte(fmt.Sprintf(`{"result": %s }`, fmresponse)))
	fmt.Println(requests)
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
	var _retRows int

	for i := range requests {
		if requests[i].RequestID == requestID {

			if requests[i].StringsToMatchLength <= requests[i].BatchSize || requests[i].ReturnedRows >= requests[i].StringsToMatchLength {
				fuzzyMatchResultsResponse.ReturnedAllRows = true
				// TODO delete item with func from dao
				fmt.Fprintf(w, "%+v", fuzzyMatchResultsResponse)
				break
			}

			fuzzyMatchResultsResponse = controller.FuzzyMatchResultsResponse{
				RequestID:   requestID,
				Mode:        requests[i].Mode,
				RequestedOn: requests[i].RequestedOn}

			fmt.Println(">>>", requests[i])

			var upto int
			if requests[i].ReturnedRows+requests[i].BatchSize >= len(requests[i].StringsToMatchIn) {
				upto = len(requests[i].StringsToMatchIn)
			} else {
				upto = requests[i].ReturnedRows + requests[i].BatchSize
			}

			//for stringToMatch := 0; stringToMatch < len(requests[i].StringsToMatch); stringToMatch++ {
			//for stringToMatch := requests[i].ReturnedRows; stringToMatch < requests[i].ReturnedRows + requests[i].BatchSize; stringToMatch++ {
			for stringToMatch := requests[i].ReturnedRows; stringToMatch < upto; stringToMatch++ {
				fmt.Println("<><><>", stringToMatch, requests[i].ReturnedRows+requests[i].BatchSize)
				var auxiliaryMatchResults []controller.AuxiliaryMatchResult

				for stringToMatchIn := 0; stringToMatchIn < len(requests[i].StringsToMatchIn); stringToMatchIn++ {
					//fmt.Println(requests[i].StringsToMatch[stringToMatch], requests[i].StringsToMatchIn[stringToMatchIn])
					auxiliaryMatchResult := controller.AuxiliaryMatchResult{
						StringMatched: requests[i].StringsToMatchIn[stringToMatchIn],
						Result: fm.FuzzyMatch(
							requests[i].StringsToMatch[stringToMatch],
							requests[i].StringsToMatchIn[stringToMatchIn],
							requests[i].Mode)}

					auxiliaryMatchResults = append(auxiliaryMatchResults, auxiliaryMatchResult)

				}

				sort.SliceStable(auxiliaryMatchResults, func(i, j int) bool {
					return auxiliaryMatchResults[i].Result > auxiliaryMatchResults[j].Result
				})

				fuzzyMatchResult := controller.FuzzyMatchResult{
					StringToMatch: requests[i].StringsToMatch[stringToMatch],
					StringMatched: auxiliaryMatchResults[0].StringMatched,
					Result:        auxiliaryMatchResults[0].Result}

				fuzzyMatchResultsResponse.Results = append(fuzzyMatchResultsResponse.Results, fuzzyMatchResult)

				_retRows = requests[i].ReturnedRows + requests[i].BatchSize
			}
		}
	}

	model.UpdateFuzzyMatchDAO(requests, requestID, _retRows)

	// TODO convert DAO to Response with CreateFuzzyMatchResultsResponse transf.function
	fmt.Fprintf(w, "%+v", fuzzyMatchResultsResponse)

	// query := r.URL.Query()
	// location := query.Get("location")

	// w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

var requests []model.FuzzyMatchDAO

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("/requests/{requestID}/", getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", post).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
