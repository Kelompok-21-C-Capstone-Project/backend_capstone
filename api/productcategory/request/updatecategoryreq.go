package request

import "backend_capstone/services/productcategory/dto"

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"status"`
	Description string `json:"description"`
}

func (req *UpdateCategoryRequest) DtoReq() *dto.UpdateCategoryDTO {
	return &dto.UpdateCategoryDTO{
		Name:        req.Name,
		IsAvailable: req.IsAvailable,
		Description: req.Description,
	}
}
