package shippers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/tanker/pkg/responses"

	"github.com/jmoiron/sqlx"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type Handler interface {
	Add(*appcontext.AppContext) HTTPHandler
	ViewAll(ctx *appcontext.AppContext) HTTPHandler
	View(ctx *appcontext.AppContext) HTTPHandler
	Delete(ctx *appcontext.AppContext) HTTPHandler
}

type handler struct {
	service Service
}

func NewHandler(ctx *appcontext.AppContext, db *sqlx.DB) Handler {
	return &handler{
		service: NewService(ctx, db),
	}
}

func (s *handler) Add(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		name := s.parseKeyFromQuery(r, "name")
		machineName := s.parseKeyFromQuery(r, "machineName")

		id, accessKey, err := s.service.Add(name, machineName)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: &Shipper{
				ID:        id,
				AccessKey: accessKey,
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
func (s *handler) ViewAll(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// page := s.parseKeyFromQuery(r, "page")
		// count := s.parseKeyFromQuery(r, "count")

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
func (s *handler) View(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := s.parseKeyFromVars(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:view:notFound", errors.New("Could not find id in the request").Error()))
			return
		}

		shippers, err := s.service.View(int64(id))
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
func (s *handler) Delete(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := s.parseKeyFromVars(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:delete:notFound", errors.New("Could not find id in the request").Error()))
			return
		}

		err = s.service.Delete(int64(id))
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:delete:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Success: "true",
		})
	}
}
