package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

const (
	connMaxLifetime = 30 * time.Minute
	defaultTimeout  = 1 * time.Second
)

var schema = `
	CREATE TABLE branch (
		bank text
		ifsc text
		micr text
		branch text
		address text
		city text
		district text
		state text
		contact text
	)
`

// Branch : struct for the data in branch table
type Branch struct {
	Bank     string `db:"bank"`
	Ifsc     string `db:"ifsc"`
	Micr     string `db:"micr"`
	Branch   string `db:"branch"`
	Address  string `db:"address"`
	City     string `db:"city"`
	District string `db:"district"`
	State    string `db:"state"`
	Contact  string `db:"contact"`
}

var db *sqlx.DB

// Init : Initialiase the database connection
func Init() {
	var err error

	db, err := sqlx.Open("postgres", config.Database().ConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	} else {
		logger.Debug("Connected to database")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping to the database failed: %s on connString %s", err, config.Database().ConnectionString())
	}

	db.SetMaxIdleConns(config.Database().MaxPoolSize())
	db.SetMaxOpenConns(config.Database().MaxPoolSize())
	db.SetConnMaxLifetime(connMaxLifetime)

	// db.MustExec(schema)

}

// Close : close the db connection
func Close() error {
	return db.Close()
}

// Get : get a reference to the database connection
func Get() *sqlx.DB {
	return db
}
