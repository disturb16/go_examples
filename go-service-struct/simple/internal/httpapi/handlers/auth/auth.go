package auth

import (
	"github.com/disturb16/go-service-struct/simple/internal/httpapi/handlers"
	"github.com/disturb16/go-service-struct/simple/internal/services/auth"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	AuthService auth.Service
}

type Result struct {
	fx.Out

	Handler handlers.Handler `group:"handlers"`
}

func New() Result {
	return Result{
		Handler: &handler{},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	group := e.Group("/auth")
	group.POST("", h.Authenticate)
	group.POST("/token/refresh", h.RefreshToken)
}
