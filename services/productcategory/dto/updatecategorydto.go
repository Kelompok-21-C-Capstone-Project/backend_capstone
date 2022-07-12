package dto

import (
	"backend_capstone/models"
)

type UpdateCategoryDTO struct {
	Name        string `json:"name" validator:"required"`
	IsAvailable bool   `json:"is_available" validator:"required,boolean"`
	Description string `json:"description" validator:"required,alphaunicode"`
	Icon        string `json:"icon"`
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
