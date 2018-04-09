package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/builds"
	"github.com/build-tanker/tanker/pkg/common/responses"
)

type buildHandler struct {
	service *builds.Service
}

func newBuildHandler(service *builds.Service) *buildHandler {
	return &buildHandler{service}
}

func (b *buildHandler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := parseKeyFromQuery(r, "file")
		shipper := parseKeyFromQuery(r, "shipper")
		bundle := parseKeyFromQuery(r, "bundle")
		platform := parseKeyFromQuery(r, "platform")
		extension := parseKeyFromQuery(r, "extension")

		url, err := b.service.Add(fileName, shipper, bundle, platform, extension)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("build:add:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: struct {
				URL string `json:"url"`
			}{
				URL: url,
			},
			Success: "true",
		})
	}
}
