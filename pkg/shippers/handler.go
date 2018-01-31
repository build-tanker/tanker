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

type ShipperHandler struct {
	service ShippersService
}

func (s *ShipperHandler) Init(ctx *appcontext.AppContext, db *sqlx.DB) {
	s.service = NewShippersService(ctx, db)
}

func (s *ShipperHandler) Add(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		name := s.parseKeyFromQuery(r, "name")
		machineName := s.parseKeyFromQuery(r, "machineName")

		id, accessKey, err := s.service.Add(name, machineName)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:add:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, responses.NewShipperAddSuccessResponse(id, accessKey))
	}
}

func (s *ShipperHandler) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (s *ShipperHandler) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

// /v1/shippers?page=1&count=25
func (s *ShipperHandler) ViewAll(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// page := s.parseKeyFromQuery(r, "page")
		// count := s.parseKeyFromQuery(r, "count")

		shippers, err := s.service.ViewAll()
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:viewall:error", err.Error()))
			return
		}
		responses.WriteJSON(w, http.StatusOK, responses.NewShipperViewAllSuccessResponse(shippers))
	}
}

// /v1/shippers/id
func (s *ShipperHandler) View(ctx *appcontext.AppContext) HTTPHandler {
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
		responses.WriteJSON(w, http.StatusOK, responses.NewShipperViewSuccessResponse(shippers))
	}
}

// /v1/shippers/id
func (s *ShipperHandler) Delete(ctx *appcontext.AppContext) HTTPHandler {
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
		responses.WriteJSON(w, http.StatusOK, responses.NewShipperDeleteSuccessResponse())
	}
}
