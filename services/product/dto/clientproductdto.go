package dto

type ProductCategory struct {
	Id       string           `json:"id,omitempty" example:"xxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name     string           `json:"label,omitempty" example:"Some Category"`
	Slug     string           `json:"slug,omitempty" example:"some-category"`
	Products []*ClientProduct `json:"products,ommitempty"`
}

type ClientProduct struct {
	Id                string `json:"id,omitempty" example:"xxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Group             string `json:"group,omitempty" example:"Telkomsel"`
	Label             string `json:"label,omitempty" example:"Telkomsel 5GB Super 24H"`
	Description       string `json:"description,omitempty" example:"Some description"`
	Stock             uint32 `json:"stock,omitempty" example:"30"`
	Price             uint32 `json:"price,omitempty" example:"50000"`
	IsDiscount        bool   `json:"is_discount" example:"false"`
	ProductCategoryId string `json:"category_id,omitempty" example:"xxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
}
