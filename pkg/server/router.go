package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/pings"
)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {
	router := mux.NewRouter()
	router.Handle("/ping", &pings.PingHandler{})
	return router
}
