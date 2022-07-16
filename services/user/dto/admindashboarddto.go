package dto

type UserDashboadDTO struct {
	PageLength int64         `json:"page_length" example:"3"`
	Datas      []UserDetails `json:"datas"`
}

type UserDetails struct {
	Id           string `json:"id"`
	Name         string `json:"nama"`
	Phone        string `json:"telepon"`
	Email        string `json:"email"`
	Transactions int64  `json:"total_transaksi"`
	Spending     int64  `json:"total_pembelian"`
}

func (UserDetails) TableName() string {
	return "users"
}
