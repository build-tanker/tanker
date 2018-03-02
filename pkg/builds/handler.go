package builds

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/responses"
	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type Handler interface {
	Add() HTTPHandler
}

type handler struct {
	ctx     *appcontext.AppContext
	service Service
}

type BuildAddResponse struct {
	URL string `json:"url"`
}

func NewHandler(ctx *appcontext.AppContext, db *sqlx.DB) Handler {
	b := NewService(ctx, db)
	return &handler{
		ctx:     ctx,
		service: b,
	}
}

func (b *handler) Add() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		accessKey := b.parseKeyFromQuery(r, "accessKey")
		bundleID := b.parseKeyFromQuery(r, "bundle")

		url, err := b.service.Add(accessKey, bundleID)
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
