package dto

import "backend_capstone/models"

type CreateTransactionDTO struct {
	ProductId       string `validate:"uuid,required"`
	PaymentMethod   string `validate:"required"`
	MethodDetails   string `validate:"uuid,required"`
	PaymentData     string `validate:"alphanum,min=8,max=32"`
	TransactionData string `validate:"alphanum,min=8,max=32"`
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
