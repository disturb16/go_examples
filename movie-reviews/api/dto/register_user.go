package dto

import models "github.com/disturb16/go_examples/movie-reviews/internal/users/model"

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *RegisterUser) ToUserModel() *models.Users {
	return &models.Users{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
