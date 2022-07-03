package request

import "backend_capstone/services/productcategory/dto"

type UpdateCategoryRequest struct {
	Name        string `json:"name" example:"some name"`
	IsAvailable bool   `json:"status" example:"true"`
	Description string `json:"description" example:"some description"`
}

func (req *UpdateCategoryRequest) DtoReq() *dto.UpdateCategoryDTO {
	return &dto.UpdateCategoryDTO{
		Name:        req.Name,
		IsAvailable: req.IsAvailable,
		Description: req.Description,
	}
}
