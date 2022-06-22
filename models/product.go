package models

import "github.com/google/uuid"

type Product struct {
	Id              uuid.UUID `gorm:"type:varchar(50);not null"`
	Name            string    `gorm:"type:varchar(50);not null"`
	Price           uint32    `gorm:"default:0;not null"`
	Stock           uint32    `gorm:"default:0;not null"`
	Discount        float32   `gorm:"default:1;type:DECIMAL(6,5)"`
	IsAvailable     bool      `gorm:"default:true;not null"`
	IsDiscount      bool      `gorm:"default:false;not null"`
	Description     string    `gorm:"type:text"`
	BrandCategoryId uuid.UUID
	BrandCategory   BrandCategory
}
