package models

import (
	"time"
)

type Payment struct {
	Id            string `gorm:"type:varchar(50)"`
	TransactionId string `gorm:"type:varchar(50)"`
	Amount        uint32 `gorm:"default:0;not null"`
	Billed        uint32 `gorm:"default:0;not null"`
	Charged       uint32 `gorm:"default:0;not null"`
	Method        string `grom:"type:text"`
	MethodDetails string `grom:"type:text"`
	Description   string `grom:"type:text"`
	Status        string `gorm:"type:varchar(30);default:'pending'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PaymentResponse struct {
	Id                    string `gorm:"type:varchar(50)"`
	TransactionResponseId string `gorm:"type:varchar(50)"`
	Amount                uint32 `gorm:"default:0;not null"`
	Billed                uint32 `gorm:"default:0;not null"`
	Charged               uint32 `gorm:"default:0;not null"`
	Method                string `grom:"type:text"`
	MethodDetails         string `grom:"type:text"`
	Description           string `grom:"type:text"`
	Status                string `gorm:"type:varchar(30);default:'pending'"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (PaymentResponse) TableName() string {
	return "payments"
}
