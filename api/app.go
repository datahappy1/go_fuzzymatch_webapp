package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/config"
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

// App returns struct
type App struct {
	Router *mux.Router
	Conf   config.Configuration
}

// ClearAppRequestData returns nil
func (a *App) ClearAppRequestData() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				log.Printf("Checking for timed out requests %s", t)
				timedOutRequestIDs := repository.GetAllTimedOutRequestIDs(a.Conf.RequestTTLInMinutes)

				for i := range timedOutRequestIDs {
					err := repository.Delete(timedOutRequestIDs[i])
					if err != nil {
						log.Printf("cannot delete request %s", timedOutRequestIDs[i])
					} else {
						log.Printf("deleted timed out request %s", timedOutRequestIDs[i])
					}
				}
			}
		}
	}()
}

// Initialize App
func (a *App) Initialize(environment string) {
	var err error
	a.Conf, err = config.GetConfiguration(environment)
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

	r.Body = http.MaxBytesReader(w, r.Body, a.Conf.MaxRequestByteSize)

	if repository.CountAll() >= a.Conf.MaxActiveRequestsCount {
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

	err = json.Unmarshal(requestBodyString, &fuzzyMatchExternalRequest)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		respondWithError(w, http.StatusInternalServerError,
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

	err = repository.Create(fuzzyMatchRequest.RequestID, fuzzyMatchRequest.StringsToMatch,
		fuzzyMatchRequest.StringsToMatchIn, fuzzyMatchRequest.Mode, a.Conf.BatchSize)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			errors.New(fmt.Sprintf("error cannot persist request %s", err)))
		return
	}

	fuzzyMatchRequestResponse := model.CreateFuzzyMatchResponse(fuzzyMatchRequest.RequestID)

	respondWithJSON(w, http.StatusOK, fuzzyMatchRequestResponse)
}

func (a *App) getLazy(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)

	requestID := ""
	if val, ok := pathParams["requestID"]; ok {
		requestID = val
		if utils.IsValidUUID(val) == false {
			respondWithError(w, http.StatusInternalServerError,
				errors.New("need a valid UUID for request ID"))
			return
		}
	}

	fuzzyMatchObject := repository.GetByRequestID(requestID)

	if fuzzyMatchObject.RequestID == "" {
		respondWithError(w, http.StatusNotFound,
			errors.New("request not found"))
		return
	}

	fuzzyMatchResults, returnedAllRows, returnedRowsUpperBound := service.CalculateFuzzyMatchingResults(fuzzyMatchObject)

	if returnedAllRows == true {
		err := repository.Delete(requestID)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError,
				errors.New(fmt.Sprintf("error cannot process request %s", err)))
			return
		}

	} else {
		fuzzyMatchObject.ReturnedRows = returnedRowsUpperBound

		err := repository.Update(requestID, fuzzyMatchObject)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError,
				errors.New(fmt.Sprintf("error cannot process request %s", err)))
			return
		}
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
