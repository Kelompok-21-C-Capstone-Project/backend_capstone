package models

import (
	"github.com/google/uuid"
)

type BrandCategory struct {
	Id          uuid.UUID `gorm:"type:varchar(50)"`
	BrandId     uuid.UUID `gorm:"type:varchar(50);not null"`
	CategoryId  uuid.UUID `gorm:"type:varchar(50);not null"`
	Slug        string    `gorm:"type:varchar(50);not null"`
	IsAvailable bool      `gorm:"default:true;not null"`
}
