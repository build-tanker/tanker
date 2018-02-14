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
	return "https://storage.googleapis.com/testBucket/206329dc-a2af-42a4-9977-13990d0c25dc?Expires=1518408172&GoogleAccessId=tanker-gcs-upload-test%40tanker-194004&Signature=OuaiGWj%2BHxJp0LOlu67SwLz1pbwLHrlNlSugrqgLD%2Fv6", nil

}

func NewTestHandler() *handler {
	s := NewMockService()
	ctx := NewTestContext()
	return &handler{
		ctx:     ctx,
		service: s,
	}
}

func TestHandlerAdd(t *testing.T) {
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
	assert.Equal(t, "https://storage.googleapis.com/testBucket/206329dc-a2af-42a4-9977-13990d0c25dc?Expires=1518408172&GoogleAccessId=tanker-gcs-upload-test%40tanker-194004&Signature=OuaiGWj%2BHxJp0LOlu67SwLz1pbwLHrlNlSugrqgLD%2Fv6", r["data"].(map[string]interface{})["url"].(string))
}
