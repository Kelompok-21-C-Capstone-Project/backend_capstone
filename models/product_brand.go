package models

import "gorm.io/gorm"

type ProductBrand struct {
	Id                string            `json:"id" gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name              string            `json:"name" gorm:"type:varchar(50);not null" example:"some name"`
	IsAvailable       *bool             `json:"is_available" gorm:"default:true;not null" example:"false"`
	Description       string            `json:"description" gorm:"type:text" example:"some description"`
	IconPath          string            `gorm:"type:text" example:"/path/to/file" json:"icon_path,omitempty"`
	ProductCategories []ProductCategory `gorm:"many2many:brand_categories" json:"product_category,omitempty"`
	Deleted           gorm.DeletedAt    `json:"deleted,omitempty"`
}

type ProductBrandResponse struct {
	Id                string                    `json:"id" gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name              string                    `json:"name" gorm:"type:varchar(50);not null" example:"some name"`
	IsAvailable       *bool                     `json:"is_available" gorm:"default:true;not null" example:"false"`
	Description       string                    `json:"description" gorm:"type:text" example:"some description"`
	IconPath          string                    `gorm:"type:text" example:"/path/to/file" json:"icon_path,omitempty"`
	ProductCategories []ProductCategoryResponse `gorm:"many2many:brand_categories" json:"product_category"`
	// Deleted           time.Time                 `json:"deleted,omitempty"`
}

func (ProductBrandResponse) TableName() string {
	return "product_brands"
}
