package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/pings"
	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}

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

	return router
}

func FakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
