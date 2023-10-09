// internal/storage/dbInit.go

package storage

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func DBInit() (*sql.DB, error) {
	log.Println("database is starting up...")

	dbPath := filepath.Join(".", "database")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err := os.Mkdir(dbPath, 0755); err != nil {
			return nil, err
		}
	}

	dbFile := filepath.Join(dbPath, "onion-gate.db")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	if err := createUsersTable(db); err != nil {
		return err
	}
	return nil
}

func createUsersTable(db *sql.DB) error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		date_created DATETIME DEFAULT CURRENT_TIMESTAMP,
		last_login DATETIME,
		is_active BOOLEAN DEFAULT TRUE
	);`

	_, err := db.Exec(usersTable)
	if err != nil {
		return err
	}

	return nil
}
