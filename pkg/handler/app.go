package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/apps"
	"github.com/build-tanker/tanker/pkg/common/responses"
)

type appHandler struct {
	service *apps.Service
}

func newAppHandler(service *apps.Service) *appHandler {
	return &appHandler{service}
}

func (a *appHandler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		org := parseKeyFromQuery(r, "org")
		name := parseKeyFromQuery(r, "name")
		bundleID := parseKeyFromQuery(r, "bundle_id")
		platform := parseKeyFromQuery(r, "platform")

		id, err := a.service.Add(org, name, bundleID, platform)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("app:add:error", err.Error()))
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
