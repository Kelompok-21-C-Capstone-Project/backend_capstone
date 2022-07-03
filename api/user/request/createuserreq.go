package request

import "backend_capstone/services/user/dto"

type RegisterUserRequest struct {
	Name     string `json:"name" example:"some name" minLength:"5"`
	Username string `json:"username" example:"some username" minLength:"5"`
	Password string `json:"password" example:"supersecret" minLength:"8"`
	Email    string `json:"email" example:"test@mail.com" format:"email"`
	Phone    string `json:"phone" example:"081111111" minLength:"10"`
}

func (data *RegisterUserRequest) DtoReq() dto.RegisterUserDTO {
	return dto.RegisterUserDTO{
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
	}
}
