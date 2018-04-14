package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/build-tanker/tanker/pkg/common/responses"
	"github.com/build-tanker/tanker/pkg/shippers"
)

type shipperHandler struct {
	service *shippers.Service
}

func newShipperHandler(service *shippers.Service) *shipperHandler {
	return &shipperHandler{service}
}

func (s *shipperHandler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		org := parseKeyFromQuery(r, "org")
		expiry := parseKeyFromQuery(r, "expiry")

		expiryInt, err := strconv.Atoi(expiry)
		if err != nil {
			expiryInt = 0
		}

		id, err := s.service.Add(org, expiryInt)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("shipper:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: &shippers.Shipper{
				ID: id,
			},
			Success: "true",
		})
	}
}

// /v1/shippers?page=1&count=25
func (s *shipperHandler) ViewAll() httpHandler {
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
func (s *shipperHandler) View() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseKeyFromVars(r, "id")

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
func (s *shipperHandler) Delete() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseKeyFromVars(r, "id")
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
