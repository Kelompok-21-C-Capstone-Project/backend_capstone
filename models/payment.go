package models

import "time"

//> model tipe data
type Payment struct {
	PaymentId   int       `gorm : "primaryKey" json:"id"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	PaymentBy   string    `json:"payment"`
	UpdatedAt   time.Time `json:"updated"`
}

func (Payment) TableName() string {
	return ("payment")
}
