package models

type PaymentVendor struct {
	Id              string `gorm:"type:varchar(50)"`
	Name            string `gorm:"type:varchar(30);not null"`
	Value           string `gorm:"type:varchar(30)"`
	BodyRequest     string `gorm:"type:text"`
	IsAvailable     bool   `gorm:"default:true"`
	IconPath        string `gorm:"type:text"`
	PaymentMethodId string `gorm:"type:varchar(50)"`
}
