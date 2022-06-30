package request

import "backend_capstone/services/product/dto"

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Stock       uint32 `json:"stock"`
	IsAvailable bool   `json:"status"`
	IsDiscount  bool   `json:"discount_status"`
	Description string `json:"description"`
	BrandId     string `json:"brand_id"`
	CategoryId  string `json:"category_id"`
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
