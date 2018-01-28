package pings

import "net/http"

type PingHandler struct{}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":\"pong\"}"))
}
