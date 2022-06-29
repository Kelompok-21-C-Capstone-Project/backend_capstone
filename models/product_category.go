package models

import "github.com/google/uuid"

type ProductCategory struct {
	Id          uuid.UUID `gorm:"type:varchar(50);not null"`
	Name        string    `gorm:"type:varchar(50);not null,unique"`
	Slug        string    `gorm:"type:varchar(50);not null"`
	IsAvailable bool      `gorm:"default:false;not null"`
	Description string    `gorm:"type:text"`
}
