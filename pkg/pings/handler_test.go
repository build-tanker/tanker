package pings

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
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
		conf := config.NewConfig([]string{".", "..", "../.."})
		log := logger.NewLogger(conf, os.Stdout)
		pingHandlerTestContext = appcontext.NewAppContext(conf, log)
	}
	return pingHandlerTestContext
}
