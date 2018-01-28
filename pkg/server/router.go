package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
)

func Router(ctx *appcontext.AppContext) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler)
	return router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":\"pong\"}"))
}
