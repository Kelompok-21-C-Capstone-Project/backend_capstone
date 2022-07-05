package dto

import (
	"backend_capstone/models"
	"backend_capstone/utils/midtransdriver/dto"
)

type CreateTransactionDTO struct {
	ProductId       string `validate:"uuid,required"`
	PaymentMethod   string `validate:"required"`
	MethodDetails   string `validate:"required"`
	PaymentData     string
	TransactionData string
	Email           string `validate:"required,email"`
	Price           uint32 `validate:"required"`
}

func (data *CreateTransactionDTO) GenerateTransactionModel(id string, userId string) *models.Transaction {
	return &models.Transaction{
		Id:          id,
		UserId:      userId,
		ProductId:   data.ProductId,
		Description: data.TransactionData,
	}
}

func (data *CreateTransactionDTO) GeneratePaymentModel(id string, transactionId string, status string, method string, methodDetails string, description string) *models.Payment {
	return &models.Payment{
		Id:            id,
		TransactionId: transactionId,
		Amount:        data.Price,
		Status:        status,
		Method:        data.PaymentMethod,
		MethodDetails: data.MethodDetails,
		Description:   data.PaymentData,
	}
}

func (data *CreateTransactionDTO) GenerateMidtransPayment(orderId string) dto.MidtransPaymentDTO {
	return dto.MidtransPaymentDTO{
		MethodDetails: data.MethodDetails,
		OrderId:       orderId,
		Paid:          int64(data.Price),
		ItemPrice:     int64(data.Price),
		ItemId:        data.ProductId,
		ItemName:      data.ProductId,
	}
}
