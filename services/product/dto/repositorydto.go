package dto

import (
	"time"
)

type ResponseBodyProduct struct {
	PageLength int       `json:"page_length" example:"3"`
	Data       []Product `json:"datas"`
}

type Product struct {
	Id          string    `json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"`
	Name        string    `json:"name,omitempty" example:"Some Product"`
	Category    string    `json:"category,omitempty" example:"Some Product"`
	Brand       string    `json:"brand,omitempty" example:"Some Product"`
	Price       uint32    `json:"price,omitempty" example:"30000"`
	Stock       uint32    `json:"stock,omitempty" example:"30"`
	Discount    float32   `json:"discount,omitempty"`
	IsAvailable bool      `json:"is_available" example:"true"`
	IsDiscount  bool      `json:"is_discount,omitempty" example:"false"`
	Description string    `json:"description,omitempty" example:"some description"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Deleted     time.Time `json:"deleted,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
