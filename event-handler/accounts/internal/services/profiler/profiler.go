package profiler

import (
	"context"
	"errors"
	"event-handler/accounts/internal/repositories/profiles"
	"event-handler/accounts/logger"
	"slices"
)

type service struct {
	repo profiles.ProfileStorer
}

type Profile struct {
	AccountID      string `json:"account_id"`
	SubscriptionID string `json:"subscription_id"`
	Status         string `json:"status"`
}

type Manager interface {
	Register(ctx context.Context, accountID, subscriptionID string) error
	UpdateStatus(ctx context.Context, accountID, status string) error
	ByID(ctx context.Context, accountID string) (Profile, error)
}

var validProfileStatuses = []string{profiles.ProfileStatusActive, profiles.ProfileStatusPending}

var (
	ErrProfileAlreadyExists  = errors.New("profile already exists")
	ErrProfileStatusNotValid = errors.New("profile status is not valid")
)

func New(r profiles.ProfileStorer) Manager {
	return &service{
		repo: r,
	}
}

func (s *service) Register(ctx context.Context, accountID, subscriptionID string) error {
	log := logger.FromCtx(ctx)

	profile, err := s.repo.ByAccountID(ctx, accountID)
	if err != nil {
		log.WithError(err).Error("failed to verify profile exists")
		return err
	}

	if profile.AccountID != "" {
		log.WithAny("account_id", accountID).Error("cannot register profile that already exists")
		return ErrProfileAlreadyExists
	}

	return s.repo.Save(ctx, accountID, subscriptionID)
}

func (s *service) UpdateStatus(ctx context.Context, accountID, status string) error {
	if !slices.Contains(validProfileStatuses, status) {
		return ErrProfileStatusNotValid
	}

	err := s.repo.UpdateStatus(ctx, accountID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ByID(ctx context.Context, accountID string) (Profile, error) {
	p := Profile{}

	record, err := s.repo.ByAccountID(ctx, accountID)
	if err != nil {
		return p, err
	}

	p.AccountID = record.AccountID
	p.SubscriptionID = record.SubscriptionID
	p.Status = record.Status

	return p, nil
}
