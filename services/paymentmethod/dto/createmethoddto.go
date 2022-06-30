package dto

import "backend_capstone/models"

type CreateMethodDTO struct {
	Name string `validate:"required"`
}

func (data *CreateMethodDTO) GenerateModel(id string, slug string) *models.PaymentMethod {
	return &models.PaymentMethod{
		Id:   id,
		Slug: slug,
		Name: data.Name,
	}
}
