package repositories

import (
	"event-handler/payments/internal/repositories/paymentstore"

	"go.uber.org/fx"
)

var Module = fx.Module("repositories", fx.Provide(
	paymentstore.New,
))
