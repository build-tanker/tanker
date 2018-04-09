package builds

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/responses"
	"github.com/gorilla/mux"
)

// HTTPHandler - type for http handlers
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// Handler - handles routes for builds
type Handler interface {
	Add() HTTPHandler
}

type handler struct {
	cnf     *config.Config
	service *Service
}

// BuildAddResponse - create a response for adding a build
type BuildAddResponse struct {
	URL string `json:"url"`
}

// NewHandler - creates a new handler for builds
func NewHandler(cnf *config.Config, db *sqlx.DB) Handler {
	b := New(cnf, db)
	return &handler{
		cnf:     cnf,
		service: b,
	}
}

func (b *handler) Add() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := b.parseKeyFromQuery(r, "file")
		shipper := b.parseKeyFromQuery(r, "shipper")
		bundle := b.parseKeyFromQuery(r, "bundle")
		platform := b.parseKeyFromQuery(r, "platform")
		extension := b.parseKeyFromQuery(r, "extension")

		url, err := b.service.Add(fileName, shipper, bundle, platform, extension)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("build:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: &BuildAddResponse{
				URL: url,
			},
			Success: "true",
		})
	}
}

func (b *handler) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (b *handler) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
