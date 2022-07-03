package request

import "backend_capstone/services/productbrand/dto"

type UpdateBrandRequest struct {
	Name        string `json:"name" example:"XL"`
	IsAvailable bool   `json:"status" example:"true"`
	Description string `json:"description" example:"Excelent Celullar"`
}

func (req *UpdateBrandRequest) DtoReq() *dto.UpdateBrandDTO {
	return &dto.UpdateBrandDTO{
		Name:        req.Name,
		IsAvailable: req.IsAvailable,
		Description: req.Description,
	}
}
