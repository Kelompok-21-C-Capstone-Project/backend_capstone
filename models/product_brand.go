package models

import "github.com/google/uuid"

type ProductBrand struct {
	Id              uuid.UUID         `gorm:"type:varchar(50);not null"`
	Name            string            `gorm:"type:varchar(50);not null"`
	IsAvailable     bool              `gorm:"default:true;not null"`
	Description     string            `gorm:"type:text"`
	IconPath        string            `gorm:"type:text"`
	ProductCategory []ProductCategory `gorm:"many2many:brand_categories"`
}
