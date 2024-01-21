package listener

import (
	"github.com/disturb16/go-service-struct/advanced/internal/listener/processors/auth"
	"github.com/disturb16/go-service-struct/advanced/internal/listener/processors/users"
	"go.uber.org/fx"
)

var Module = fx.Module("listener", fx.Provide(
	auth.New,
	users.New,
))
