package profiles

import (
	"github.com/disturb16/go-service-struct/advanced/internal/httpapi/handlers"
	"github.com/disturb16/go-service-struct/advanced/internal/services/profiles"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	service profiles.Service
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
	group := e.Group("/profiles")
	group.GET("", h.Get)
	group.PUT("", h.Update)

}
