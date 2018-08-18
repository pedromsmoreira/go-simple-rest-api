package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func CreateRoutes(hch *HealthCheckHandler) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/healthchecks/shallow").
		Name("Shallow Healthchecks").
		HandlerFunc(hch.Shallow)

	return router
}
