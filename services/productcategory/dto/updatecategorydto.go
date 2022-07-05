package dto

import (
	"backend_capstone/models"
)

type UpdateCategoryDTO struct {
	Name        string `validator:"required"`
	IsAvailable bool   `validator:"required,boolean"`
	Description string `validator:"required,alphaunicode"`
	Icon        string
}

func (data *UpdateCategoryDTO) GenerateModel(id string, slug string) *models.ProductCategory {
	return &models.ProductCategory{
		Id:          id,
		Slug:        slug,
		Name:        data.Name,
		IsAvailable: data.IsAvailable,
		Description: data.Description,
		Icon:        data.Icon,
	}
}
