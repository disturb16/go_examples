package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const schema string = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		order_date datetime NOT NULL,
		payment_method TEXT NOT NULL,
		payment_reference TEXT NOT NULL,
		user_id INTEGER NOT NULL,

		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

type Repository struct {
	db *sqlx.DB
}

func New(ctx context.Context) (Repository, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", "./data.db")
	if err != nil {
		return Repository{}, err
	}

	db.SetMaxOpenConns(1)

	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		return Repository{}, err
	}

	return Repository{db: db}, nil
}

func (r *Repository) Close(ctx context.Context) error {
	return r.db.Close()
}
