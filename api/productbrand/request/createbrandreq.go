package request

import "backend_capstone/services/productbrand/dto"

type CreateBrandRequest struct {
	Name        string `json:"name" example:"Telkomsel"`
	Description string `json:"description" example:"Telekomunikasi"`
}

func (req *CreateBrandRequest) DtoReq() *dto.CreateBrandDTO {
	return &dto.CreateBrandDTO{
		Name:        req.Name,
		Description: req.Description,
	}
}
