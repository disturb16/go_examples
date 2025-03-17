package main

import (
	"context"
	"database/sql"
	"event-handler/payments/config"
	"event-handler/payments/database"
	"event-handler/payments/internal/contracts"
	"event-handler/payments/internal/httpapi"
	"event-handler/payments/internal/repositories"
	"event-handler/payments/internal/services"
	"event-handler/payments/logger"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc       fx.Lifecycle
	Config   *config.Config
	DB       *sql.DB
	Echo     *echo.Echo
	Handlers []contracts.HttpHandler `group:"handlers"`
}

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			database.New,
			httpapi.New,
			logger.New,
			amqpPublisher,
		),
		repositories.Module,
		services.Module,
		httpapi.Module,
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func amqpPublisher(c *config.Config, log *logger.Logger) (*amqp.Publisher, error) {
	amqpConfig := amqp.NewDurableQueueConfig(c.RabbitMQAddr)
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewSlogLogger(log.Logger))
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func setLifeCycle(p Params) {
	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			for _, h := range p.Handlers {
				h.RegisterRoutes(p.Echo)
			}

			err := database.PopulateDb(p.DB)
			if err != nil {
				return err
			}

			go func() {
				p.Echo.Logger.Fatal(p.Echo.Start(p.Config.Address))
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			if err := p.Echo.Shutdown(ctx); err != nil {
				log.Println(err)
			}

			if err := p.DB.Close(); err != nil {
				log.Println(err)
			}

			return nil
		},
	})
}
