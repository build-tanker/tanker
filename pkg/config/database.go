package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	name        string
	host        string
	user        string
	password    string
	port        int
	maxPoolSize int
}

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

func (db DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		db.name,
		db.user,
		db.password,
		db.host,
		db.port,
	)
}

func (db DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.user,
		db.password,
		db.host,
		db.port,
		db.name,
	)
}

func (db DatabaseConfig) MaxPoolSize() int {
	return db.maxPoolSize
}
