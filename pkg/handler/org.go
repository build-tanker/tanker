package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/common/responses"
	"github.com/build-tanker/tanker/pkg/orgs"
)

type orgHandler struct {
	service *orgs.Service
}

func newOrgHandler(service *orgs.Service) *orgHandler {
	return &orgHandler{service}
}

func (a *orgHandler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		name := parseKeyFromQuery(r, "name")
		imageURL := parseKeyFromQuery(r, "image_url")

		id, err := a.service.Add(name, imageURL)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("org:add:error", err.Error()))
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
