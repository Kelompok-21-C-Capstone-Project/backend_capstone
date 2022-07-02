package request

import "backend_capstone/services/user/dto"

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
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
