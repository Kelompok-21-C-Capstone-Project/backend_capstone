package midtransdriver

import (
	"backend_capstone/configs"
	"backend_capstone/utils/midtransdriver/dto"
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransDriver struct {
	ServerKey string
	Env       midtrans.EnvironmentType
}

func NewMidtransService(configs *configs.AppConfig) *MidtransDriver {
	log.Print("Enter NewMidtransService")
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
func (d *MidtransDriver) CreateBankTransferPayment(midtranspaymentDTO dto.MidtransPaymentDTO) *coreapi.ChargeReq {
	log.Print("Enter midtransdriver.CreateBankTransferTransaction")
	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  midtranspaymentDTO.OrderId,
			GrossAmt: midtranspaymentDTO.Paid,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(midtranspaymentDTO.MethodDetails),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    midtranspaymentDTO.ItemId,
				Price: midtranspaymentDTO.ItemPrice,
				Qty:   1,
				Name:  midtranspaymentDTO.ItemName,
			},
		},
	}
}

// Create ChargeReq for ShopeePay
func (d *MidtransDriver) CreateShopeePayPayment(midtranspaymentDTO dto.MidtransPaymentDTO) *coreapi.ChargeReq {
	log.Print("Enter midtransdriver.CreateShopeePayTransaction")
	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  midtranspaymentDTO.OrderId,
			GrossAmt: midtranspaymentDTO.Paid,
		},
		ShopeePay: &coreapi.ShopeePayDetails{},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    midtranspaymentDTO.ItemId,
				Price: midtranspaymentDTO.ItemPrice,
				Qty:   1,
				Name:  midtranspaymentDTO.ItemName,
			},
		},
	}
}

// Create ChargeReq for Gopay
func (d *MidtransDriver) CreateGopayPayment(midtranspaymentDTO dto.MidtransPaymentDTO) *coreapi.ChargeReq {
	log.Print("Enter midtransdriver.CreateGopayTransaction")
	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  midtranspaymentDTO.OrderId,
			GrossAmt: midtranspaymentDTO.Paid,
		},
		Gopay: &coreapi.GopayDetails{},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    midtranspaymentDTO.ItemId,
				Price: midtranspaymentDTO.ItemPrice,
				Qty:   1,
				Name:  midtranspaymentDTO.ItemName,
			},
		},
	}
}

// Create ChargeReq for Qris
func (d *MidtransDriver) CreateQrisPayment(midtranspaymentDTO dto.MidtransPaymentDTO) *coreapi.ChargeReq {
	log.Print("Enter midtransdriver.CreateQrisTransaction")
	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  midtranspaymentDTO.OrderId,
			GrossAmt: midtranspaymentDTO.Paid,
		},
		Qris: &coreapi.QrisDetails{},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    midtranspaymentDTO.ItemId,
				Price: midtranspaymentDTO.ItemPrice,
				Qty:   1,
				Name:  midtranspaymentDTO.ItemName,
			},
		},
	}
}

func (d *MidtransDriver) DoPayment(method string, midtranspaymentDTO dto.MidtransPaymentDTO) interface{} {
	log.Print("Enter midtransdriver.DoPayment")

	chargeReq := new(coreapi.ChargeReq)

	// 2. Initiate charge request
	switch method {
	case "bank_transfer":
		chargeReq = d.CreateBankTransferPayment(midtranspaymentDTO)
	case "gopay":
		chargeReq = d.CreateGopayPayment(midtranspaymentDTO)
	case "shopeepay":
		chargeReq = d.CreateShopeePayPayment(midtranspaymentDTO)
	case "qris":
		chargeReq = d.CreateQrisPayment(midtranspaymentDTO)
	}

	// 3. Request to Midtrans using global config
	coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
	log.Print("Response :", coreApiRes)

	return nil
}
