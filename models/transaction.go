package models

import "time"

//> model tipe data
type Transaction struct {
	TransactionId     int       `gorm:"primaryKey" json:"id"`
	UserId            int       `json:"user id"`
	PaymentId         int       `json:"payment_id"`
	ProductId         int       `json:"product_id"`
	Description       string    `json:"description"`
	TransactionDate   time.Time `json:"date"`
	TransactionDetail string    `json:"transaction detail"`
	UpdatedAt         time.Time `json:"updated"`
}

func (Transaction) TableName() string {
	return "transaction"
}
