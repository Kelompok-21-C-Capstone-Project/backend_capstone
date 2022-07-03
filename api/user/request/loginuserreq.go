package request

import "backend_capstone/services/user/dto"

type LoginUserRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func (data *LoginUserRequest) DtoReq() dto.LoginUserDTO {
	return dto.LoginUserDTO{
		Identifier: data.Identifier,
		Password:   data.Password,
	}
}
