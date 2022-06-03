package models

import "time"

//> model tipe data
type Product struct {
	ProductId   int       `gorm:"primaryKey" json:"id"`
	ProductName string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated"`
}

func (Product) TableName() string {
	return ("product")
}
