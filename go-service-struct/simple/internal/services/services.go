package services

import (
	"github.com/disturb16/go-service-struct/simple/internal/services/auth"
	"github.com/disturb16/go-service-struct/simple/internal/services/profiles"
	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	auth.New,
	profiles.New,
))
