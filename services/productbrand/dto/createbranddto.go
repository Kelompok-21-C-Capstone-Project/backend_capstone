package dto

import (
	"backend_capstone/models"

	"github.com/google/uuid"
)

type CreateBrandDTO struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

func (data *CreateBrandDTO) GenerateModel(id uuid.UUID) *models.ProductBrand {
	return &models.ProductBrand{
		Id:          id,
		Name:        data.Name,
		IsAvailable: false,
		Description: data.Description,
	}
}
