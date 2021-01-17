package main

import (
	"fmt"
	"log"
	"net/http"

	fm "github.com/datahappy1/go_fuzzymatch"
	"github.com/datahappy1/go_fuzzymatch_webapp/model"
	"github.com/gorilla/mux"
)

var requests = []model.FuzzyMatchRequest

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	fmrequest := model.CreateFuzzyMatchRequest(
		r.FormValue("stringSToMatch"), 
		r.FormValue("stringsToMatchIn"), 
		r.FormValue("mode"))
	// curl sample request:
	// curl -d "stringsToMatch=['apple gmbh corp', 'bear']&stringsToMatchIn=['apple inc', 'apple gmbh', 'hair']&mode=deepDive" -X POST http://localhost:8080/api/v1
	// curl sample response:
	// {"result": 73 }
	fmresponse := fm.FuzzyMatch(fmrequest.StringsToMatch, fmrequest.StringsToMatchIn, fmrequest.Mode)
	w.Write([]byte(fmt.Sprintf(`{"result": %d }`, fmresponse)))

}

// func params(w http.ResponseWriter, r *http.Request) {
// 	pathParams := mux.Vars(r)
// 	w.Header().Set("Content-Type", "application/json")

// 	userID := -1
// 	var err error
// 	if val, ok := pathParams["userID"]; ok {
// 		userID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	commentID := -1
// 	if val, ok := pathParams["commentID"]; ok {
// 		commentID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	query := r.URL.Query()
// 	location := query.Get("location")

// 	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
// }

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)

	// api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
