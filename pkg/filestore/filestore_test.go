package filestore

import (
	"os"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig([]string{".", "..", "../.."})
		log := logger.NewLogger(conf, os.Stdout)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}
