package request

import "backend_capstone/services/user/dto"

type UpdateUserRequest struct {
	Name     string `json:"name" example:"some name"`
	Password string `json:"password" example:"supersecret"`
	Phone    string `json:"phone" example:"0822222222"`
}

func (data *UpdateUserRequest) DtoReq() dto.UpdateUserDTO {
	return dto.UpdateUserDTO{
		Name:     data.Name,
		Password: data.Password,
		Phone:    data.Phone,
	}
}
