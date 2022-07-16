package dto

import (
	"backend_capstone/models"
)

type UpdateProductDTO struct {
	Name        string `validate:"required"`
	Price       uint32 `validate:"required"`
	PriceBuy    uint32 `validate:"required"`
	Profit      uint32 `validate:"required"`
	Stock       uint32
	IsAvailable bool
	IsDiscount  bool
	Description string `validate:"required"`
	BrandId     string `validate:"required"`
	CategoryId  string `validate:"required"`
}

func (data *UpdateProductDTO) GenerateModel(id string, ProductBrandCategoryId string) *models.Product {
	return &models.Product{
		Id:                     id,
		Name:                   data.Name,
		Price:                  data.Price,
		PriceBuy:               data.PriceBuy,
		Profit:                 data.Profit,
		Stock:                  data.Stock,
		Description:            data.Description,
		IsAvailable:            &data.IsAvailable,
		IsDiscount:             &data.IsDiscount,
		ProductBrandCategoryId: ProductBrandCategoryId,
	}
}

type UpdateStockDTO struct {
	AdminId string                  `validate:"required"`
	Datas   []UpdateProductStockDTO `validate:"required"`
}

type UpdateProductStockDTO struct {
	Id    string `validate:"required"`
	Stock uint32 `validate:"required"`
}
