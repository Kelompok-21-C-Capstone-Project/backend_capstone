package models

import "time"

//> model tipe data
type User struct {
	UserId    int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated"`
}

func (User) TableName() string {
	return ("user")
}
