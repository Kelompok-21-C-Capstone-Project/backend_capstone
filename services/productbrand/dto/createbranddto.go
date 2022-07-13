package dto

import (
	"backend_capstone/models"
)

type CreateBrandDTO struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	IsAvailable bool
}

func (data *CreateBrandDTO) GenerateModel(id string) *models.ProductBrand {
	return &models.ProductBrand{
		Id:          id,
		Name:        data.Name,
		IsAvailable: data.IsAvailable,
		Description: data.Description,
	}
}
