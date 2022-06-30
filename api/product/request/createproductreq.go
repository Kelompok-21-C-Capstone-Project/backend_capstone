package request

import "backend_capstone/services/product/dto"

type CreateProductRequest struct {
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Stock       uint32 `json:"stock"`
	Description string `json:"description"`
	BrandId     string `json:"brand_id"`
	CategoryId  string `json:"category_id"`
}

func (data *CreateProductRequest) DtoReq() *dto.CraeteProductDTO {
	return &dto.CraeteProductDTO{
		Name:        data.Name,
		Price:       data.Price,
		Stock:       data.Stock,
		Description: data.Description,
		BrandId:     data.BrandId,
		CategoryId:  data.CategoryId,
	}
}
