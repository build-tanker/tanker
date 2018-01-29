package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/pings"
	"github.com/sudhanshuraheja/tanker/pkg/shippers"
)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	shipperService := shippers.ShipperHandler{}

	router := mux.NewRouter()
	// GET___ .../ping
	router.Handle("/ping", &pings.PingHandler{})

	// Shippers
	// POST__ .../v1/shippers?name=shipper_name&machineName=machine_name
	router.HandleFunc("/v1/shippers", shipperService.Add()).Methods("POST")
	// GET___ .../v1/shippers?page=1&count=10
	router.HandleFunc("/v1/shippers", FakeHandler).Methods("GET")
	// GET___ .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{id}", FakeHandler).Methods("GET")
	// PUT___ .../v1/shippers/id?name=shipper_name&machineName=machine_name
	router.HandleFunc("/v1/shippers/{id}", FakeHandler).Methods("PUT")
	// DELETE .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{id}", FakeHandler).Methods("DELETE")

	// Builds
	// POST__ .../v1/builds?accessKey=a1b2c3&buildSize=80&checksum=a1b2c3
	router.HandleFunc("/v1/builds", FakeHandler).Methods("POST")
	// POST__ .../v1/builds/abcdef?accessKey=a1b2c3
	router.HandleFunc("/v1/builds/{id}", FakeHandler).Methods("POST")
	return router
}

func FakeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"fake\":\"yes\"}"))
}
