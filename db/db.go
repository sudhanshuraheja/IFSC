package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/sudhanshuraheja/IFSC/config"
)

const (
	connMaxLifetime = 30 * time.Minute
	defaultTimeout  = 1 * time.Second
)

var db *sqlx.DB

// Init : Initialiase the database connection
func Init() {
	var err error

	db, err := sqlx.Open("postgres", config.Database().ConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping to the database failed: %s on connString %s", err, config.Database().ConnectionString())
	}

	db.SetMaxIdleConns(config.Database().MaxPoolSize())
	db.SetMaxOpenConns(config.Database().MaxPoolSize())
	db.SetConnMaxLifetime(connMaxLifetime)
}

// Close : close the db connection
func Close() error {
	return db.Close()
}

// Get : get a reference to the database connection
func Get() *sqlx.DB {
	return db
}
