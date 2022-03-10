package main

import (
	"context"
	"gotests/api"
	"gotests/database"
	"gotests/internal/books"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			database.LoadConfig,
			database.NewPostgresDB,
			books.NewRepository,
			books.NewService,
			echo.New,
			api.New,
		),

		fx.Invoke(
			api.RegisterRoutes,
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				e.Logger.Fatal(e.Start(":8081"))
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Close()
		},
	})
}
