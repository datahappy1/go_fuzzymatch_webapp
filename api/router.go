package api

import (
	"net/http"
)

func (a *App) initializeRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/requests/{requestID}/", a.getLazy).Methods(http.MethodGet)
	api.HandleFunc("/requests/", a.post).Methods(http.MethodPost)

	ui := a.Router.PathPrefix("/").Subrouter()
	fileServerStaticRoot := http.FileServer(http.Dir("./ui/dist/"))
	ui.PathPrefix("/").Handler(fileServerStaticRoot)
}
