package service

import (
	"errors"

	"github.com/disturb16/go_examples/movie-reviews/internal/users"
	models "github.com/disturb16/go_examples/movie-reviews/internal/users/model"
)

var (
	UserNotFoundError = errors.New("User not found")
)

type service struct {
	repo users.Repository
}

// NewUserService returns a new user service impolementation.
func NewUserService(repo users.Repository) users.Service {
	return &service{repo}
}

// Register creates a new user.
func (s *service) Register(user *models.Users) error {

	// TODO: Check if the user already exists

	return s.repo.Save(user)
}

// Login returns a user if the credentials are correct.
func (s *service) Login(email, password string) (*models.Users, error) {

	u, err := s.repo.Login(email, password)
	if err != nil {
		return nil, UserNotFoundError
	}

	return u, nil
}
