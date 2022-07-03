package models

type ProductCategory struct {
	Id          string `gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name        string `gorm:"type:varchar(50);not null,unique" example:"Some Name"`
	Slug        string `gorm:"type:varchar(50);not null" example:"some-name"`
	IsAvailable bool   `gorm:"default:false;not null" example:"true"`
	Description string `gorm:"type:text" example:"some description"`
}
