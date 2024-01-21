package users

import (
	"msgbroker"

	"github.com/disturb16/go-service-struct/advanced/internal/services/profiles"
	"go.uber.org/fx"
)

type processor struct {
	profileService profiles.Service
}

type Result struct {
	fx.Out

	Procesor msgbroker.Processor `group:"procesors"`
}

func New(service profiles.Service) Result {
	return Result{
		Procesor: &processor{
			profileService: service,
		},
	}
}

func (p *processor) Register(b msgbroker.Broker) {
	b.AddProcessor("user-update", p.OnUserUpdate)
	b.AddProcessor("user-deleted", p.onUserDelete)
}
