package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

const schemaQuery string = `
create table if not exists users
(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName TEXT,
    LastName TEXT
)`

func New(db *sql.DB) (*Repository, error) {

	// Populate schema
	_, err := db.Exec(schemaQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}
