package database

import (
	"database/sql"
	"event-handler/payments/config"
	"fmt"

	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func New(c *config.Config) (*sql.DB, error) {
	dbPath := fmt.Sprintf("file:%s", c.DBPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PopulateDb(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		return err
	}

	return nil
}
