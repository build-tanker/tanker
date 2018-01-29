package responses

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)
}

type Response struct {
	Data    interface{}     `json:"data,omitempty"`
	Success string          `json:"success,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
