package users

import models "github.com/disturb16/go_examples/movie-reviews/internal/users/model"

//go:generate mockery --name=Repository --output=internal/users --inpackage=true
type Repository interface {
	Save(user *models.Users) error
	Login(email, password string) (*models.Users, error)
}

//go:generate mockery --name=Service --output=internal/users --inpackage=true
type Service interface {
	Register(user *models.Users) error
	Login(email, password string) (*models.Users, error)
}
