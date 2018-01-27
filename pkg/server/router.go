package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router : route requests to handlers
func Router() http.Handler {
	router := mux.NewRouter()
	router.Handle("/ping", pingHandler())
	return router
}

func pingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"success\":\"pong\"}"))
	})
}
