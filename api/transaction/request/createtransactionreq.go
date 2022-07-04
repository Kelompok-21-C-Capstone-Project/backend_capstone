package request

import "backend_capstone/services/transaction/dto"

type CreateTransactionRequest struct {
	ProductId       string `json:"product_id"`
	PaymentMethod   string `json:"payment_method"`
	MethodDetails   string `json:"payment_details"`
	PaymentData     string `json:"payment_data"`
	TransactionData string `json:"transaction_data"`
	Email           string `json:"email"`
	Price           uint32 `json:"price"`
}

func (data *CreateTransactionRequest) DtoReq() dto.CreateTransactionDTO {
	return dto.CreateTransactionDTO{
		ProductId:       data.ProductId,
		PaymentMethod:   data.PaymentMethod,
		MethodDetails:   data.MethodDetails,
		PaymentData:     data.PaymentData,
		TransactionData: data.TransactionData,
		Email:           data.Email,
		Price:           data.Price,
	}
}
