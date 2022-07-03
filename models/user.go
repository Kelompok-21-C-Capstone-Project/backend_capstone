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
	Name         string `gorm:"type:varchar(50);not null"`
	Role         string `gorm:"type:varchar(30);default:'user';not null"`
	Email        string `gorm:"type:varchar(30);unique;not null"`
	Phone        string `gorm:"type:varchar(15);unique;not null"`
	Password     string `gorm:"type:varchar(100);not null"`
	IconPath     string `gorm:"type:text"`
	Transactions []Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserResponse struct {
	Id       string `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" format:"uuid"`
	Username string `json:"username" example:"username"`
	Name     string `json:"name" example:"somen name"`
	Role     string `json:"role" example:"user"`
	Email    string `json:"email" example:"testmail@mail.com" format:"email"`
	Phone    string `json:"phone" example:"0811111111"`
}

func (UserResponse) TableName() string {
	return "users"
}
