package dto

type ResponseBodyProductBrand struct {
	PageLength int            `json:"page_length" example:"3"`
	Data       []ProductBrand `json:"datas"`
}

type ProductBrand struct {
	Id          string `json:"id,omitempty" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Name        string `json:"name,omitempty" example:"some name"`
	Category    string `json:"category,omitempty" example:"E-Money"`
	IsAvailable bool   `json:"is_available" example:"false"`
	Description string `json:"description,omitempty" example:"some description"`
	IconPath    string `json:"icon_path,omitempty" example:"mdi-some-brand"`
	Product     int    `json:"product_sum" example:"33"`
}

func (ProductBrand) TableName() string {
	return "product_brands"
}
