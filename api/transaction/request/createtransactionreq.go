package request

import "backend_capstone/services/transaction/dto"

type CreateTransactionRequest struct {
	ProductId       string `json:"product_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"`
	PaymentMethod   string `json:"payment_method" example:"Virtual Account"`
	MethodDetails   string `json:"payment_details" example:"Mandiri Virtual Account"`
	PaymentData     string `json:"payment_data,omitempty" example:"Misal nomer yang akan ditagih (nomer hp untuk gopay)"`
	TransactionData string `json:"transaction_data,omitempty" example:"Data seperti nomer hp/nomer padam yg menerima produk pembelian"`
	Email           string `json:"email" example:"testmail@mail.com"`
	Price           uint32 `json:"price" example:"300000"`
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
