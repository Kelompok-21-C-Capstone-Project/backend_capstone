package midtransdriver

import (
	"backend_capstone/configs"
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransDriver struct {
	ServerKey string
	Env       midtrans.EnvironmentType
}

func NewMidtransDriver(configs *configs.AppConfig) *MidtransDriver {
	var api MidtransDriver

	// 1. Set you ServerKey with globally
	midtrans.ServerKey = configs.API_Midtrans.SERVER_KEY
	midtrans.Environment = midtrans.Sandbox

	switch configs.API_Midtrans.ENV {
	case "sandbox":
		api.ServerKey = configs.API_Midtrans.SERVER_KEY
		api.Env = midtrans.Sandbox
	}
	return &api
}

func (d *MidtransDriver) GetPaymentMethod() interface{} {

	// 2. Initiate charge request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeCreditCard,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
		CreditCard: &coreapi.CreditCardDetails{
			TokenID:        "YOUR-CC-TOKEN",
			Authentication: true,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}

	// 3. Request to Midtrans using global config
	coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
	log.Print("Response :", coreApiRes)
	return false
}
