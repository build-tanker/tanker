package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/common/config"
)

// Handler exposes all handlers
type Handler struct {
	health *healthHandler
}

// HTTPHandler is the type which can handle a URL
type httpHandler func(w http.ResponseWriter, r *http.Request)

// New creates a new handler
func New(conf *config.Config, db *sqlx.DB) *Handler {

	// Finally, create handlers
	health := newHealthHandler()

	return &Handler{health}
}

// Route pipes requests to the correct handlers
func (h *Handler) Route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.health.ping()).Methods(http.MethodGet)

	// GET___ .../ping
	// router.HandleFunc("/ping", pingHandler.Ping(ctx)).Methods(http.MethodGet)

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

func fakeHandler() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}

func parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
