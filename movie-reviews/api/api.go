package api

import "github.com/disturb16/go_examples/movie-reviews/internal/users"

type API struct {
	userService users.Service
}

func NewApi(userService users.Service) *API {
	return &API{userService}
}
