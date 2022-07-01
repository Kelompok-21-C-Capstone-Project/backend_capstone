package models

import (
	"time"
)

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

type User struct {
	Id           string `gorm:"type:varchar(50);not null"`
	Username     string `gorm:"type:varchar(50);unique;not null"`
	FirstName    string `gorm:"type:varchar(50);not null"`
	LastName     string `gorm:"type:varchar(50);not null"`
	Role         string `gorm:"type:varchar(30);default:'user';not null"`
	Email        string `gorm:"type:varchar(30);unique;not null"`
	Phone        string `gorm:"type:varchar(15);unique;not null"`
	Password     string `gorm:"type:varchar(30);not null"`
	IconPath     string `gorm:"type:text"`
	Transactions []Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
