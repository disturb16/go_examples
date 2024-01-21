package main

import (
	"context"
	"database/sql"
	"log"
	"msgbroker"

	"github.com/disturb16/go-service-struct/advanced/config"
	"github.com/disturb16/go-service-struct/advanced/database"
	"github.com/disturb16/go-service-struct/advanced/internal/listener"
	"github.com/disturb16/go-service-struct/advanced/internal/repositories"
	"github.com/disturb16/go-service-struct/advanced/internal/services"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc        fx.Lifecycle
	Config    *config.Config
	DB        *sql.DB
	Broker    msgbroker.Broker
	Procesors []msgbroker.Processor `group:"procesors"`
}

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			database.New,
			msgbroker.New,
		),
		repositories.Module,
		services.Module,
		listener.Module,
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(p Params) {
	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			for _, procesor := range p.Procesors {
				procesor.Register(p.Broker)
			}

			go func() {
				p.Broker.Start()
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			if err := p.DB.Close(); err != nil {
				log.Println(err)
			}

			p.Broker.Stop()

			return nil
		},
	})
}
