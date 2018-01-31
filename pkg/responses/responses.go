package responses

import (
	"encoding/json"
	"net/http"

	"github.com/sudhanshuraheja/tanker/pkg/model"
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

func NewErrorResponse(code string, message string) *Response {
	return &Response{
		Success: "false",
		Errors: []ErrorResponse{
			ErrorResponse{
				Code:    code,
				Message: message,
			},
		},
	}
}

func NewShipperAddSuccessResponse(id int64, accessKey string) *Response {
	return &Response{
		Data: &model.Shipper{
			ID:        id,
			AccessKey: accessKey,
		},
		Success: "true",
	}
}

func NewShipperViewAllSuccessResponse(shippers []model.Shipper) *Response {
	return &Response{
		Data:    shippers,
		Success: "true",
	}
}

func NewShipperViewSuccessResponse(shipper model.Shipper) *Response {
	return &Response{
		Data:    shipper,
		Success: "true",
	}
}

func NewShipperDeleteSuccessResponse() *Response {
	return &Response{
		Success: "true",
	}
}
