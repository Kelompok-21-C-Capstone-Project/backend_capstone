package dto

import (
	"backend_capstone/models"
)

type CreateBrandDTO struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

func (data *CreateBrandDTO) GenerateModel(id string) *models.ProductBrand {
	return &models.ProductBrand{
		Id:          id,
		Name:        data.Name,
		IsAvailable: false,
		Description: data.Description,
	}
}
