package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id                     string               `gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx" json:"id,omitempty"`
	Name                   string               `gorm:"type:varchar(50);not null" example:"Some Product" json:"name,omitempty"`
	Price                  uint32               `gorm:"default:0;not null" example:"30000" json:"price,omitempty"`
	Stock                  uint32               `gorm:"default:0;not null" example:"30" json:"stock,omitempty"`
	Discount               float32              `gorm:"default:1;type:DECIMAL(6,5)" json:"discount,omitempty"`
	IsAvailable            bool                 `gorm:"default:true;not null" example:"true" json:"is_available"`
	IsDiscount             bool                 `gorm:"default:false;not null" example:"false" json:"is_discount,omitempty"`
	Description            string               `gorm:"type:text" example:"some description" json:"description,omitempty"`
	Transactions           []Transaction        `json:"transactions,omitempty"`
	ProductBrandCategoryId string               `json:"category_id,omitempty"`
	ProductBrandCategory   ProductBrandCategory `json:"category,omitempty"`
	UpdatedAt              time.Time            `json:"updated_at,omitempty"`
	Deleted                gorm.DeletedAt       `json:"deleted,omitempty"`
}

type ProductResponse struct {
	Id                     string                       `gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx" json:"id,omitempty"`
	Name                   string                       `gorm:"type:varchar(50);not null" example:"Some Product" json:"name,omitempty"`
	Price                  uint32                       `gorm:"default:0;not null" example:"30000" json:"price,omitempty"`
	Stock                  uint32                       `gorm:"default:0;not null" example:"30" json:"stock,omitempty"`
	Discount               float32                      `gorm:"default:1;type:DECIMAL(6,5)" json:"discount,omitempty"`
	IsAvailable            bool                         `gorm:"default:true;not null" example:"true" json:"is_available"`
	IsDiscount             bool                         `gorm:"default:false;not null" example:"false" json:"is_discount,omitempty"`
	Description            string                       `gorm:"type:text" example:"some description" json:"description,omitempty"`
	Transactions           []TransactionResponse        `json:"transactions,omitempty"`
	ProductBrandCategoryId string                       `json:"category_id,omitempty"`
	ProductBrandCategory   ProductBrandCategoryResponse `json:"category,omitempty"`
	UpdatedAt              time.Time                    `json:"updated_at,omitempty"`
	Deleted                time.Time                    `json:"deleted,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}
