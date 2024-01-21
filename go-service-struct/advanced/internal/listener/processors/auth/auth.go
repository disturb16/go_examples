package auth

import (
	"msgbroker"

	"go.uber.org/fx"
)

type processor struct {
	Broker msgbroker.Broker
}

type Result struct {
	fx.Out

	Procesor msgbroker.Processor `group:"procesors"`
}

func New() Result {
	return Result{
		Procesor: &processor{},
	}
}

func (p *processor) Register(b msgbroker.Broker) {
	b.AddProcessor("user-logged-out", p.onLogout)
}
