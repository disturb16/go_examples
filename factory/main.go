package main

import (
	"fmt"

	"github.com/user/factory-pattern/configuration"
	"github.com/user/factory-pattern/repository"
)

// Factory Pattern

func main() {
	config := &configuration.Configuration{
		Engine: "sqlserver",
		Host:   "localhost",
	}

	repo, err := repository.New(config)
	if err != nil {
		panic(err)
	}

	err = repo.Save("")
	if err != nil {
		panic(err)
	}

	data := repo.Find(1)
	fmt.Println(data)
}
