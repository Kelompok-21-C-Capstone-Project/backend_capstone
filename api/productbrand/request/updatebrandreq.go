package request

import "backend_capstone/services/productbrand/dto"

type UpdateBrandRequest struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"status"`
	Description string `json:"description"`
}

func (req *UpdateBrandRequest) DtoReq() *dto.UpdateBrandDTO {
	return &dto.UpdateBrandDTO{
		Name:        req.Name,
		IsAvailable: req.IsAvailable,
		Description: req.Description,
	}
}
