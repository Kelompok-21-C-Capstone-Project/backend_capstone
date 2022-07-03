package request

import "backend_capstone/services/user/dto"

type RegisterAdminRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Key      string `json:"key"`
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
