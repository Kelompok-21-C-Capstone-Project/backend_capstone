package midtransdriver

import (
	"backend_capstone/configs"
	"backend_capstone/models"
	"backend_capstone/utils/midtransdriver/dto"
	"log"
	"reflect"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransDriver struct {
	ServerKey         string
	Env               midtrans.EnvironmentType
	MidtransOperation coreapi.Client
}

func NewMidtransService(configs *configs.AppConfig) *MidtransDriver {
	log.Print("Enter NewMidtransService")
	var api MidtransDriver

	// 1. Set you ServerKey with globally
	midtrans.ServerKey = configs.API_Midtrans.SERVER_KEY
	midtrans.Environment = midtrans.Sandbox

	c := coreapi.Client{}
	c.New(configs.API_Midtrans.SERVER_KEY, midtrans.Sandbox)

	switch configs.API_Midtrans.ENV {
	case "sandbox":
		api.ServerKey = configs.API_Midtrans.SERVER_KEY
		api.Env = midtrans.Sandbox
		api.MidtransOperation = c
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
			Bank: midtrans.Bank(strings.ToLower(midtranspaymentDTO.MethodDetails)),
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
		PaymentType: coreapi.PaymentTypeShopeepay,
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
		PaymentType: coreapi.PaymentTypeGopay,
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
		PaymentType: coreapi.PaymentTypeQris,
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

func (d *MidtransDriver) DoPayment(method string, midtranspaymentDTO dto.MidtransPaymentDTO) (data *models.Payment) {
	log.Print("Enter midtransdriver.DoPayment")

	chargeReq := new(coreapi.ChargeReq)
	coreApiRes := new(coreapi.ChargeResponse)
	var va_number string

	// 2. Initiate charge request
	switch strings.ToLower(method) {
	case "virtual account":
		chargeReq = d.CreateBankTransferPayment(midtranspaymentDTO)
		coreApiRes, _ = d.MidtransOperation.ChargeTransaction(chargeReq)
		if midtranspaymentDTO.MethodDetails == "mandiri" || midtranspaymentDTO.MethodDetails == "permata" {
			va_number = coreApiRes.PermataVaNumber
		} else {
			va_number = reflect.ValueOf(coreApiRes.VaNumbers).Index(0).FieldByName("VANumber").Interface().(string)
		}
		log.Print("Response :", coreApiRes)
	case "gopay":
		chargeReq = d.CreateGopayPayment(midtranspaymentDTO)
		coreApiRes, _ = d.MidtransOperation.ChargeTransaction(chargeReq)
		va_number = reflect.ValueOf(coreApiRes.Actions).Index(0).FieldByName("URL").Interface().(string)
		log.Print("Response :", coreApiRes)
	case "shopeepay":
		chargeReq = d.CreateShopeePayPayment(midtranspaymentDTO)
		coreApiRes, _ = d.MidtransOperation.ChargeTransaction(chargeReq)
		log.Print("Response :", coreApiRes)
	case "qris":
		chargeReq = d.CreateQrisPayment(midtranspaymentDTO)
		coreApiRes, _ = d.MidtransOperation.ChargeTransaction(chargeReq)
		log.Print("Response :", coreApiRes)
	}
	data = &models.Payment{
		TransactionId: midtranspaymentDTO.OrderId,
		Amount:        uint32(midtranspaymentDTO.Paid),
		Method:        method,
		MethodDetails: midtranspaymentDTO.MethodDetails,
		Status:        coreApiRes.TransactionStatus,
		Description:   va_number,
	}
	log.Print(va_number)
	log.Print(data)
	return nil
}
