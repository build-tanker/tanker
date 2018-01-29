package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/pings"
	"github.com/sudhanshuraheja/tanker/pkg/shippers"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	pingService := pings.PingHandler{}

	shipperService := shippers.ShipperHandler{}
	shipperService.Init(ctx, db)

	router := mux.NewRouter()
	// GET___ .../ping
	router.HandleFunc("/ping", pingService.Ping(ctx))

	// Shippers
	// POST__ .../v1/shippers?name=shipper_name&machineName=machine_name
	router.HandleFunc("/v1/shippers", shipperService.Add(ctx)).Methods("POST")
	// GET___ .../v1/shippers?page=1&count=25
	router.HandleFunc("/v1/shippers", FakeHandler(ctx, db)).Methods("GET")
	// GET___ .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{id}", FakeHandler(ctx, db)).Methods("GET")
	// PUT___ .../v1/shippers/id?name=shipper_name&machineName=machine_name
	router.HandleFunc("/v1/shippers/{id}", FakeHandler(ctx, db)).Methods("PUT")
	// DELETE .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{id}", FakeHandler(ctx, db)).Methods("DELETE")

	// Builds
	// POST__ .../v1/builds?accessKey=a1b2c3&buildSize=80&checksum=a1b2c3
	router.HandleFunc("/v1/builds", FakeHandler(ctx, db)).Methods("POST")
	// POST__ .../v1/builds/abcdef?accessKey=a1b2c3
	router.HandleFunc("/v1/builds/{id}", FakeHandler(ctx, db)).Methods("POST")
	return router
}

func FakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
