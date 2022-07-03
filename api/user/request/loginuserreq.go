package request

import "backend_capstone/services/user/dto"

type LoginUserRequest struct {
	Identifier string `json:"identifier" example:"username" minLength:"5"`
	Password   string `json:"password" example:"supersecret" minLength:"8"`
}

func (data *LoginUserRequest) DtoReq() dto.LoginUserDTO {
	return dto.LoginUserDTO{
		Identifier: data.Identifier,
		Password:   data.Password,
	}
}
