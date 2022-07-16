package dto

import "backend_capstone/models"

type UpdateUserDTO struct {
	Name     string `validate:"required,min=5"`
	Password string `validate:"required,min=8"`
	Phone    string `validate:"required,min=10"`
}

func (data *UpdateUserDTO) GenerateModel() *models.User {
	return &models.User{
		Name:     data.Name,
		Phone:    data.Phone,
		Password: data.Password,
	}
}
