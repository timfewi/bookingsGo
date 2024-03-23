package driver

import (
	"database/sql"
	"fmt"
	"time"

	// Import the pqx driver for PostgreSQL
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/timfewi/bookingsGo/internal/helpers"
)

// DB is the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL connects to the database
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		fmt.Println("Error connecting to database")
		helpers.ServerError(nil, err)
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxIdleConns)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		fmt.Println("Error connecting to database")
		helpers.ServerError(nil, err)
		return nil, err
	}

	return dbConn, nil
}

// NewDatabase creates a new database connection
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println("Error opening database connection")
		helpers.ServerError(nil, err)
		return nil, err

	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error pinging database")
		helpers.ServerError(nil, err)
		return nil, err
	}

	return db, nil
}

// testDB tries to ping the database connection
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error pinging database")
		helpers.ServerError(nil, err)
		return err
	}

	return nil
}
