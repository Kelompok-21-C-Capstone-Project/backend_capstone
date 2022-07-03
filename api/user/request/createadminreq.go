package request

import "backend_capstone/services/user/dto"

type RegisterAdminRequest struct {
	Name     string `json:"name" example:"some name" minLength:"5"`
	Username string `json:"username" example:"username" minLength:"5"`
	Password string `json:"password" example:"supersecret" minLength:"8"`
	Email    string `json:"email" example:"somemail@mail.com" format:"email"`
	Phone    string `json:"phone" example:"0833333333" minLength:"10"`
	Key      string `json:"key" example:"adminsupersecret"`
}

func (data *RegisterAdminRequest) DtoReq() dto.RegisterAdminDTO {
	return dto.RegisterAdminDTO{
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
		Key:      data.Key,
	}
}
