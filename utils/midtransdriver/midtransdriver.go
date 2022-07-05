package midtransdriver

import (
	"backend_capstone/configs"
	"backend_capstone/models"
	"backend_capstone/utils/midtransdriver/dto"
	"errors"
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
	bank := strings.ToLower(midtranspaymentDTO.MethodDetails)
	switch {
	case strings.Contains(bank, "bni"):
		bank = "bni"
	case strings.Contains(bank, "mandiri"):
		bank = "mandiri"
	case strings.Contains(bank, "cimb"):
		bank = "cimb"
	case strings.Contains(bank, "bca"):
		bank = "bca"
	case strings.Contains(bank, "bri"):
		bank = "bri"
	case strings.Contains(bank, "maybank"):
		bank = "maybank"
	case strings.Contains(bank, "permata"):
		bank = "permata"
	case strings.Contains(bank, "mega"):
		bank = "mega"
	}
	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  midtranspaymentDTO.OrderId,
			GrossAmt: midtranspaymentDTO.Paid,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(bank),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    midtranspaymentDTO.ItemId,
				Price: midtranspaymentDTO.ItemPrice,
				Qty:   1,
				Name:  midtranspaymentDTO.ItemName,
			},
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 1,
			Unit:           "hour",
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
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 1,
			Unit:           "hour",
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
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 1,
			Unit:           "hour",
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
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 1,
			Unit:           "hour",
		},
	}
}

func (d *MidtransDriver) DoPayment(method string, midtranspaymentDTO dto.MidtransPaymentDTO) (data *models.Payment, err error) {
	log.Print("Enter midtransdriver.DoPayment")

	chargeReq := new(coreapi.ChargeReq)
	coreApiRes := new(coreapi.ChargeResponse)
	coreApiErr := new(midtrans.Error)
	var va_number string

	// 2. Initiate charge request
	switch strings.ToLower(method) {
	case "virtual account":
		chargeReq = d.CreateBankTransferPayment(midtranspaymentDTO)
		coreApiRes, coreApiErr = d.MidtransOperation.ChargeTransaction(chargeReq)
		if midtranspaymentDTO.MethodDetails == "mandiri" || midtranspaymentDTO.MethodDetails == "permata" {
			va_number = coreApiRes.PermataVaNumber
		} else {
			va_number = reflect.ValueOf(coreApiRes.VaNumbers).Index(0).FieldByName("VANumber").Interface().(string)
		}
		// log.Print("Response :", coreApiRes)
	case "gopay":
		chargeReq = d.CreateGopayPayment(midtranspaymentDTO)
		coreApiRes, coreApiErr = d.MidtransOperation.ChargeTransaction(chargeReq)
		va_number = reflect.ValueOf(coreApiRes.Actions).Index(0).FieldByName("URL").Interface().(string)
		// log.Print("Response :", coreApiRes)
	case "shopeepay":
		chargeReq = d.CreateShopeePayPayment(midtranspaymentDTO)
		coreApiRes, coreApiErr = d.MidtransOperation.ChargeTransaction(chargeReq)
		// log.Print("Response :", coreApiRes)
	case "qris":
		chargeReq = d.CreateQrisPayment(midtranspaymentDTO)
		coreApiRes, coreApiErr = d.MidtransOperation.ChargeTransaction(chargeReq)
		// log.Print("Response :", coreApiRes)
	}
	if coreApiErr != nil {
		log.Print("Error midtransdriver.DoPayment")
		err = errors.New("Error during payment")
		return
	}

	log.Print("Response :", coreApiRes)

	switch coreApiRes.TransactionStatus {
	case "pending":
		midtranspaymentDTO.MethodDetails = "Pending"
	case "deny":
		midtranspaymentDTO.MethodDetails = "Rejected"
	}

	data = &models.Payment{
		TransactionId: midtranspaymentDTO.OrderId,
		Amount:        uint32(midtranspaymentDTO.Paid),
		Billed:        uint32(midtranspaymentDTO.Paid) + 1500,
		Charged:       uint32(1500),
		Method:        method,
		MethodDetails: midtranspaymentDTO.MethodDetails,
		Status:        coreApiRes.TransactionStatus,
		Description:   va_number, //nomer hp atau nomer virtual account bank
	}
	return
}
