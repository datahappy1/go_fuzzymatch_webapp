package main

import (
	//"database/sql"
	//_ "github.com/lib/pq"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	fm "github.com/datahappy1/go_fuzzymatch/pkg"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/config"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/controller"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/gorilla/mux"
)

// App returns struct
type App struct {
	Router *mux.Router
	//DB     *sql.DB
	conf config.Configuration
}

func (a *App) Initialize(environment string) {
	var err error
	a.conf, err = config.GetConfiguration(environment)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

//func (a *App) Initialize(user, password, dbname string) {
//connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

//var err error
//a.DB, err = sql.Open("mysql", connectionString)
//if err != nil {
//	log.Fatal(err)
//}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/requests/{requestID}/", a.getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", a.post).Methods(http.MethodPost)

	ui := a.Router.PathPrefix("/").Subrouter()
	fileServerStaticRoot := http.FileServer(http.Dir("./ui/dist/"))
	ui.PathPrefix("/").Handler(fileServerStaticRoot)
	//ui.PathPrefix("/src/").Handler(fileServerStaticRoot)
	//ui.PathPrefix("/dist/").Handler(fileServerStaticRoot)

}

func (a *App) post(w http.ResponseWriter, r *http.Request) {
	// TODO status for too large request
	r.Body = http.MaxBytesReader(w, r.Body, a.conf.MaxRequestByteSize)

	if model.EvaluateRequestCount(a.conf.MaxActiveRequestsCount) == false {
		respondWithError(w, http.StatusTooManyRequests, errors.New("too many overall requests in flight, try later"))
		return
	}

	requestedFromIP, err := controller.GetIP(r)
	fmt.Println(err)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errors.New("cannot determine IP address"))
		return
	}

	isTooManyRequestRatePerIP, inFlightRequestID := model.EvaluateRequestRatePerIP(requestedFromIP)
	if isTooManyRequestRatePerIP == true {
		respondWithError(w, http.StatusTooManyRequests, errors.New(fmt.Sprintf("too many requests from IP address in flight, "+
			"collect previous request %s data first", inFlightRequestID)))
		return
	}

	var fuzzyMatchExternalRequest controller.FuzzyMatchExternalRequest

	requestBodyString, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, errors.New("cannot read request body"))
		return
	}

	//rawRequestBody := string(requestBodyString[:])
	err = json.Unmarshal(requestBodyString, &fuzzyMatchExternalRequest)
	if err != nil {

		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", requestBodyString)

		respondWithError(w, http.StatusInternalServerError, errors.New("error decoding request data"))
		return
	}

	fuzzyMatchRequest, err := controller.CreateFuzzyMatchRequest(
		controller.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatch),
		controller.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatchIn),
		fuzzyMatchExternalRequest.Mode, requestedFromIP)

	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, errors.New("error invalid request"))
		return
	}

	// curl sample request:
	// https://stackoverflow.com/questions/11834238/curl-post-command-line-on-windows-restful-service
	// Windows cmd:
	// curl -X POST -d "{""stringsToMatch"":""'apple, gmbh','corp'"",""stringsToMatchIn"":""hair"",""mode"":""combined""}" http://localhost:8080/api/v1/requests/

	// *Nix terminal:
	//	curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
	//	--header 'Content-Type: text/plain' \
	//	--data-raw '{
	//	"stringsToMatch": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.",
	//		"stringsToMatchIn": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.'\''",
	//		"mode": "combined"
	//}'

	model.CreateFuzzyMatchDAOInRequestsData(fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch,
		fuzzyMatchRequest.StringsToMatchIn, fuzzyMatchRequest.Mode, fuzzyMatchRequest.RequestedFromIP, a.conf.BatchSize)

	fuzzyMatchRequestResponse := controller.CreateFuzzyMatchResponse(fuzzyMatchRequest.RequestID)

	respondWithJSON(w, http.StatusOK, fuzzyMatchRequestResponse)

}

func (a *App) getLazy(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	// curl sample request
	// curl -X GET http://localhost:8080/api/v1/requests/66e3a79e-a05e-4d63-856f-5a12ed673965/

	requestID := ""
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		if controller.IsValidUUID(val) == false {
			respondWithError(w, http.StatusInternalServerError, errors.New("need a valid UUID for request ID"))
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
				model.RequestsData[i].RequestedFromIP,
				model.RequestsData[i].BatchSize,
				model.RequestsData[i].ReturnedRows)
			break
		}
	}

	if fuzzyMatchDAO.RequestID == "" {
		respondWithError(w, http.StatusNotFound, errors.New("request not found"))
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

	respondWithJSON(w, http.StatusOK, fuzzyMatchResultsResponse)
}

func respondWithError(w http.ResponseWriter, code int, message error) {
	respondWithJSON(w, code, map[string]string{"error": message.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response, _ := json.Marshal(payload)
	// TODO handle write response error
	w.Write(response)

}
