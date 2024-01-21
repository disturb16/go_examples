package profiles

import "database/sql"

type Repository interface{}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
