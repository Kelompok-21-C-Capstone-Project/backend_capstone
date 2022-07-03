package models

import (
	"time"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

//> model tipe data
type Transaction struct {
	Id          string    `gorm:"type:varchar(50)" json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	UserId      string    `gorm:"type:varchar(50)" json:"user_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	ProductId   string    `gorm:"type:varchar(50)" json:"product_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Payment     Payment   `json:"payment,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty" example:"some description"`
	CreatedAt   time.Time `json:"date,omitempty"`
	UpdatedAt   time.Time `json:"updated,omitempty"`
}
