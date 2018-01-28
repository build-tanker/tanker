package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	name                   string
	version                string
	logLevel               string
	port                   string
	enableStaticFileServer bool
	enableGzipCompression  bool
	enableDelayMiddleware  bool
	database               DatabaseConfig
}

func NewConfig() *Config {
	config := &Config{}

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigName("application")
	viper.SetConfigType("toml")

	viper.SetDefault("application.name", "tanker")
	viper.SetDefault("application.version", "NotDefined")
	viper.SetDefault("application.logLevel", "debug")
	viper.SetDefault("server.port", "4000")
	viper.SetDefault("server.enableStaticFileServer", false)
	viper.SetDefault("server.enableGzipCompression", true)
	viper.SetDefault("server.enableDelayMiddleware", false)

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

func (c *Config) Name() string {
	return c.name
}

func (c *Config) Version() string {
	return c.version
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func (c *Config) Port() string {
	return c.port
}

func (c *Config) EnableStaticFileServer() bool {
	return c.enableStaticFileServer
}

func (c *Config) EnableGzipCompression() bool {
	return c.enableGzipCompression
}

func (c *Config) EnableDelayMiddleware() bool {
	return c.enableDelayMiddleware
}

func (c *Config) Database() DatabaseConfig {
	return c.database
}

func (c *Config) readLatestConfig() {
	c.name = viper.GetString("application.name")
	c.version = viper.GetString("application.version")
	c.logLevel = viper.GetString("application.logLevel")
	c.port = viper.GetString("server.port")
	c.enableStaticFileServer = viper.GetBool("server.enableStaticFileServer")
	c.enableGzipCompression = viper.GetBool("server.enableGzipCompression")
	c.enableDelayMiddleware = viper.GetBool("server.enableDelayMiddleware")
	c.database = NewDatabaseConfig()
}
