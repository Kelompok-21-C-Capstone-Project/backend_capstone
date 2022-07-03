package request

import "backend_capstone/services/product/dto"

type CreateProductRequest struct {
	Name        string `json:"name" example:"Some Product"`
	Price       uint32 `json:"price" example:"30000"`
	Stock       uint32 `json:"stock" example:"20"`
	Description string `json:"description" example:"some description"`
	BrandId     string `json:"brand_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	CategoryId  string `json:"category_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
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
