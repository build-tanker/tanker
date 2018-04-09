package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

const connMaxLifetime = 30 * time.Minute

// New initialize a new postgres connection
func New(url string, maxOpenConns int) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalln("Could not connect to database:", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Ping to the database failed:", err.Error(), "on connString", url)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	log.Println("Connected to database")

	return db
}
