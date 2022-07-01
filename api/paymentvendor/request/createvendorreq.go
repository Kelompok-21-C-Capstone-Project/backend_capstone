package request

import "backend_capstone/services/paymentvendor/dto"

type CreateVendorRequest struct {
	Name            string `json:"name"`
	Value           string `json:"value"`
	PaymentMethodId string `json:"method_id"`
}

func (data *CreateVendorRequest) DtoReq() *dto.CreateVendorDTO {
	return &dto.CreateVendorDTO{
		Name:            data.Name,
		Value:           data.Value,
		PaymentMethodId: data.PaymentMethodId,
	}
}
