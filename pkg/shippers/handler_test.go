package shippers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/tanker/pkg/common/config"
)

type MockService struct {
	conf *config.Config
}

func NewMockService(conf *config.Config) *MockService {
	return &MockService{
		conf: conf,
	}
}

func (m *MockService) Add(appGroup string, expiry int) (string, error) {
	return "testId", nil
}

func (m *MockService) Delete(id string) error {
	return nil
}

func (m *MockService) View(id string) (Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return Shipper{
		ID:        "testId",
		AppGroup:  "testAppGroup",
		Expiry:    10,
		Deleted:   false,
		CreatedAt: testTime,
		UpdatedAt: testTime,
	}, nil
}

func (m *MockService) ViewAll() ([]Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return []Shipper{
		Shipper{
			ID:        "testId",
			AppGroup:  "testAppGroup",
			Expiry:    10,
			Deleted:   false,
			CreatedAt: testTime,
			UpdatedAt: testTime,
		},
	}, nil
}

func newTestHandler() *handler {
	conf := config.New([]string{".", "..", "../.."})
	return &handler{
		service: NewMockService(conf),
	}
}

func TestHandlerAdd(t *testing.T) {
	h := newTestHandler()

	req, err := http.NewRequest(http.MethodPost, "/v1/shippers", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers", h.Add()).Methods(http.MethodPost)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":{\"id\":\"testId\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"},\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerView(t *testing.T) {
	h := newTestHandler()

	req, err := http.NewRequest(http.MethodGet, "/v1/shippers/15", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers/{id}", h.View()).Methods(http.MethodGet)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":{\"id\":\"testId\",\"app_group\":\"testAppGroup\",\"expiry\":10,\"created_at\":\"2018-01-31T01:01:01.000000001Z\",\"updated_at\":\"2018-01-31T01:01:01.000000001Z\"},\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerViewAll(t *testing.T) {
	h := newTestHandler()

	req, err := http.NewRequest(http.MethodGet, "/v1/shippers", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers", h.ViewAll()).Methods(http.MethodGet)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":[{\"id\":\"testId\",\"app_group\":\"testAppGroup\",\"expiry\":10,\"created_at\":\"2018-01-31T01:01:01.000000001Z\",\"updated_at\":\"2018-01-31T01:01:01.000000001Z\"}],\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerDelete(t *testing.T) {
	h := newTestHandler()

	req, err := http.NewRequest(http.MethodDelete, "/v1/shippers/15", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers/{id}", h.Delete()).Methods(http.MethodDelete)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"success\":\"true\"}\n", response.Body.String())
}
