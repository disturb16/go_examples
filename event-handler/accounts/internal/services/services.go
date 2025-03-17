package services

import (
	"event-handler/accounts/internal/services/profiler"

	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	profiler.New,
))
