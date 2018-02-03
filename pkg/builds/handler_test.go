package builds

import (
	"encoding/json"
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

func (m MockService) Add(accessKey string, bundle string) (string, error) {
	return "https://mockBucket.storage.googleapis.com/mockFile", nil
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

	req, err := http.NewRequest(http.MethodPost, "/v1/builds?accessKey=a1b2c3&bundle=com.me.app", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/builds", h.Add()).Methods(http.MethodPost)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)

	var r map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &r)
	assert.Equal(t, "true", r["success"].(string))
	assert.Equal(t, "https://mockBucket.storage.googleapis.com/mockFile", r["data"].(map[string]interface{})["url"].(string))
}
