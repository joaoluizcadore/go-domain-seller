package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.sqlite3")
	if err != nil {
		panic(err)
	}
	return db
}
