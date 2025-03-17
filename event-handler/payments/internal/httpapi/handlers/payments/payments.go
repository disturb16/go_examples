package payments

import (
	"event-handler/payments/internal/contracts"
	"event-handler/payments/internal/services/paymentsmanager"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	paymentManager paymentsmanager.Manager
}

type Result struct {
	fx.Out

	Handler contracts.HttpHandler `group:"handlers"`
}

func New(pm paymentsmanager.Manager) Result {
	return Result{
		Handler: &handler{
			paymentManager: pm,
		},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.POST("/payments", h.RegisterPayment)
}
