package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseConfig - database configuration settings for postgres
type DatabaseConfig struct {
	name        string
	host        string
	user        string
	password    string
	port        int
	maxPoolSize int
}

// NewDatabaseConfig - initialise settings for a new DB
func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		name:        viper.GetString("database.name"),
		host:        viper.GetString("database.host"),
		user:        viper.GetString("database.user"),
		password:    viper.GetString("database.password"),
		port:        viper.GetInt("database.port"),
		maxPoolSize: viper.GetInt("database.maxPoolSize"),
	}
}

// ConnectionString - get the connectionstring to connect to postgres
func (db DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		db.name,
		db.user,
		db.password,
		db.host,
		db.port,
	)
}

// ConnectionURL - get the connection URL to connect to postgres
func (db DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.user,
		db.password,
		db.host,
		db.port,
		db.name,
	)
}

// MaxPoolSize - get the max pool size for db connections
func (db DatabaseConfig) MaxPoolSize() int {
	return db.maxPoolSize
}
