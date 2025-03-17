package repositories

import (
	"event-handler/accounts/internal/repositories/profiles"

	"go.uber.org/fx"
)

var Module = fx.Module("repositories", fx.Provide(
	profiles.New,
))
