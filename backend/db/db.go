package db

import (
  "database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Ares1605/casual-chess-backend/env"
)

func Conn() (*sql.DB, error) {
	dbPath := env.Get("DB_PATH")
	if dbPath == "" {
	  return nil, errors.New("DB_PATH environment variable must be specified")
	}
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
}


