package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "username"
	password = "username666"
	dbname   = "lollipop_db"
)

type Database struct {
	db *sql.DB
}

func New() (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain database connection: %s", err.Error())
	}

	return &Database{
		db: db,
	}, nil
}
