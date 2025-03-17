package listener

import (
	"context"
	"event-handler/accounts/config"
	"event-handler/accounts/internal/contracts"
	"event-handler/accounts/internal/listener/handlers/paymentsevent"
	"event-handler/accounts/logger"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"go.uber.org/fx"
)

type eventsListener struct {
	router     *message.Router
	subscriber *amqp.Subscriber
}

type EventsListenerRouter interface {
	Start(context.Context) error
	RegisterHandler(h contracts.EventHandler)
	Stop() error
}

var Module = fx.Module("evetslistener", fx.Provide(
	paymentsevent.New,
))

func New(c *config.Config, log *logger.Logger) (EventsListenerRouter, error) {
	wlog := watermill.NewSlogLogger(log.Logger)
	router, err := message.NewRouter(message.RouterConfig{}, wlog)
	if err != nil {
		return nil, err
	}

	amqpConfig := amqp.NewDurableQueueConfig(c.RabbitMQAddr)
	subscriber, err := amqp.NewSubscriber(amqpConfig, wlog)
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(middleware.CorrelationID)
	router.AddMiddleware(middleware.Recoverer)
	router.AddMiddleware(middleware.Retry{
		MaxRetries:      5,
		InitialInterval: time.Millisecond * 100,
		Multiplier:      1.5,
		Logger:          wlog,
	}.Middleware)

	el := &eventsListener{
		router:     router,
		subscriber: subscriber,
	}

	return el, nil
}

func (el *eventsListener) Start(ctx context.Context) error {
	return el.router.Run(ctx)
}

func (el *eventsListener) Stop() error {
	return el.router.Close()
}

func (el *eventsListener) RegisterHandler(h contracts.EventHandler) {
	el.router.AddNoPublisherHandler(h.Name(), h.Topic(), el.subscriber, h.Handle)
}
