package dto

import (
	"backend_capstone/models"
)

type CreateCategoryDTO struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

func (data *CreateCategoryDTO) GenerateModel(id string, slug string) *models.ProductCategory {
	return &models.ProductCategory{
		Id:          id,
		Slug:        slug,
		Name:        data.Name,
		IsAvailable: false,
		Description: data.Description,
	}
}
