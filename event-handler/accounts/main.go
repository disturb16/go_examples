package main

import (
	"context"
	"database/sql"
	"event-handler/accounts/config"
	"event-handler/accounts/database"
	"event-handler/accounts/internal/contracts"
	"event-handler/accounts/internal/httpapi"
	"event-handler/accounts/internal/listener"
	"event-handler/accounts/internal/repositories"
	"event-handler/accounts/internal/services"
	"event-handler/accounts/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc                   fx.Lifecycle
	Config               *config.Config
	DB                   *sql.DB
	Echo                 *echo.Echo
	Handlers             []contracts.HttpHandler `group:"handlers"`
	Logger               *logger.Logger
	EventHandlers        []contracts.EventHandler `group:"event_handlers"`
	EventsListenerRouter listener.EventsListenerRouter
}

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			logger.New,
			database.New,
			httpapi.New,
			listener.New,
		),
		repositories.Module,
		services.Module,
		httpapi.Module,
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
			p.Logger.Info("starting server")
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

			for _, h := range p.EventHandlers {
				p.EventsListenerRouter.RegisterHandler(h)
			}

			go func() {
				p.EventsListenerRouter.Start(context.Background())
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			if err := p.Echo.Shutdown(ctx); err != nil {
				p.Logger.WithError(err).Error("failed to shutdown server")
			}

			if err := p.DB.Close(); err != nil {
				p.Logger.WithError(err).Error("failed to close db connection")
			}

			if err := p.EventsListenerRouter.Stop(); err != nil {
				p.Logger.WithError(err).Error("failed to stop event listener router")
			}

			return nil
		},
	})
}
