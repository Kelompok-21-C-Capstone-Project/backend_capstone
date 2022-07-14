package dto

import (
	"backend_capstone/models"
)

type CreateCategoryDTO struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	IsAvailable bool
	Icon        string
}

func (data *CreateCategoryDTO) GenerateModel(id string, slug string) *models.ProductCategory {
	return &models.ProductCategory{
		Id:          id,
		Slug:        slug,
		Name:        data.Name,
		IsAvailable: &data.IsAvailable,
		Description: data.Description,
		Icon:        data.Icon,
	}
}
