package models

type PaymentVendor struct {
	Id              string `gorm:"type:varchar(50)"`
	Name            string `gorm:"type:varchar(30);not null"`
	BodyRequest     string `gorm:"type:text;not null"`
	IconPath        string `gorm:"type:text"`
	PaymentMethodId string `gorm:"type:varchar(50)"`
}
