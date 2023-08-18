package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/data.sqlite3")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		panic(err)
	}
	return db
}
