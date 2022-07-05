package request

import (
	"backend_capstone/services/productcategory/dto"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" example:"soma name"`
	Description string `json:"description" example:"some description"`
	Icon        string `json:"icon" example:"mdi-data-icon"`
}

func (req *CreateCategoryRequest) DtoReq() *dto.CreateCategoryDTO {
	return &dto.CreateCategoryDTO{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}
}
