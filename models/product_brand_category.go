package models

import (
	"github.com/google/uuid"
)

type ProductBrandCategory struct {
	Id                uuid.UUID `gorm:"type:varchar(50)"`
	ProductBrandId    uuid.UUID `gorm:"type:varchar(50);not null"`
	ProductCategoryId uuid.UUID `gorm:"type:varchar(50);not null"`
	Slug              string    `gorm:"type:varchar(50);not null"`
	IsAvailable       bool      `gorm:"default:true;not null"`
	Products          []Product
}
