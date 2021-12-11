package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/disturb16/go-examples/dependency-injection/configuration"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSqliteConnection(ctx context.Context, config *configuration.Configuration) (*sql.DB, error) {
	log.Println("Connecting to database...")

	source := fmt.Sprintf("./%s.db", config.DB.Name)
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
