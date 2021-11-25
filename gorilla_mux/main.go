package main

import (
	"context"
	"fmt"

	"github.com/disturb16/go_example/gorillamux/api"
	"github.com/disturb16/go_example/gorillamux/dbconnection"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	var number int
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db := dbconnection.New()
	err = db.QueryRowContext(ctx, "SELECT 1").Scan(&number)
	if err != nil {
		panic(err)
	}

	fmt.Println(number)

	srv := api.New()
	err = srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
