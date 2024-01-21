package httpapi

import (
	"github.com/disturb16/go-service-struct/advanced/internal/httpapi/handlers/auth"
	"github.com/disturb16/go-service-struct/advanced/internal/httpapi/handlers/profiles"
	"go.uber.org/fx"
)

var Module = fx.Module("httpapi", fx.Provide(
	auth.New,
	profiles.New,
))
