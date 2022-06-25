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
	log.Print("Enter NewMidtransDriver")
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

func (d *MidtransDriver) PutApprovePaymentMethod() interface{} {
	approveRes, _ := coreapi.ApproveTransaction("54321")
	log.Print("Response :", approveRes)
	return nil
}

// Create ChargeReq for bank transfer
func (d *MidtransDriver) CreateBankTransferTransaction() interface{} {
	log.Print("Enter midtransdriver.CreateBankTransferTransaction")
	return nil
}

// Create ChargeReq for ShopeePay
func (d *MidtransDriver) CreateShopeePayTransaction() interface{} {
	log.Print("Enter midtransdriver.CreateShopeePayTransaction")
	return nil
}

// Create ChargeReq for Gopay
func (d *MidtransDriver) CreateGopayTransaction() interface{} {
	log.Print("Enter midtransdriver.CreateGopayTransaction")
	return nil
}

// Create ChargeReq for ConvStore
func (d *MidtransDriver) CreateConvStoreTransaction() interface{} {
	log.Print("Enter midtransdriver.CreateConvStoreTransaction")
	return nil
}

// Create ChargeReq for Qris
func (d *MidtransDriver) CreateQrisTransaction() interface{} {
	log.Print("Enter midtransdriver.CreateQrisTransaction")
	return nil
}

func (d *MidtransDriver) GetPaymentMethod() interface{} {
	log.Print("Enter midtransdriver.GetPaymentMethod")

	// 2. Initiate charge request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "54321",
			GrossAmt: 200000,
		},
		CreditCard: &coreapi.CreditCardDetails{
			TokenID:        "YOUR-CC-TOKEN",
			Authentication: true,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM2",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}

	// 3. Request to Midtrans using global config
	coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
	log.Print("Response :", coreApiRes)

	return nil
}
