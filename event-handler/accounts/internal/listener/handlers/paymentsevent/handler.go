package paymentsevent

import (
	"context"
	"encoding/json"
	"event-handler/accounts/internal/contracts"
	"event-handler/accounts/internal/events"
	"event-handler/accounts/internal/services/profiler"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

type handler struct {
	profilesManager profiler.Manager
}

type Result struct {
	fx.Out

	Handler contracts.EventHandler `group:"event_handlers"`
}

func New(pm profiler.Manager) Result {
	return Result{
		Handler: &handler{
			profilesManager: pm,
		},
	}
}

func (h *handler) Name() string {
	return "accounts_payments_handler"
}

func (h *handler) Topic() string {
	return "payments"
}

func (h *handler) Handle(msg *message.Message) error {
	ev := events.PaymentCreated{}

	err := json.Unmarshal(msg.Payload, &ev)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = h.profilesManager.UpdateStatus(ctx, ev.AccountID, "active")
	if err != nil {
		return err
	}

	return nil
}
