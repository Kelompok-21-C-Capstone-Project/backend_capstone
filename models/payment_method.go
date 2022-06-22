package models

import "github.com/google/uuid"

type PaymentMethod struct {
	Id   uuid.UUID `gorm:"type:varchar(50)"`
	Name string    `gorm:"type:varchar(50);not null"`
}
