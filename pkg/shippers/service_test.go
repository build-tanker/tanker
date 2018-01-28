package shippers

import (
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

var shipperServiceTestContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if shipperServiceTestContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		shipperServiceTestContext = appcontext.NewAppContext(conf, log)
	}
	return shipperDatastoreTestContext
}
