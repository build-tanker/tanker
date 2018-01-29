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

type Shipper struct {
	ID          int64  `json:"id,omitempty"`
	AccessKey   string `json:"access_key,omitempty"`
	Name        string `json:"name,omitempty"`
	MachineName string `json:"machine_name,omitempty"`
	CreatedAt   int    `json:"created_at,omitempty"`
	UpdatedAt   int    `json:"updated_at,omitempty"`
}

func NewShipperAddSuccessResponse(id int64, accessKey string) *Response {
	return &Response{
		Data: &Shipper{
			ID:        id,
			AccessKey: accessKey,
		},
	}
}
