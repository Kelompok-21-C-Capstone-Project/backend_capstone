package models

import (
	"time"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

type Payment struct {
	Id              string `gorm:"type:varchar(50)"`
	TransactionId   string `gorm:"type:varchar(50)"`
	Amount          uint32 `gorm:"default:0;not null"`
	Details         string `grom:"type:text"`
	Description     string `grom:"type:text"`
	Status          string `gorm:"type:varchar(30);default:'mengunggu'"`
	PaymentVendorId string
	PaymentVendor   PaymentVendor
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
