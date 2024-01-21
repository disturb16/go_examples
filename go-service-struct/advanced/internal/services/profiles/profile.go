package profiles

import (
	"context"

	"github.com/disturb16/go-service-struct/advanced/internal/repositories/profiles"
)

type Service interface {
	Update(ctx context.Context, username string) error
	Get(id string) (string, error)
}

type service struct {
	repository profiles.Repository
}

func New(repository profiles.Repository) Service {
	return &service{
		repository: repository,
	}
}
