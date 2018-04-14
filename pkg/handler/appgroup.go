package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/appgroups"
	"github.com/build-tanker/tanker/pkg/common/responses"
)

type appGroupHandler struct {
	service *appgroups.Service
}

func newAppGroupHandler(service *appgroups.Service) *appGroupHandler {
	return &appGroupHandler{service}
}

func (a *appGroupHandler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		name := parseKeyFromQuery(r, "name")
		imageURL := parseKeyFromQuery(r, "image_url")

		id, err := a.service.Add(name, imageURL)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("appGroup:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: struct {
				ID string `json:"id"`
			}{
				ID: id,
			},
			Success: "true",
		})
	}
}
