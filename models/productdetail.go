package models

import "time"

//> model tipe data
type Detail struct {
	ProductDetailId int       `gorm : "primaryKey" json:"id"`
	ProductId       int       `json:"id"`
	ProductList     string    `json:"list"`
	Description     string    `json:"description"`
	StockProduct    int       `json:"stock"`
	UpdatedAt       time.Time `json:"updated"`
}

func (Detail) TableName() string {
	return ("product detail")
}
