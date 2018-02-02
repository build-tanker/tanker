package builds

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"source.golabs.io/core/tanker/pkg/appcontext"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type Handler interface {
	Add() HTTPHandler
}

type handler struct {
	ctx     *appcontext.AppContext
	service Service
}

func NewHandler(ctx *appcontext.AppContext) Handler {
	s := NewService()
	return &handler{
		ctx:     ctx,
		service: s,
	}
}

func (s *handler) Add() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *handler) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (s *handler) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	fmt.Println(vars)
	return vars[key]
}
