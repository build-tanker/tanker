package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/access"
	"github.com/build-tanker/tanker/pkg/apps"
	"github.com/build-tanker/tanker/pkg/builds"
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/fileserver"
	"github.com/build-tanker/tanker/pkg/orgs"
	"github.com/build-tanker/tanker/pkg/shippers"
)

// Handler exposes all handlers
type Handler struct {
	health  *healthHandler
	build   *buildHandler
	shipper *shipperHandler
	app     *appHandler
	org     *orgHandler
}

// HTTPHandler is the type which can handle a URL
type httpHandler func(w http.ResponseWriter, r *http.Request)

// New creates a new handler
func New(conf *config.Config, db *sqlx.DB) *Handler {
	// Crete FileServer
	fileServer := fileserver.New(conf, fileserver.NewGoogleCloudStorage())

	// Create services
	buildService := builds.New(conf, db, fileServer)
	shipperService := shippers.New(conf, db)
	accessService := access.New(conf, db)
	appService := apps.New(conf, db, accessService)
	orgService := orgs.New(conf, db, accessService)

	// Finally, create handlers
	health := newHealthHandler()
	build := newBuildHandler(buildService)
	shipper := newShipperHandler(shipperService)
	app := newAppHandler(appService)
	org := newOrgHandler(orgService)

	return &Handler{health, build, shipper, app, org}
}

// Route pipes requests to the correct handlers
func (h *Handler) Route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.health.ping()).Methods(http.MethodGet)

	router.HandleFunc("/v1/org", h.org.Add()).Methods(http.MethodPost)

	router.HandleFunc("/v1/builds", h.build.Add()).Methods(http.MethodPost)

	router.HandleFunc("/v1/shippers", h.shipper.Add()).Methods(http.MethodPost)
	router.HandleFunc("/v1/shippers", h.shipper.ViewAll()).Methods(http.MethodGet)
	router.HandleFunc("/v1/shippers/{id}", h.shipper.View()).Methods(http.MethodGet)
	router.HandleFunc("/v1/shippers/{id}", h.shipper.View()).Methods(http.MethodDelete)

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
