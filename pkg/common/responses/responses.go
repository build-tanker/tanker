package responses

import (
	"encoding/json"
	"net/http"
)

// WriteJSON encodes the JSON
func WriteJSON(w http.ResponseWriter, status int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)
}

// Response defines the standard output response
type Response struct {
	Data    interface{}     `json:"data,omitempty"`
	Success string          `json:"success,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

// ErrorResponse defines the standard error response
type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// NewErrorResponse writes a new error to the response
func NewErrorResponse(code string, message string) *Response {
	return &Response{
		Success: "false",
		Errors: []ErrorResponse{
			{
				Code:    code,
				Message: message,
			},
		},
	}
}
