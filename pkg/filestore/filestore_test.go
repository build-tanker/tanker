package filestore

import (
	"os"

	"github.com/gojekfarm/tanker/pkg/appcontext"
	"github.com/gojekfarm/tanker/pkg/config"
	"github.com/gojekfarm/tanker/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {

	if testContext == nil {
		conf := config.NewConfig([]string{".", "../config/testutil"})
		log := logger.NewLogger(conf, os.Stdout)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}
