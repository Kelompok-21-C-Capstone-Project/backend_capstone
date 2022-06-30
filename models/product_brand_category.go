package models

type ProductBrandCategory struct {
	Id                string `gorm:"type:varchar(50)"`
	ProductBrandId    string `gorm:"type:varchar(50);not null"`
	ProductCategoryId string `gorm:"type:varchar(50);not null"`
	Slug              string `gorm:"type:varchar(50);not null"`
	IsAvailable       bool   `gorm:"default:true;not null"`
	Products          []Product
}
