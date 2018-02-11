package appcontext

import (
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

// AppContext - global context for config and logging
type AppContext struct {
	config *config.Config
	logger logger.Logger
}

// NewAppContext - function to create a global context for conf and logging
func NewAppContext(config *config.Config, logger logger.Logger) *AppContext {
	return &AppContext{
		config: config,
		logger: logger,
	}
}

// GetConfig - fetch the config from the global AppContext
func (a *AppContext) GetConfig() *config.Config {
	return a.config
}

// GetLogger - fetch the logger from the global AppContext
func (a *AppContext) GetLogger() logger.Logger {
	return a.logger
}
