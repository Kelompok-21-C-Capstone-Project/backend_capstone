package dto

type ResponseBodyProductCategory struct {
	PageLength int               `json:"page_length" example:"3"`
	Data       []ProductCategory `json:"datas"`
}

type ProductCategory struct {
	Id          string `json:"id,omitempty" gorm:"type:varchar(50);not null" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name        string `json:"name,omitempty" gorm:"type:varchar(50);not null,unique" example:"Some Name"`
	Slug        string `json:"slug,omitempty" gorm:"type:varchar(50);not null" example:"some-name"`
	IsAvailable bool   `json:"is_available" gorm:"default:false;not null" example:"true"`
	Description string `json:"description,omitempty" gorm:"type:text" example:"some description"`
	Icon        string `json:"icon,omitempty" gorm:"type:text" example:"mdi-some-category"`
	Product     int    `json:"product_sum" example:"33"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
