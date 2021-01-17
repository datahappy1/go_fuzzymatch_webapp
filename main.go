package main

import (
	"fmt"
	"log"
	"net/http"

	fm "github.com/datahappy1/go_fuzzymatch"
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

	fmrequest := model.CreateFuzzyMatchRequest(
		model.SplitFormStringValueToArrayOfStrings(r.FormValue("stringsToMatch")),
		model.SplitFormStringValueToArrayOfStrings(r.FormValue("stringsToMatchIn")),
		r.FormValue("mode"))
	// curl sample request:
	// curl -d "stringsToMatch='apple gmbh corp', 'bear'&stringsToMatchIn='apple inc', 'apple gmbh', 'hair'&mode=deepDive" -X POST http://localhost:8080/api/v1/requests/
	requests = append(requests, fmrequest)

	fmresponse := model.CreateFuzzyMatchResponse(fmrequest.RequestID)
	//w.Write([]byte(fmt.Sprintf(`{"result": %s }`, fmresponse)))
	//fmt.Println(requests)
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

	// first read values from requests variable by requestID and use in fm.FuzzyMatch
	//result := []string{}
	var fmResultsResponse model.FuzzyMatchResultsResponse

	for i := range requests {
		fmt.Println(requests[i])
		if requests[i].RequestID == requestID {
			res := fm.FuzzyMatch(
				requests[i].StringsToMatch[0],
				requests[i].StringsToMatchIn[0],
				requests[i].Mode)

			fmt.Println(res)

			fmResultsResponse = model.FuzzyMatchResultsResponse{
				RequestID: requestID,
				Mode:      requests[i].Mode}
			// Results:   res} // need to match model.FuzzyMatchResultsResponse
		}
	}

	// fmresponse := fm.FuzzyMatch(fmrequest.StringsToMatch, fmrequest.StringsToMatchIn, fmrequest.Mode)
	fmt.Fprintf(w, "%+v", fmResultsResponse)

	// query := r.URL.Query()
	// location := query.Get("location")

	// w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

var requests []model.FuzzyMatchRequest

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("/requests/{requestID}/", getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", post).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
