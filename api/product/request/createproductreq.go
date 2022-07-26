package request

import "backend_capstone/services/product/dto"

type CreateProductRequest struct {
	Name        string `json:"name" example:"Some Product"`
	Price       uint32 `json:"harga_jual" example:"30000"`
	PriceBuy    uint32 `json:"harga_beli" example:"30000"`
	Profit      uint32 `json:"komisi" example:"30000"`
	Stock       uint32 `json:"stock" example:"20"`
	IsAvailable bool   `json:"is_available"  example:"true"`
	Description string `json:"description" example:"some description"`
	BrandId     string `json:"brand_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	CategoryId  string `json:"category_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
}

func (data *CreateProductRequest) DtoReq() *dto.CraeteProductDTO {
	return &dto.CraeteProductDTO{
		Name:        data.Name,
		Price:       data.Price,
		PriceBuy:    data.PriceBuy,
		Profit:      data.Profit,
		Stock:       data.Stock,
		Description: data.Description,
		BrandId:     data.BrandId,
		CategoryId:  data.CategoryId,
		IsAvailable: data.IsAvailable,
	}
}
