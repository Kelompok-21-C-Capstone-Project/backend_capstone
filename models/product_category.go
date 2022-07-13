package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	Id          string         `json:"id,omitempty" gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name        string         `json:"name,omitempty" gorm:"type:varchar(50);not null,unique" example:"Some Name"`
	Slug        string         `json:"slug,omitempty" gorm:"type:varchar(50);not null" example:"some-name"`
	IsAvailable bool           `json:"is_available" gorm:"default:false;not null" example:"true"`
	Description string         `json:"description,omitempty" gorm:"type:text" example:"some description"`
	Icon        string         `json:"icon,omitempty" gorm:"type:text" example:"mdi-some-category"`
	Deleted     gorm.DeletedAt `json:"deleted,omitempty"`
}

type ProductCategoryResponse struct {
	Id          string    `json:"id,omitempty" gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name        string    `json:"name,omitempty" gorm:"type:varchar(50);not null,unique" example:"Some Name"`
	Slug        string    `json:"slug,omitempty" gorm:"type:varchar(50);not null" example:"some-name"`
	IsAvailable bool      `json:"is_available" gorm:"default:false;not null" example:"true"`
	Description string    `json:"description,omitempty" gorm:"type:text" example:"some description"`
	Icon        string    `json:"icon,omitempty" gorm:"type:text" example:"mdi-some-category"`
	Deleted     time.Time `json:"deleted,omitempty"`
}

func (ProductCategoryResponse) TableName() string {
	return "product_categories"
}
