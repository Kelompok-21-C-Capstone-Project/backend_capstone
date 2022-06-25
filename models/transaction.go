package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

//> model tipe data
type Transaction struct {
	Id          uuid.UUID `gorm:"type:varchar(50)" json:"id"`
	UserId      uuid.UUID `gorm:"type:varchar(50)" json:"user_id"`
	ProductId   uuid.UUID `gorm:"type:varchar(50)" json:"product_id"`
	Payment     Payment
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"date"`
	UpdatedAt   time.Time `json:"updated"`
}
