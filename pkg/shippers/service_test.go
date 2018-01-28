package shippers

import (
	"testing"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

func NewTestContext() *appcontext.AppContext {
	conf := config.NewConfig()
	log := logger.NewLogger(conf)
	return appcontext.NewAppContext(conf, log)
}

func TestShippersServiceAdd(t *testing.T) {
	// ctx := NewTestContext()
	// shippersService := NewShippersService(ctx)

	// _, err := shippersService.Add("test", "machine.test")
	// assert.Nil(t, err)
	// assert.Equal(t, "test", ship.Name)
	// assert.Equal(t, "machine.test", ship.MachineName)
}
