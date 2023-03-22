package main

import (
	"context"
	"faker-generator/generator"
	"faker-generator/repository"
	"log"
)

func main() {
	ctx := context.Background()

	repo, err := repository.New(ctx)
	if err != nil {
		log.Println(err)
	}

	users, err := generator.GenerateUsers()
	if err != nil {
		log.Println(err)
	}

	for _, u := range users {
		orders, err := generator.GenerateOrders()
		if err != nil {
			log.Println(err)
		}

		err = repo.InsertUserWithOrders(ctx, u, orders)
		if err != nil {
			log.Println(err)
		}
	}

	repo.Close(ctx)
}
