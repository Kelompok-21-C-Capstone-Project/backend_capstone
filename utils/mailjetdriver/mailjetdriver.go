package mailjetdriver

import (
	"backend_capstone/configs"
	"backend_capstone/services/transaction"
)

type MailjetDriver struct {
	PrivateKey string
	PublicKey  string
}

func NewTransactionMailjetService(configs *configs.AppConfig) transaction.Mailjet {
	var api *MailjetDriver
	return api
}

func (d *MailjetDriver) SendBill() (err error) {
	return
}

func (d *MailjetDriver) SendInvoice() (err error) {
	return
}
