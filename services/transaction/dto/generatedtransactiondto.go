package dto

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
