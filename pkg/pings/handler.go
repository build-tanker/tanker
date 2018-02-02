package pings

import (
	"net/http"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/responses"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type PingHandler struct{}

func (p *PingHandler) Ping(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}
