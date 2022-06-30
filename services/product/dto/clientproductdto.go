package dto

import "github.com/google/uuid"

type ProductCategory struct {
	Id       uuid.UUID
	Name     string           `json:"label,omitempty"`
	Slug     string           `json:"slug,omitempty"`
	Products *[]ClientProduct `json:"products,ommitempty"`
}

type ClientProduct struct {
	Id                uuid.UUID `json:"id,omitempty"`
	Group             string    `json:"group,omitempty"`
	Label             string    `json:"label,omitempty"`
	Description       string    `json:"description,omitempty"`
	Stock             uint32    `json:"stock,omitempty"`
	Price             uint32    `json:"price,omitempty"`
	IsDiscount        bool      `json:"is_discount,omitempty"`
	ProductCategoryId uuid.UUID `json:"category_id,omitempty"`
}
