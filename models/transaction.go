package models

import (
	"time"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

//> model tipe data
type Transaction struct {
	Id          string `gorm:"type:varchar(50)" json:"id"`
	UserId      string `gorm:"type:varchar(50)" json:"user_id"`
	ProductId   string `gorm:"type:varchar(50)" json:"product_id"`
	Payment     Payment
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"date"`
	UpdatedAt   time.Time `json:"updated"`
}
