package shippers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/responses"
)

// HTTPHandler - handler incoming requests
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// Handler - handle requests for shippers
type Handler interface {
	Add() HTTPHandler
	ViewAll() HTTPHandler
	View() HTTPHandler
	Delete() HTTPHandler
}

type handler struct {
	cnf     *config.Config
	service Service
}

// NewHandler - create a new handler for shippers
func NewHandler(cnf *config.Config, db *sqlx.DB) Handler {
	return &handler{
		cnf:     cnf,
		service: NewService(cnf, db),
	}
}

func (s *handler) Add() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		appGroup := s.parseKeyFromQuery(r, "appGroup")
		expiry := s.parseKeyFromQuery(r, "expiry")

		expiryInt, err := strconv.Atoi(expiry)
		if err != nil {
			expiryInt = 0
		}

		id, err := s.service.Add(appGroup, expiryInt)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: &Shipper{
				ID: id,
			},
			Success: "true",
		})
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
	return vars[key]
}

// /v1/shippers?page=1&count=25
func (s *handler) ViewAll() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		shippers, err := s.service.ViewAll()
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:viewall:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data:    shippers,
			Success: "true",
		})
	}
}

// /v1/shippers/id
func (s *handler) View() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		id := s.parseKeyFromVars(r, "id")

		shippers, err := s.service.View(id)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:view:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data:    shippers,
			Success: "true",
		})
	}
}

// /v1/shippers/id
func (s *handler) Delete() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		id := s.parseKeyFromVars(r, "id")
		if id == "" {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:delete:notFound", errors.New("Could not find id in the request").Error()))
			return
		}

		err := s.service.Delete(id)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:delete:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Success: "true",
		})
	}
}
