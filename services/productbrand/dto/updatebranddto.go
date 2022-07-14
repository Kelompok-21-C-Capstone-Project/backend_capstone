package dto

import (
	"backend_capstone/models"
)

type UpdateBrandDTO struct {
	Name        string `json:"name" validator:"required"`
	IsAvailable bool   `json:"is_available" validator:"required,boolean"`
	Description string `json:"description" validator:"required,alphaunicode"`
}

func (data *UpdateBrandDTO) GenerateModel(id string) *models.ProductBrand {
	return &models.ProductBrand{
		Id:          id,
		Name:        data.Name,
		IsAvailable: &data.IsAvailable,
		Description: data.Description,
	}
}
