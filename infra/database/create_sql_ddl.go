package database

import (
	"database/sql"
	"log"
)

func CreateSqlDDL(db *sql.DB) {
	ddls := []string{
		`CREATE TABLE IF NOT EXISTS visits (id INTEGER PRIMARY KEY AUTOINCREMENT, domain VARCHAR(255) NOT NULL, ip VARCHAR(255) NOT NULL, visit_date DATETIME not null)`,
		`CREATE INDEX IF NOT EXISTS idx_visits_domain ON visits (domain)`,
	}

	for _, ddl := range ddls {
		log.Printf("Executing DDL: %v", ddl)
		_, err := db.Exec(ddl)
		if err != nil {
			log.Printf("Error to exec the DDL: %v -> %v", ddl, err.Error())
			panic(err)
		}
	}
}
