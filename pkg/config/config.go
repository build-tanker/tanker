package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for tanker
type Config struct {
	name                   string
	version                string
	logLevel               string
	fileStore              string
	port                   string
	enableStaticFileServer bool
	enableGzipCompression  bool
	enableDelayMiddleware  bool
	gcsJSONConfig          string
	gcsBucket              string
	database               DatabaseConfig
}

// NewConfig - create a new configuration
func NewConfig(paths []string) *Config {
	config := &Config{}

	viper.AutomaticEnv()

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("application")
	viper.SetConfigType("toml")

	viper.SetDefault("application.name", "tanker")
	viper.SetDefault("application.version", "NotDefined")
	viper.SetDefault("application.logLevel", "debug")
	viper.SetDefault("application.fileStore", "googlecloud")
	viper.SetDefault("server.port", "4000")
	viper.SetDefault("server.enableStaticFileServer", false)
	viper.SetDefault("server.enableGzipCompression", true)
	viper.SetDefault("server.enableDelayMiddleware", false)
	viper.SetDefault("googlecloud.jsonConfig", "")
	viper.SetDefault("googlecloud.bucket", "shrieking-cat")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// Name - get the app name from config
func (c *Config) Name() string {
	return c.name
}

// Version - get the app version from config
func (c *Config) Version() string {
	return c.version
}

// LogLevel - get the log level from config
func (c *Config) LogLevel() string {
	return c.logLevel
}

// FileStore - get the filestore from config
func (c *Config) FileStore() string {
	return c.fileStore
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.port
}

// EnableStaticFileServer - get if the static file server is enabled or not from the config
func (c *Config) EnableStaticFileServer() bool {
	return c.enableStaticFileServer
}

// EnableGzipCompression - get if gzip compression is enabled or not in the config
func (c *Config) EnableGzipCompression() bool {
	return c.enableGzipCompression
}

// EnableDelayMiddleware - get if delay middlware is enabled or not from the config
func (c *Config) EnableDelayMiddleware() bool {
	return c.enableDelayMiddleware
}

// Database - load the database config
func (c *Config) Database() DatabaseConfig {
	return c.database
}

// GCSJSONConfig - get the google cloud storage json path from config
func (c *Config) GCSJSONConfig() string {
	return c.gcsJSONConfig
}

// GCSBucket - get the google cloud storage bucket from config
func (c *Config) GCSBucket() string {
	return c.gcsBucket
}

func (c *Config) readLatestConfig() {
	c.name = viper.GetString("application.name")
	c.version = viper.GetString("application.version")
	c.logLevel = viper.GetString("application.logLevel")
	c.fileStore = viper.GetString("application.fileStore")
	c.port = viper.GetString("server.port")
	c.enableStaticFileServer = viper.GetBool("server.enableStaticFileServer")
	c.enableGzipCompression = viper.GetBool("server.enableGzipCompression")
	c.enableDelayMiddleware = viper.GetBool("server.enableDelayMiddleware")
	c.gcsJSONConfig = viper.GetString("googlecloud.jsonConfig")
	c.gcsBucket = viper.GetString("googlecloud.bucket")
	c.database = NewDatabaseConfig()
}
