package filestore

import (
	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/config"
	"source.golabs.io/core/tanker/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}
