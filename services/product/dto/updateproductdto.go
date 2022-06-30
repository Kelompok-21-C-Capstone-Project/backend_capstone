package dto

import (
	"backend_capstone/models"

	"github.com/google/uuid"
)

type UpdateProductDTO struct {
	Name        string `validate:"required"`
	Price       uint32 `validate:"required"`
	Stock       uint32 `validate:"required"`
	IsAvailable bool   `validate:"required"`
	IsDiscount  bool   `validate:"required"`
	Description string `validate:"required"`
	BrandId     string `validate:"required"`
	CategoryId  string `validate:"required"`
}

func (data *UpdateProductDTO) GenerateModel(id uuid.UUID, ProductBrandCategoryId uuid.UUID) *models.Product {
	return &models.Product{
		Id:                     id,
		Name:                   data.Name,
		Price:                  data.Price,
		Stock:                  data.Stock,
		Description:            data.Description,
		IsAvailable:            data.IsAvailable,
		IsDiscount:             data.IsDiscount,
		ProductBrandCategoryId: ProductBrandCategoryId,
	}
}
