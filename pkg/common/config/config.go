package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for tanker
type Config struct {
	server struct {
		port      string
		host      string
		fileStore string
	}
	gcs struct {
		JSONConfig string
		bucket     string
	}
	database struct {
		name        string
		host        string
		user        string
		password    string
		port        int
		maxPoolSize int
	}
}

// New creates a new configuration
func New(paths []string) *Config {
	config := &Config{}

	viper.AutomaticEnv()

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("tanker")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.server.port
}

// Host - get the host from config
func (c *Config) Host() string {
	return c.server.host
}

// ConnectionString - get the connectionstring to connect to postgres
func (c *Config) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		c.database.name,
		c.database.user,
		c.database.password,
		c.database.host,
		c.database.port,
	)
}

// ConnectionURL - get the connection URL to connect to postgres
func (c *Config) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.database.user,
		c.database.password,
		c.database.host,
		c.database.port,
		c.database.name,
	)
}

// MaxPoolSize - get the max pool size for db connections
func (c *Config) MaxPoolSize() int {
	return c.database.maxPoolSize
}

// FileStore - get the filestore being used in the app
func (c *Config) FileStore() string {
	return c.server.fileStore
}

// GCSJSONConfig - get the google cloud json config path
func (c *Config) GCSJSONConfig() string {
	return c.gcs.JSONConfig
}

// GCSBucket - get the google cloud bucket
func (c *Config) GCSBucket() string {
	return c.gcs.bucket
}

func (c *Config) mustGetString(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("The key %s was not found, crashing\n", key)
	}
	return value
}

func (c *Config) mustGetInt(key string) int {
	value := c.mustGetString(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("They key %s was not an integer (%s), crashing\n", key, value)
	}
	return intValue
}

func (c *Config) getString(key string, defaultValue string) string {
	value := viper.GetString(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (c *Config) getInt(key string, defaultValue int) int {
	value := viper.GetInt(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func (c *Config) readLatestConfig() {
	c.server.host = c.getString("SERVER_HOST", "http://localhost")
	c.server.port = c.getString("SERVER_PORT", "4000")
	c.server.fileStore = c.getString("SERVER_FILESTORE", "googlecloud")

	c.gcs.JSONConfig = c.mustGetString("GCS_JSON_CONFIG")
	c.gcs.bucket = c.mustGetString("GCS_BUCKET")

	c.database.host = c.mustGetString("DB_HOST")
	c.database.port = c.mustGetInt("DB_PORT")
	c.database.name = c.mustGetString("DB_NAME")
	c.database.user = c.mustGetString("DB_USER")
	c.database.password = c.mustGetString("DB_PASSWORD")
	c.database.maxPoolSize = c.getInt("DB_MAX_POOL_SIZE", 5)
}
