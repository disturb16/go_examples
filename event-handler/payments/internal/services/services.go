package services

import (
	"event-handler/payments/internal/services/paymentsmanager"

	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	paymentsmanager.New,
))
