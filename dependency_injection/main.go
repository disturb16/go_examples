package main

import (
	"context"

	"github.com/disturb16/go-examples/dependency-injection/configuration"
	"github.com/disturb16/go-examples/dependency-injection/database"
)

func main() {
	var err error
	ctx := context.Background()

	config, err := configuration.Load("config.yml")
	if err != nil {
		panic(err)
	}

	// create connection
	db, err := database.CreateSqliteConnection(ctx, config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
