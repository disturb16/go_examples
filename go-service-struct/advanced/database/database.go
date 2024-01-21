package database

import (
	"context"
	"database/sql"
)

func New(ctx context.Context) *sql.DB {
	return &sql.DB{}
}
