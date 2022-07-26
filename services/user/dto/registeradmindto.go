package dto

import "backend_capstone/models"

type RegisterAdminDTO struct {
	Name     string `validate:"required,min=5"`
	Username string `validate:"required,min=5"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,min=10"`
	Key      string `validate:"required"`
}

func (data *RegisterAdminDTO) GenerateModel(id string) *models.User {
	return &models.User{
		Id:       id,
		Username: data.Username,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Password: data.Password,
		Role:     "admin",
	}
}
