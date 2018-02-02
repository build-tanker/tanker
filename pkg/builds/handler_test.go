package builds

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func NewMockService() Service {
	return MockService{}
}

func NewTestHandler() *handler {
	s := NewMockService()
	ctx := NewTestContext()
	return &handler{
		ctx:     ctx,
		service: s,
	}
}

func TestAdd(t *testing.T) {
	h := NewTestHandler()

	req, err := http.NewRequest(http.MethodPost, "/v1/builds?accessKey=a1b2c3", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/builds", h.Add()).Methods(http.MethodPost)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "", response.Body.String())
}
