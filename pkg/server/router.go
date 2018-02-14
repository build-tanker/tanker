package server

import (
	"net/http"

	"github.com/gojektech/tanker/pkg/builds"

	"github.com/jmoiron/sqlx"

	"github.com/gojektech/tanker/pkg/appcontext"
	"github.com/gojektech/tanker/pkg/pings"
	"github.com/gojektech/tanker/pkg/shippers"
	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}

	shipperHandler := shippers.NewHandler(ctx, db)
	buildsHandler := builds.NewHandler(ctx, db)

	router := mux.NewRouter()
	// GET___ .../ping
	router.HandleFunc("/ping", pingHandler.Ping(ctx)).Methods(http.MethodGet)

	// Shippers
	// POST__ .../v1/shippers?name=shipper_name&machineName=machine_name
	router.HandleFunc("/v1/shippers", shipperHandler.Add()).Methods(http.MethodPost)
	// GET___ .../v1/shippers?page=1&count=25
	router.HandleFunc("/v1/shippers", shipperHandler.ViewAll()).Methods(http.MethodGet)
	// GET___ .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{id}", shipperHandler.View()).Methods(http.MethodGet)
	// PUT___ .../v1/shippers/id?name=shipper_name&machineName=machine_name
	// router.HandleFunc("/v1/shippers/{id}", FakeHandler(ctx, db)).Methods(http.MethodPut)
	// DELETE .../v1/shippers/id
	router.HandleFunc("/v1/shippers/{accessKey}", shipperHandler.Delete()).Methods(http.MethodDelete)

	// Builds
	// POST__ .../v1/builds?accessKey=a1b2c3&bundle=com.me.app
	router.HandleFunc("/v1/builds", buildsHandler.Add()).Methods(http.MethodPost)
	return router
}

func FakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
