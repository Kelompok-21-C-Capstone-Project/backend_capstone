package request

import "backend_capstone/services/paymentvendor/dto"

type UpdateVendorRequest struct {
	Name            string `json:"name"`
	Value           string `json:"value"`
	IsAvailable     bool   `json:"status"`
	PaymentMethodId string `json:"method_id"`
}

func (data *UpdateVendorRequest) DtoReq() *dto.UpdateVendorDTO {
	return &dto.UpdateVendorDTO{
		Name:            data.Name,
		Value:           data.Value,
		IsAvailable:     data.IsAvailable,
		PaymentMethodId: data.PaymentMethodId,
	}
}
