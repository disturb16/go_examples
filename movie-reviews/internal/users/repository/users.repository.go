package repository

import (
	"github.com/disturb16/go_examples/movie-reviews/internal/users"
	models "github.com/disturb16/go_examples/movie-reviews/internal/users/model"
	"github.com/jmoiron/sqlx"
)

const (
	// CreateUserQuery is the query to create a user
	createUserQuery = `
		INSERT INTO users (name, email, password)
		VALUES (:name ,:email, :password)
	`

	// GetUserByEmailQuery is the query to get a user by email
	getUserByEmailQuery = `
		SELECT
			id,
			name,
			email,
			password
		FROM users
		WHERE email = $1
			AND password = $2
	`
)

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) users.Repository {
	return &repository{db}
}

func (r *repository) Save(user *models.Users) error {
	_, err := r.db.NamedExec(createUserQuery, user)
	return err
}

func (r *repository) Login(email, password string) (*models.Users, error) {
	user := &models.Users{}
	err := r.db.Get(user, getUserByEmailQuery, email, password)
	return user, err
}
