package request

import "backend_capstone/services/user/dto"

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (data *UpdateUserRequest) DtoReq() dto.UpdateUserDTO {
	return dto.UpdateUserDTO{
		Name:     data.Name,
		Password: data.Password,
		Phone:    data.Phone,
	}
}
