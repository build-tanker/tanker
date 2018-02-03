package builds

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/responses"
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

func NewHandler(ctx *appcontext.AppContext) Handler {
	b := NewService(ctx)
	return &handler{
		ctx:     ctx,
		service: b,
	}
}

func (b *handler) Add() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		accessKey := b.parseKeyFromQuery(r, "accessKey")
		bundleID := b.parseKeyFromQuery(r, "bundleID")

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
	fmt.Println(vars)
	return vars[key]
}
