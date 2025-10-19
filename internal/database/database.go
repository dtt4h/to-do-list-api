package database

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Ensure the directory exists
	if err := os.MkdirAll("./internal/database", 0755); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}
	DB, err = sql.Open("sqlite", "./internal/database/tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// create DB
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending'
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB has been initialized")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
