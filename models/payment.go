package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

type Payment struct {
	Id              uuid.UUID `gorm:"type:varchar(50)"`
	TransactionId   uuid.UUID `gorm:"type:varchar(50)"`
	Amount          uint32    `gorm:"default:0;not null"`
	Details         string    `grom:"type:text"`
	Description     string    `grom:"type:text"`
	Status          string    `gorm:"type:varchar(30);default:'mengunggu'"`
	PaymentVendorId uuid.UUID
	PaymentVendor   PaymentVendor
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
