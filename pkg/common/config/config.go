package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for tanker
type Config struct {
	server struct {
		port string
		host string
	}
	gcs struct {
		JSONConfig  string
		bucket      string
		credentials googleCredentials
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

type googleCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
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

// GCSJSONConfig - get the google cloud json config path
func (c *Config) GCSJSONConfig() string {
	return c.gcs.JSONConfig
}

// GCSBucket - get the google cloud bucket
func (c *Config) GCSBucket() string {
	return c.gcs.bucket
}

// GCSCredentials - get the google credentials
func (c *Config) GCSCredentials() googleCredentials {
	return c.gcs.credentials
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

	c.gcs.JSONConfig = c.mustGetString("GCS_JSON_CONFIG")
	c.gcs.bucket = c.mustGetString("GCS_BUCKET")
	c.readGoogleConfig()

	c.database.host = c.mustGetString("DB_HOST")
	c.database.port = c.mustGetInt("DB_PORT")
	c.database.name = c.mustGetString("DB_NAME")
	c.database.user = c.mustGetString("DB_USER")
	c.database.password = c.mustGetString("DB_PASSWORD")
	c.database.maxPoolSize = c.getInt("DB_MAX_POOL_SIZE", 5)
}

func (c *Config) readGoogleConfig() {
	data, err := ioutil.ReadFile(c.gcs.JSONConfig)
	if err != nil {
		log.Fatalf("Could not read the google config file at %s due to %s", c.gcs.JSONConfig, err.Error())
	}
	err = json.Unmarshal(data, &c.gcs.credentials)
	if err != nil {
		log.Fatalf("Could not unmarshal gcs json, %s", err.Error())
	}
}
