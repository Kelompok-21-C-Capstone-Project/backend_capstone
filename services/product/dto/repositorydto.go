package dto

import (
	"time"
)

type ResponseBodyProduct struct {
	PageLength int       `json:"page_length" example:"3"`
	Data       []Product `json:"datas"`
}

type Product struct {
	Id                     string    `json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"`
	Name                   string    `json:"name,omitempty" example:"Some Product"`
	Price                  uint32    `json:"harga_jual,omitempty" example:"30000"`
	PriceBuy               uint32    `json:"harga_beli,omitempty" example:"30000"`
	Profit                 uint32    `json:"komisi,omitempty" example:"30000"`
	Stock                  uint32    `json:"stock,omitempty" example:"30"`
	Discount               float32   `json:"discount,omitempty"`
	IsAvailable            *bool     `json:"is_available" example:"true"`
	IsDiscount             *bool     `json:"is_discount,omitempty" example:"false"`
	Description            string    `json:"description,omitempty" example:"some description"`
	ProductBrandCategoryId string    `json:"category_id,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
