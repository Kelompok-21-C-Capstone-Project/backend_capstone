package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id                     string               `gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx" json:"id,omitempty"`
	Name                   string               `gorm:"type:varchar(50);not null" example:"Some Product" json:"name,omitempty"`
	Price                  uint32               `gorm:"default:0;not null" example:"30000" json:"harga_jual,omitempty"`
	PriceBuy               uint32               `gorm:"default:0;not null" example:"30000" json:"harga_beli,omitempty"`
	Profit                 uint32               `gorm:"default:0;not null" example:"30000" json:"komisi,omitempty"`
	Stock                  uint32               `gorm:"default:0;not null" example:"30" json:"stock,omitempty"`
	Discount               float32              `gorm:"default:1;type:DECIMAL(6,5)" json:"discount,omitempty"`
	IsAvailable            *bool                `gorm:"default:true;not null" example:"true" json:"is_available"`
	IsDiscount             *bool                `gorm:"default:false;not null" example:"false" json:"is_discount,omitempty"`
	Description            string               `gorm:"type:text" example:"some description" json:"description,omitempty"`
	Transactions           []Transaction        `json:"transactions,omitempty"`
	ProductBrandCategoryId string               `json:"category_id,omitempty"`
	ProductBrandCategory   ProductBrandCategory `json:"category,omitempty"`
	UpdatedAt              time.Time            `json:"updated_at,omitempty"`
	Deleted                gorm.DeletedAt       `json:"deleted,omitempty"`
}

type Supply struct {
	Id        string    `gorm:"type:varchar(50);not null"`
	Name      string    `gorm:"type:text;not null"`
	CodeNo    string    `gorm:"type:varchar(50);not null"`
	SumStock  uint32    `gorm:"default:0;not null"`
	Products  []Product `gorm:"many2many:supply_products"`
	AdminId   string    `gorm:"type:varchar(50);not null"`
	Admin     User
	CreatedAt time.Time
}

type SupplyProduct struct {
	Id        uint32
	SupplyId  string `gorm:"type:varchar(50);not null"`
	ProductId string `gorm:"type:varchar(50);not null"`
	Amount    uint32 `gorm:"default:0;not null"`
	Cost      uint32 `gorm:"default:0;not null"`
}

type ProductResponse struct {
	Id                     string                       `json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"`
	Name                   string                       `json:"name,omitempty" example:"Some Product"`
	Price                  uint32                       `json:"harga_jual,omitempty" example:"30000"`
	PriceBuy               uint32                       `json:"harga_beli,omitempty" example:"30000"`
	Profit                 uint32                       `json:"komisi,omitempty" example:"30000"`
	Stock                  uint32                       `json:"stock,omitempty" example:"30"`
	Discount               float32                      `json:"discount,omitempty"`
	IsAvailable            *bool                        `json:"is_available" example:"true"`
	IsDiscount             *bool                        `json:"is_discount,omitempty" example:"false"`
	Description            string                       `json:"description,omitempty" example:"some description"`
	Transactions           []TransactionResponse        `json:"transactions" gorm:"foreignKey:ProductId"`
	ProductBrandCategoryId string                       `json:"category_id,omitempty"`
	ProductBrandCategory   ProductBrandCategoryResponse `json:"category,omitempty"`
	UpdatedAt              time.Time                    `json:"updated_at,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}
