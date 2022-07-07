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
	Id             string    `json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	TransactionId  string    `json:"transaction_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Status         string    `json:"status" example:"Pending"`
	VaNumber       string    `json:"va_number,omitempty" example:"2312312412523"`
	PaymentDetails string    `json:"payment_method,omitempty" example:"Mandiri Virtual Account"`
	Billed         uint32    `json:"bill,omitempty" example:"11500"`
	Charger        uint32    `json:"charged,omitempty" example:"1500"`
	Product        string    `json:"product_name,omitempty" example:"Telkomsel 5GB Super Tiktok"`
	ProductPrice   uint32    `json:"product_price,omitempty" example:"10000"`
	Deadline       time.Time `json:"dead_line,omitempty" example:"2022-07-05T14:32:46.996074+07"`
}

func (BillClient) TableName() string {
	return "transactions"
}

type ClientTransactionsResponse struct {
	Id                 string    `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Category           string    `json:"type" example:"paket-data"`
	Status             string    `json:"status" example:"Pending"`
	Product            string    `json:"label" example:"Telkomsel 5GB Super Tiktok"`
	TransactionDetails string    `json:"phone_number" example:"2312312412523"`
	Charged            uint32    `json:"price" example:"11500"`
	CreatedAt          time.Time `json:"created_at" example:"2022-07-05T14:32:46.996074+07"`
	PaymentMethod      string    `json:"payment_method" example:"Virtual Account"`
}

func (ClientTransactionsResponse) TableName() string {
	return ""
}

type MidtransAfterPayment struct {
	TransactionId string `json:"order_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Status        string `json:"transaction_status,omitempty" example:"Pending"`
	Code          string `json:"status_code,omitempty"`
}
