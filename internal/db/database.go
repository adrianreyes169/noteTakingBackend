package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnection() *sql.DB {

	db, err := sql.Open("sqlite3", "notes.db")

	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}

	err2 := db.Ping()

	if err2 != nil {
		log.Fatal("Cannot connect to database: ", err2)
	}

	sqlBytes, err3 := os.ReadFile("internal/db/migrations.sql")
	if err3 != nil {
		log.Fatal("Cannot read migrations.sql: ", err3)
	}

	_, err4 := db.Exec(string(sqlBytes))
	if err4 != nil {
		log.Fatal("Could not run migrations: ", err4)
	}

	return db

}
