package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: PENETAPAN ATRIBUT ENTITAS

type Payment struct {
	Id              uuid.UUID `gorm:"type:varchar(50)"`
	TransactionId   uuid.UUID `gorm:"type:varchar(50)"`
	Details         string    `grom:"type:text"`
	PaymentVendorId uuid.UUID
	PaymentVendor   PaymentVendor
	CreatedAt       time.Time
	UpdatedAt       time.Time
}