package auth

import (
	"context"

	"github.com/disturb16/go-service-struct/advanced/internal/repositories/users"
)

type Service interface {
	Authenticate(ctx context.Context, username, password string) error
}

type service struct {
	repo users.Repository
}

func New(repo users.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Authenticate(ctx context.Context, username, password string) error {
	return nil
}
