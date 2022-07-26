package dto

import (
	"backend_capstone/models"
)

type CraeteProductDTO struct {
	Name        string `validate:"required"`
	Price       uint32 `validate:"required"`
	PriceBuy    uint32 `validate:"required"`
	Profit      uint32 `validate:"required"`
	Stock       uint32
	Description string `validate:"required"`
	BrandId     string `validate:"required"`
	CategoryId  string `validate:"required"`
	IsAvailable bool
}

func (data *CraeteProductDTO) GenerateModel(id string, ProductBrandCategoryId string) *models.Product {
	return &models.Product{
		Id:                     id,
		Name:                   data.Name,
		Price:                  data.Price,
		PriceBuy:               data.PriceBuy,
		Profit:                 data.Profit,
		Stock:                  data.Stock,
		Description:            data.Description,
		ProductBrandCategoryId: ProductBrandCategoryId,
		IsAvailable:            &data.IsAvailable,
	}
}
