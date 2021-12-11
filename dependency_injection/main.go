package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/disturb16/go-examples/dependency-injection/configuration"
	"github.com/disturb16/go-examples/dependency-injection/database"
	"github.com/disturb16/go-examples/dependency-injection/repository"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			configuration.New,
			database.CreateSqliteConnection,
			repository.New,
		),
		fx.Invoke(
			configureLifyeCycleHooks,
		),
	)

	app.Run()
}

func configureLifyeCycleHooks(lc fx.Lifecycle, db *sql.DB, repo *repository.Repository) {
	lc.Append(
		fx.Hook{
			OnStart: func(c context.Context) error {
				fmt.Println("Starting application")

				err := repo.SaveUser(c, "Charles", "Doe")
				if err != nil {
					return err
				}

				uu, err := repo.GetUsers(c)
				if err != nil {
					return err
				}

				fmt.Println(uu)

				return nil
			},

			OnStop: func(c context.Context) error {
				fmt.Println("Closing database connection")
				return db.Close()
			},
		},
	)
}
