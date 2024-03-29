package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/config"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/data"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/repository"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/service"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type repo struct {
	DB data.Database
}

// App is struct
type App struct {
	Router *mux.Router
	Conf   config.Configuration
	Repo   repo
}

// InitializeAPI returns nil
func (a *App) InitializeAPI(environment string) {
	var err error
	a.Conf, err = config.GetConfiguration(environment)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeAPIRoutes()
}

// InitializeStatic returns nil
func (a *App) InitializeStatic() {
	a.initializeStaticRoutes()
}

// InitializeDB returns nil
func (a *App) InitializeDB() {
	requestsPseudoTable := data.CreateRequestsPseudoTable()
	a.Repo.DB = data.Database{RequestsPseudoTable: requestsPseudoTable}
}

// ClearAppRequestData returns nil
func (a *App) ClearAppRequestData() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				log.Printf("Checking for timed out requests %s", t)
				timedOutRequestIDs := repository.GetAllTimedOutRequestIDs(a.Repo.DB, a.Conf.RequestTTLInMinutes)

				for i := range timedOutRequestIDs {
					repository.Delete(a.Repo.DB, timedOutRequestIDs[i])
					log.Printf("deleted timed out request %s", timedOutRequestIDs[i])
				}
			}
		}
	}()
}

// Run App returns nil
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) post(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, a.Conf.MaxRequestByteSize)

	if repository.CountAll(a.Repo.DB) >= a.Conf.MaxActiveRequestsCount {
		respondWithError(w, http.StatusTooManyRequests,
			errors.New("too many overall requests in flight, try later"))
		return
	}

	var fuzzyMatchExternalRequest model.FuzzyMatchExternalRequest

	requestBodyString, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable,
			errors.New("cannot read request body"))
		return
	}

	//fmt.Println(string(requestBodyString[:]))
	err = json.Unmarshal(requestBodyString, &fuzzyMatchExternalRequest)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		respondWithError(w, http.StatusNotAcceptable,
			errors.New("error decoding request data"))
		return
	}

	fuzzyMatchRequest, err := model.CreateFuzzyMatchRequest(
		utils.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatch),
		utils.SplitFormStringValueToSliceOfStrings(fuzzyMatchExternalRequest.StringsToMatchIn),
		fuzzyMatchExternalRequest.Mode)

	if err != nil {
		respondWithError(w, http.StatusNotAcceptable,
			errors.New("error invalid request"))
		return
	}

	fuzzyMatchObject, err := model.CreateFuzzyMatch(
		fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch, fuzzyMatchRequest.StringsToMatchIn,
		fuzzyMatchRequest.Mode, a.Conf.BatchSize, 0)
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, err)
		return
	}

	repository.Create(a.Repo.DB, fuzzyMatchObject)

	fuzzyMatchRequestResponse := model.CreateFuzzyMatchResponse(fuzzyMatchObject.RequestID)

	respondWithJSON(w, http.StatusOK, fuzzyMatchRequestResponse)
}

func (a *App) getLazy(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	requestID := ""
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		if utils.IsValidUUID(val) == false {
			respondWithError(w, http.StatusNotAcceptable,
				errors.New("need a valid UUID for request ID"))
			return
		}
	}

	fuzzyMatchObject := repository.GetByRequestID(a.Repo.DB, requestID)

	if fuzzyMatchObject.RequestID == "" {
		respondWithError(w, http.StatusNotFound,
			errors.New("request not found"))
		return
	}

	fuzzyMatchResults, returnedAllRows, returnedRowsUpperBound := service.CalculateFuzzyMatchingResults(fuzzyMatchObject)

	if returnedAllRows == true {
		repository.Delete(a.Repo.DB, requestID)

	} else {
		fuzzyMatchObject.ReturnedRows = returnedRowsUpperBound
		repository.Update(a.Repo.DB, fuzzyMatchObject)
	}

	fuzzyMatchResultsResponse := model.CreateFuzzyMatchResultsResponse(
		requestID, fuzzyMatchObject.Mode, fuzzyMatchObject.RequestedOn, returnedAllRows, fuzzyMatchResults,
	)

	respondWithJSON(w, http.StatusOK, fuzzyMatchResultsResponse)
}

func respondWithError(w http.ResponseWriter, code int, message error) {
	respondWithJSON(w, code, map[string]string{"error": message.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response, _ := json.Marshal(payload)
	_, err := w.Write(response)

	if err != nil {
		fmt.Println(err)
	}

}
