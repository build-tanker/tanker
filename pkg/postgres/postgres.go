package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

const (
	connMaxLifetime = 30 * time.Minute
	defaultTimeout  = 1 * time.Second
)

var database *sqlx.DB

// Init : Initialiase the database connection
func Init() {
	var err error

	database, err = sqlx.Open("postgres", config.Database().ConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	} else {
		logger.Debug("Connected to database")
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("Ping to the database failed: %s on connString %s", err, config.Database().ConnectionString())
	}

	database.SetMaxIdleConns(config.Database().MaxPoolSize())
	database.SetMaxOpenConns(config.Database().MaxPoolSize())
	database.SetConnMaxLifetime(connMaxLifetime)
}

// Close : close the db connection
func Close() error {
	logger.Debug("Closing the DB connection")
	return database.Close()
}

// Get : get a reference to the database connection
func Get() *sqlx.DB {
	return database
}
