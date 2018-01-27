package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseConfig : configuration for the db
type DatabaseConfig struct {
	name        string
	host        string
	user        string
	password    string
	port        int
	maxPoolSize int
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		name:        viper.GetString("database.name"),
		host:        viper.GetString("database.host"),
		user:        viper.GetString("database.user"),
		password:    viper.GetString("database.password"),
		port:        viper.GetInt("database.port"),
		maxPoolSize: viper.GetInt("database.maxPoolSize"),
	}
}

// ConnectionString : string to connect to the db
func (db DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		db.name,
		db.user,
		db.password,
		db.host,
		db.port,
	)
}

// ConnectionURL : url to connect to db
func (db DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.user,
		db.password,
		db.host,
		db.port,
		db.name,
	)
}

// MaxPoolSize : export the maximum connection pool size
func (db DatabaseConfig) MaxPoolSize() int {
	return db.maxPoolSize
}
