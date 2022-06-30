package dto

type ProductCategory struct {
	Id       string
	Name     string           `json:"label,omitempty"`
	Slug     string           `json:"slug,omitempty"`
	Products []*ClientProduct `json:"products,ommitempty"`
}

type ClientProduct struct {
	Id                string `json:"id,omitempty"`
	Group             string `json:"group,omitempty"`
	Label             string `json:"label,omitempty"`
	Description       string `json:"description,omitempty"`
	Stock             uint32 `json:"stock,omitempty"`
	Price             uint32 `json:"price,omitempty"`
	IsDiscount        bool   `json:"is_discount"`
	ProductCategoryId string `json:"category_id,omitempty"`
}
