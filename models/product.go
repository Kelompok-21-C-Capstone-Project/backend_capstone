package models

import (
	"time"
)

type Product struct {
	Id                     string  `gorm:"type:varchar(50);not null"`
	Name                   string  `gorm:"type:varchar(50);not null"`
	Price                  uint32  `gorm:"default:0;not null"`
	Stock                  uint32  `gorm:"default:0;not null"`
	Discount               float32 `gorm:"default:1;type:DECIMAL(6,5)"`
	IsAvailable            bool    `gorm:"default:true;not null"`
	IsDiscount             bool    `gorm:"default:false;not null"`
	Description            string  `gorm:"type:text"`
	Transactions           []Transaction
	ProductBrandCategoryId string
	ProductBrandCategory   ProductBrandCategory
	UpdatedAt              time.Time
}
