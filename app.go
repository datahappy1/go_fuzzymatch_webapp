package main

import (
	"encoding/json"
	"errors"
	"fmt"
	fm "github.com/datahappy1/go_fuzzymatch/pkg"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/config"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/repository"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// App returns struct
type App struct {
	Router *mux.Router
	conf   config.Configuration
}

// Initialize App
func (a *App) Initialize(environment string) {
	var err error
	a.conf, err = config.GetConfiguration(environment)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

// Run App
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
}

func (a *App) post(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, a.conf.MaxRequestByteSize)

	if repository.CountAll() >= a.conf.MaxActiveRequestsCount {
		respondWithError(w, http.StatusTooManyRequests, errors.New("too many overall requests in flight, try later"))
		return
	}

	requestedFromIP, err := utils.GetIP(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errors.New("cannot determine IP address"))
		return
	}

	inFlightRequestID := repository.GetByIP(requestedFromIP).RequestID
	if inFlightRequestID != "" {
		respondWithError(w, http.StatusTooManyRequests,
			errors.New(fmt.Sprintf("too many requests from IP address, collect request %s data first", inFlightRequestID)))
		return
	}

	var fuzzyMatchExternalRequest model.FuzzyMatchExternalRequest

	requestBodyString, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, errors.New("cannot read request body"))
		return
	}

	err = json.Unmarshal(requestBodyString, &fuzzyMatchExternalRequest)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		respondWithError(w, http.StatusInternalServerError, errors.New("error decoding request data"))
		return
	}

	fuzzyMatchRequest, err := model.CreateFuzzyMatchRequest(
		utils.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatch),
		utils.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatchIn),
		fuzzyMatchExternalRequest.Mode, requestedFromIP)

	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, errors.New("error invalid request"))
		return
	}

	repository.Create(fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch,
		fuzzyMatchRequest.StringsToMatchIn, fuzzyMatchRequest.Mode, fuzzyMatchRequest.RequestedFromIP, a.conf.BatchSize)

	fuzzyMatchRequestResponse := model.CreateFuzzyMatchResponse(fuzzyMatchRequest.RequestID)

	respondWithJSON(w, http.StatusOK, fuzzyMatchRequestResponse)
}

func (a *App) getLazy(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)

	requestID := ""
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		if utils.IsValidUUID(val) == false {
			respondWithError(w, http.StatusInternalServerError, errors.New("need a valid UUID for request ID"))
			return
		}
	}

	var fuzzyMatchObject model.FuzzyMatchModel
	var fuzzyMatchResults model.FuzzyMatchResults
	var fuzzyMatchResultsResponse model.FuzzyMatchResultsResponse
	var returnedRowsUpperBound int
	var returnedAllRows bool

	fuzzyMatchObject = repository.GetByRequestID(requestID)

	if fuzzyMatchObject.RequestID == "" {
		respondWithError(w, http.StatusNotFound, errors.New("request not found"))
		return
	}

	if fuzzyMatchObject.ReturnedRows+fuzzyMatchObject.BatchSize >= fuzzyMatchObject.StringsToMatchLength {
		returnedRowsUpperBound = fuzzyMatchObject.StringsToMatchLength
		returnedAllRows = true
	} else {
		returnedRowsUpperBound = fuzzyMatchObject.ReturnedRows + fuzzyMatchObject.BatchSize
		returnedAllRows = false
	}

	for stringToMatch := fuzzyMatchObject.ReturnedRows; stringToMatch < returnedRowsUpperBound; stringToMatch++ {
		var auxiliaryMatchResults []model.AuxiliaryMatchResult

		for stringToMatchIn := 0; stringToMatchIn < fuzzyMatchObject.StringsToMatchInLength; stringToMatchIn++ {
			auxiliaryMatchResult := model.AuxiliaryMatchResult{
				StringMatched: fuzzyMatchObject.StringsToMatchIn[stringToMatchIn],
				Result: fm.FuzzyMatch(
					fuzzyMatchObject.StringsToMatch[stringToMatch],
					fuzzyMatchObject.StringsToMatchIn[stringToMatchIn],
					fuzzyMatchObject.Mode)}

			auxiliaryMatchResults = append(auxiliaryMatchResults, auxiliaryMatchResult)
		}

		sort.SliceStable(auxiliaryMatchResults, func(i, j int) bool {
			return auxiliaryMatchResults[i].Result > auxiliaryMatchResults[j].Result
		})

		fuzzyMatchResult := model.FuzzyMatchResult{
			StringToMatch: fuzzyMatchObject.StringsToMatch[stringToMatch],
			StringMatched: auxiliaryMatchResults[0].StringMatched,
			Result:        auxiliaryMatchResults[0].Result}

		fuzzyMatchResults = append(fuzzyMatchResults, fuzzyMatchResult)
	}

	fuzzyMatchResultsResponse = model.CreateFuzzyMatchResultsResponse(
		requestID, fuzzyMatchObject.Mode, fuzzyMatchObject.RequestedOn, returnedAllRows, fuzzyMatchResults,
	)

	if returnedAllRows == true {
		repository.Delete(requestID)
	} else {
		repository.Update(requestID, returnedRowsUpperBound)
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
	w.Write(response)

}
