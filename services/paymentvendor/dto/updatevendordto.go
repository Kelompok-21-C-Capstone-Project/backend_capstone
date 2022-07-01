package dto

import "backend_capstone/models"

type UpdateVendorDTO struct {
	Name            string `validate:"required"`
	Value           string `validate:"required"`
	IsAvailable     bool   `validate:"required"`
	PaymentMethodId string `validate:"uuid,required"`
}

func (data *UpdateVendorDTO) GenerateModel(id string) *models.PaymentVendor {
	return &models.PaymentVendor{
		Id:              id,
		Name:            data.Name,
		Value:           data.Value,
		IsAvailable:     data.IsAvailable,
		PaymentMethodId: data.PaymentMethodId,
	}
}
