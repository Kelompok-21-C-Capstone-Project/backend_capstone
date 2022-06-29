package dto

import (
	"backend_capstone/models"

	"github.com/google/uuid"
)

type UpdateBrandDTO struct {
	Name        string `validator:"required"`
	IsAvailable bool   `validator:"required,boolean"`
	Description string `validator:"required,alphaunicode"`
}

func (data *UpdateBrandDTO) GenerateModel(id uuid.UUID) *models.ProductBrand {
	return &models.ProductBrand{
		Id:          id,
		Name:        data.Name,
		IsAvailable: data.IsAvailable,
		Description: data.Description,
	}
}
