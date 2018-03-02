package server

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/builds"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/pings"
	"github.com/build-tanker/tanker/pkg/shippers"
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

	// Auth
	// POST .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123
	// GET .../v1/users/15
	// PUT .../v1/users/15 access_token=tkn&name=name&deleted=true
	// DELETE .../v1/users/15

	// AppGroup
	// POST .../v1/appGroup name=name
	// GET .../v1/appGroup/15
	// PUT .../v1/appGroup/15 name=name
	// DELETE .../v1/appGroup/15

	// App
	// POST ../v1/app appGroup=appGroup&name=name&bundleId=bundle_id&platform=platform
	// GET .../v1/app/15
	// PUT .../v1/app/15 name=name
	// DELETE .../v1/app/15

	// Access
	// POST ../v1/access person=person&appGroup=app_group&app=app&access_level=admin&access_given_by=person
	// GET ../v1/access/15
	// PUT ../v1/access/15 access_level=admin
	// DELETE ../v1/access/15

	// Shipper
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

	// Build
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
