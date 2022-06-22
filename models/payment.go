package models

import "github.com/google/uuid"

type Payment struct {
	Id              uuid.UUID `gorm:"type:varchar(50)"`
	Name            string    `gorm:"type:varchar(30);not null"`
	BodyRequest     string    `gorm:"type:text;not null"`
	PaymentMethodId uuid.UUID `gorm:"type:varchar(50)"`
	PaymentMethod   PaymentMethod
}
