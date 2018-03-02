package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for tanker
type Config struct {
	port     string
	database DatabaseConfig
}

// NewConfig - create a new configuration
func NewConfig(paths []string) *Config {
	config := &Config{}

	viper.AutomaticEnv()

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("tanker")
	viper.SetConfigType("toml")

	viper.SetDefault("server.port", "4000")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.port
}

// Database - load the database config
func (c *Config) Database() DatabaseConfig {
	return c.database
}

func (c *Config) readLatestConfig() {
	c.port = viper.GetString("server.port")
	c.database = NewDatabaseConfig()
}
