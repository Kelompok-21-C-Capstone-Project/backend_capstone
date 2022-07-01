package request

import "backend_capstone/services/paymentmethod/dto"

type CreateUpdateMethodRequest struct {
	Name string `json:"name"`
}

func (data *CreateUpdateMethodRequest) DtoReq() *dto.CreateMethodDTO {
	return &dto.CreateMethodDTO{
		Name: data.Name,
	}
}
