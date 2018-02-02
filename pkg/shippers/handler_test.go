package shippers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"

	"source.golabs.io/core/tanker/pkg/appcontext"
)

type MockService struct {
	ctx *appcontext.AppContext
}

func NewMockService(ctx *appcontext.AppContext) *MockService {
	return &MockService{
		ctx: ctx,
	}
}

func (m *MockService) Add(name string, machineName string) (int64, string, error) {
	return 15, "testAccessKey", nil
}

func (m *MockService) Delete(id int64) error {
	return nil
}

func (m *MockService) View(id int64) (Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return Shipper{
		ID:          15,
		AccessKey:   "testAccessKey",
		Name:        "testName",
		MachineName: "testMachineName",
		CreatedAt:   testTime,
		UpdatedAt:   testTime,
	}, nil
}

func (m *MockService) ViewAll() ([]Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return []Shipper{
		Shipper{
			ID:          15,
			AccessKey:   "testAccessKey",
			Name:        "testName",
			MachineName: "testMachineName",
			CreatedAt:   testTime,
			UpdatedAt:   testTime,
		},
	}, nil
}

func NewTestHandler() *handler {
	ctx := NewTestContext()
	return &handler{
		service: NewMockService(ctx),
	}
}

func TestHandlerAdd(t *testing.T) {
	h := NewTestHandler()

	req, err := http.NewRequest(http.MethodPost, "/v1/shippers", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers", h.Add()).Methods(http.MethodPost)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":{\"id\":15,\"access_key\":\"testAccessKey\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"},\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerView(t *testing.T) {
	h := NewTestHandler()

	req, err := http.NewRequest(http.MethodGet, "/v1/shippers/15", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers/{id}", h.View()).Methods(http.MethodGet)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":{\"id\":15,\"access_key\":\"testAccessKey\",\"name\":\"testName\",\"machine_name\":\"testMachineName\",\"created_at\":\"2018-01-31T01:01:01.000000001Z\",\"updated_at\":\"2018-01-31T01:01:01.000000001Z\"},\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerViewAll(t *testing.T) {
	h := NewTestHandler()

	req, err := http.NewRequest(http.MethodGet, "/v1/shippers", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/shippers", h.ViewAll()).Methods(http.MethodGet)
	router.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"data\":[{\"id\":15,\"access_key\":\"testAccessKey\",\"name\":\"testName\",\"machine_name\":\"testMachineName\",\"created_at\":\"2018-01-31T01:01:01.000000001Z\",\"updated_at\":\"2018-01-31T01:01:01.000000001Z\"}],\"success\":\"true\"}\n", response.Body.String())
}

func TestHandlerDelete(t *testing.T) {
	h := NewTestHandler()

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
