package httpapi

import (
	"event-handler/accounts/internal/httpapi/handlers/profile"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func New() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	return e
}

var Module = fx.Module("httpapi", fx.Provide(
	profile.New,
))
