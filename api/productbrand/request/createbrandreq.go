package request

import "backend_capstone/services/productbrand/dto"

type CreateBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (req *CreateBrandRequest) DtoReq() *dto.CreateBrandDTO {
	return &dto.CreateBrandDTO{
		Name:        req.Name,
		Description: req.Description,
	}
}
