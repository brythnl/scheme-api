package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func Connect(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, errors.New("DSN not set")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
