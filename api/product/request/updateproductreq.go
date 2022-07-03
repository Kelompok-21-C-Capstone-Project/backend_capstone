package request

import "backend_capstone/services/product/dto"

type UpdateProductRequest struct {
	Name        string `json:"name" example:"Some Product"`
	Price       uint32 `json:"price" example:"30000"`
	Stock       uint32 `json:"stock" example:"20"`
	IsAvailable bool   `json:"status"  example:"true"`
	IsDiscount  bool   `json:"is_discount" example:"true"`
	Description string `json:"description" example:"some description"`
	BrandId     string `json:"brand_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	CategoryId  string `json:"category_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
}

func (data *UpdateProductRequest) DtoReq() *dto.UpdateProductDTO {
	return &dto.UpdateProductDTO{
		Name:        data.Name,
		Price:       data.Price,
		Stock:       data.Stock,
		IsAvailable: data.IsAvailable,
		IsDiscount:  data.IsDiscount,
		Description: data.Description,
		BrandId:     data.BrandId,
		CategoryId:  data.CategoryId,
	}
}
