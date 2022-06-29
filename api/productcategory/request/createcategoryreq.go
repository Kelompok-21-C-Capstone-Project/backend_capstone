package request

import (
	"backend_capstone/services/productcategory/dto"
)

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (req *CreateCategoryRequest) DtoReq() *dto.CreateCategoryDTO {
	return &dto.CreateCategoryDTO{
		Name:        req.Name,
		Description: req.Description,
	}
}
