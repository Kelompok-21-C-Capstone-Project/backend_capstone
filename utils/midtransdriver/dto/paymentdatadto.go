package dto

type MidtransPaymentDTO struct {
	MethodDetails string
	OrderId       string
	Paid          int64
	ItemId        string
	ItemPrice     int64
	ItemName      string
}
