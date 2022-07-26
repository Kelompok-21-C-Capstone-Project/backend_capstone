package dto

type DataSupplyDTO struct {
	PageLength int64              `json:"page_length" example:"3"`
	Data       []SupplyProductDTO `json:"datas"`
}

type SupplyProductDTO struct {
	Id   string `json:"no_faktur" example:"INV/ADSA/ASD"`
	Name string `json:"product_data" example:"Paket, PLN, Pulsa"`
	Sum  uint32 `json:"jumlah" example:"30"`
	Date string `json:"tanggal" example:"12-02-2022"`
}

func (SupplyProductDTO) TableName() string {
	return "supplies"
}
