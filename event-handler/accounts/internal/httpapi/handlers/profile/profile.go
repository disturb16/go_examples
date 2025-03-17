package profile

import (
	"event-handler/accounts/internal/contracts"
	"event-handler/accounts/internal/services/profiler"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	profileManager profiler.Manager
}

type Result struct {
	fx.Out

	Handler contracts.HttpHandler `group:"handlers"`
}

type Response struct {
	Data any `json:"data"`
}

func New(pm profiler.Manager) Result {
	return Result{
		Handler: &handler{
			profileManager: pm,
		},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/profiles/:accountID", h.ProfileByAccountID)
	e.POST("/profiles", h.CreateProfile)
}
