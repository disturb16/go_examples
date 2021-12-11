package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/disturb16/go-examples/dependency-injection/configuration"
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

func New(ctx context.Context, config *configuration.Configuration, db *sql.DB) (*Repository, error) {

	// Populate schema
	_, err := db.ExecContext(ctx, schemaQuery)
	if err != nil {
		return nil, err
	}

	fmt.Println(config.DB.Name)

	return &Repository{db}, nil
}
