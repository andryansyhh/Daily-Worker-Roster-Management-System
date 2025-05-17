package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewClientDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "file:roster.db?cache=shared&_foreign_keys=on")
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	if err := runMigration(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ DB connected & migrated")
	return db
}

func runMigration(db *sql.DB) error {
	content, err := os.ReadFile("./migration/migration.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(content))
	return err
}
