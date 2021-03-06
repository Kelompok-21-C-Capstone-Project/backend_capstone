package models

import (
	"gorm.io/gorm"
)

type ProductBrandCategory struct {
	Id                string         `gorm:"type:varchar(50)" json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	ProductBrandId    string         `gorm:"type:varchar(50);not null" json:"brand_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	ProductCategoryId string         `gorm:"type:varchar(50);not null" json:"category_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Slug              string         `gorm:"type:varchar(50);not null" json:"slug,omitempty"`
	IsAvailable       *bool          `gorm:"default:true;not null" json:"is_available" example:"true"`
	Products          []Product      `json:"products,omitempty"`
	Deleted           gorm.DeletedAt `json:"deleted,omitempty"`
}

type ProductBrandCategoryResponse struct {
	Id                string            `gorm:"type:varchar(50)" json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	ProductBrandId    string            `gorm:"type:varchar(50);not null" json:"brand_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	ProductCategoryId string            `gorm:"type:varchar(50);not null" json:"category_id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Slug              string            `gorm:"type:varchar(50);not null" json:"slug,omitempty"`
	IsAvailable       *bool             `gorm:"default:true;not null" json:"is_available" example:"true"`
	Products          []ProductResponse `gorm:"foreignKey:ProductBrandCategoryId" json:"products"`
	// Deleted           time.Time         `json:"deleted,omitempty"`
}

func (ProductBrandCategoryResponse) TableName() string {
	return "product_brand_categories"
}
