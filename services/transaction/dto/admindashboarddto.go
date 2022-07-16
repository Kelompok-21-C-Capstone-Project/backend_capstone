package dto

type DashboardDetailTransactionDTO struct {
	Sum          DetailPenjualanDTO `json:"summary"`
	Transactions DataTransactionDTO `json:"transaksi"`
}

type DetailPenjualanDTO struct {
	Sum              uint32 `json:"total_penjualan"`
	Profit           uint32 `json:"total_komisi"`
	TransactionCount uint32 `json:"total_transaksi"`
}

type DataTransactionDTO struct {
	PageLength int64                  `json:"page_length" example:"3"`
	Data       []DetailTransactionDTO `json:"datas"`
}

type DetailTransactionDTO struct {
	Id            string `json:"invoice" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Status        string `json:"status" example:"Pending"`
	OrderTime     string `json:"order_time" example:"01/02/2022 23:12"`
	PaymmentTime  string `json:"payment_time" example:"01/02/2022 23:12"`
	Category      string `json:"category" example:"Paket Data"`
	UserId        string `json:"user_id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	UserName      string `json:"user_name" example:"Hadi"`
	Paid          uint32 `json:"nominal" example:"35000"`
	PaymentMethod string `json:"payment_method" example:"BCA Virtual Account"`
}

func (DetailTransactionDTO) TableName() string {
	return "transactions"
}
