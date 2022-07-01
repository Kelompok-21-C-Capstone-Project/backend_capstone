package dto

import "backend_capstone/models"

type CreateVendorDTO struct {
	Name            string `validate:"required"`
	Value           string `validate:"required"`
	PaymentMethodId string `validate:"required"`
}

func (data *CreateVendorDTO) GenerateModel(id string) *models.PaymentVendor {
	return &models.PaymentVendor{
		Id:              id,
		Name:            data.Name,
		Value:           data.Value,
		PaymentMethodId: data.PaymentMethodId,
	}
}
