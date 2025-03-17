package paymentsmanager

import (
	"context"
	"encoding/json"
	"errors"
	"event-handler/payments/internal/events"
	"event-handler/payments/internal/repositories/paymentstore"
	"event-handler/payments/logger"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type service struct {
	repo            paymentstore.PaymentStorer
	eventsPublisher *amqp.Publisher
}

type Manager interface {
	RegisterPayment(ctx context.Context, accountID string, total int64) error
}

var ErrInvalidPaymentAmount = errors.New("invalid payment amount")

func New(repo paymentstore.PaymentStorer, ep *amqp.Publisher) Manager {
	return &service{
		repo:            repo,
		eventsPublisher: ep,
	}
}

func (s *service) RegisterPayment(ctx context.Context, accountID string, total int64) error {
	if total < 0 {
		return ErrInvalidPaymentAmount
	}

	log := logger.FromCtx(ctx)
	paymentID := uuid.New().String()
	err := s.repo.Save(ctx, accountID, paymentID, total)
	if err != nil {
		return err
	}

	err = s.sendPaymentCreatedEvent(accountID, paymentID, total)
	if err != nil {
		log.WithError(err).Error("failed to send payment created event")
	}

	return nil
}

func (s *service) sendPaymentCreatedEvent(accountID, paymentID string, total int64) error {
	ev := events.PaymentCreated{
		AccountID: accountID,
		PaymentID: paymentID,
		TotalPaid: total,
		CreatedAt: time.Now().In(time.UTC).Format(time.DateOnly),
	}

	payload, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)

	err = s.eventsPublisher.Publish("payments", msg)
	if err != nil {
		return err
	}

	return nil
}
