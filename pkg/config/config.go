package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config object
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

var config *Config

// Init config from file
func Init() {
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
		readLatestConfig()
	})

	readLatestConfig()
}

func readLatestConfig() {
	config = &Config{
		name:     viper.GetString("application.name"),
		version:  viper.GetString("application.version"),
		logLevel: viper.GetString("application.logLevel"),
		port:     viper.GetString("server.port"),
		enableStaticFileServer: viper.GetBool("server.enableStaticFileServer"),
		enableGzipCompression:  viper.GetBool("server.enableGzipCompression"),
		enableDelayMiddleware:  viper.GetBool("server.enableDelayMiddleware"),
		database:               newDatabaseConfig(),
	}

}

// Name : Exporting Name
func Name() string {
	return config.name
}

// Version : Export application version
func Version() string {
	return config.version
}

// LogLevel : Export the log level
func LogLevel() string {
	return config.logLevel
}

// Port : Export the port
func Port() string {
	return config.port
}

// EnableStaticFileServer : Export if we are enabling the static file server
func EnableStaticFileServer() bool {
	return config.enableStaticFileServer
}

// EnableGzipCompression : Export if we want to enable Gzip compression
func EnableGzipCompression() bool {
	return config.enableGzipCompression
}

// EnableDelayMiddleware : Export if we want to enable delay middleware
func EnableDelayMiddleware() bool {
	return config.enableDelayMiddleware
}

// Database : export the database configuration
func Database() DatabaseConfig {
	return config.database
}
