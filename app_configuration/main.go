package main

import (
	"context"
	"fmt"

	"github.com/disturb16/go-examples/app-configuration/configuration"
	"github.com/disturb16/go-examples/app-configuration/database"
)

func main() {
	var id int
	var err error
	ctx := context.Background()

	config, err := configuration.Load("config.yml")
	if err != nil {
		panic(err)
	}

	// create connection
	db := database.CreateConnection(ctx, config)
	defer db.Close()

	// test query
	err = db.QueryRowContext(ctx, "SELECT id from books").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	fmt.Println("All good")
}
