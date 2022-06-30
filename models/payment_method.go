package models

type PaymentMethod struct {
	Id             string `gorm:"type:varchar(50)"`
	Name           string `gorm:"type:varchar(50);not null"`
	Slug           string `gorm:"type:varchar(50);not null"`
	PaymentVendors []PaymentVendor
}
