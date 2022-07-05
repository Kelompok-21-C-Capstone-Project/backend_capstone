package dto

import "time"

type Payment struct {
	Id                   string `json:"id,omitempty"`
	VirtualAccountNumber string `json:"va_number,omitempty"`
	PaidPrice            uint32 `json:"paid_price,omitempty"`
	Detail               TransactionDetails
}

type TransactionDetails struct {
	PaymentId string
	Price     uint32 `json:"price,omitempty"`
	PaidPrice uint32 `json:"paid_price,omitempty"`
	Vendor    string `json:"vendor_name,omitempty"`
	Product   string `json:"product_name,omitempty"`
}

type BillClient struct {
	Id             string    `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	TransactionId  string    `json:"transaction_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	VaNumber       string    `json:"va_number" example:"2312312412523"`
	PaymentDetails string    `json:"payment_method" example:"Mandiri Virtual Account"`
	Billed         uint32    `json:"bill" example:"11500"`
	Charger        uint32    `json:"charged" example:"1500"`
	Product        string    `json:"product_name" example:"Telkomsel 5GB Super Tiktok"`
	ProductPrice   uint32    `json:"product_price" example:"10000"`
	Deadline       time.Time `json:"dead_line" example:"2022-07-05T14:32:46.996074+07"`
}

func (BillClient) TableName() string {
	return "transactions"
}
