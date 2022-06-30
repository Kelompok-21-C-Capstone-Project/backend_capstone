package dto

import (
	"backend_capstone/models"
)

type CraeteProductDTO struct {
	Name        string `validate:"required"`
	Price       uint32 `validate:"required"`
	Stock       uint32 `validate:"required"`
	Description string `validate:"required"`
	BrandId     string `validate:"required"`
	CategoryId  string `validate:"required"`
}

func (data *CraeteProductDTO) GenerateModel(id string, ProductBrandCategoryId string) *models.Product {
	return &models.Product{
		Id:                     id,
		Name:                   data.Name,
		Price:                  data.Price,
		Stock:                  data.Stock,
		Description:            data.Description,
		ProductBrandCategoryId: ProductBrandCategoryId,
	}
}
