package pings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/config"
	"source.golabs.io/core/tanker/pkg/logger"
)

var pingHandlerTestContext *appcontext.AppContext

func TestPingHandler(t *testing.T) {
	ctx := NewPingHandlerTestContext()
	pingHandler := PingHandler{}

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler.Ping(ctx))

	handler.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"success\":\"pong\"}\n", response.Body.String())
}

func NewPingHandlerTestContext() *appcontext.AppContext {
	if pingHandlerTestContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		pingHandlerTestContext = appcontext.NewAppContext(conf, log)
	}
	return pingHandlerTestContext
}
