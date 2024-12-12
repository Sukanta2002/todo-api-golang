package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./todo.db")

	if err != nil {
		log.Fatalln("Database connection erroe", err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT 0
	);`)

	return db
}
