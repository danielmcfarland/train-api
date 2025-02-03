package routes

import (
	"net/http"
)

func LoadRoutes(router *http.ServeMux) {
	handler := &Handler{}

	router.HandleFunc("/api/v1/stations/{id}", handler.getStation)
}
