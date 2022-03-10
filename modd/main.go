package main

import (
	"context"
	"log"
	"watch-files/engine"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			func() string {
				return "Hello there!"
			},
			engine.New,
		),

		fx.Invoke(
			func(lc fx.Lifecycle, message string, e *engine.Engine) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Println(message)
						log.Println(e.RenderIndex())
						return nil
					},
					OnStop: func(ctx context.Context) error {
						log.Println("Shutting down")
						return nil
					},
				})
			},
		),
	)

	app.Run()
}
