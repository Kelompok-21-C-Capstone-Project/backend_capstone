package request

import (
	"backend_capstone/services/product/dto"
	"log"
)

type PayloadUpdateStock struct {
	AdminId string               `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Date    string               `json:"date" example:"27 Jun 2022"`
	Datas   []UpdateProductStock `json:"datas"`
}

type UpdateProductStock struct {
	Id    string `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Stock uint32 `json:"add" example:"3"`
}

func (data *PayloadUpdateStock) DtoReq() *dto.UpdateStockDTO {
	dtoData := new(dto.UpdateStockDTO)
	dtoData.AdminId = data.AdminId
	log.Print(*data)
	for index := range data.Datas {
		dtoData.Datas = append(dtoData.Datas, dto.UpdateProductStockDTO{
			Id:    data.Datas[index].Id,
			Stock: data.Datas[index].Stock,
		})
	}
	return dtoData
}

type UpdateProductRequest struct {
	Name        string `json:"name" example:"Some Product"`
	Price       uint32 `json:"harga_jual" example:"30000"`
	PriceBuy    uint32 `json:"harga_beli" example:"30000"`
	Profit      uint32 `json:"komisi" example:"30000"`
	Stock       uint32 `json:"stock" example:"20"`
	IsAvailable bool   `json:"is_available"  example:"true"`
	IsDiscount  bool   `json:"is_discount" example:"true"`
	Description string `json:"description" example:"some description"`
	BrandId     string `json:"brand_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	CategoryId  string `json:"category_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
}

func (data *UpdateProductRequest) DtoReq() *dto.UpdateProductDTO {
	return &dto.UpdateProductDTO{
		Name:        data.Name,
		Price:       data.Price,
		PriceBuy:    data.PriceBuy,
		Profit:      data.Profit,
		Stock:       data.Stock,
		IsAvailable: data.IsAvailable,
		IsDiscount:  data.IsDiscount,
		Description: data.Description,
		BrandId:     data.BrandId,
		CategoryId:  data.CategoryId,
	}
}
