package api

import (
	"net/http"
)

func (a *App) initializeRoutes() {
	apiRouter := a.Router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/requests/{requestID}/", a.getLazy).Methods(http.MethodGet)
	apiRouter.HandleFunc("/requests/", a.post).Methods(http.MethodPost)

	uiRouter := a.Router.PathPrefix("/").Subrouter()
	fileServerStaticRoot := http.FileServer(http.Dir("./ui/dist/"))
	uiRouter.PathPrefix("/").Handler(fileServerStaticRoot)
}
