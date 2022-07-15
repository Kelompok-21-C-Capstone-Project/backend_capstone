package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	dtoMidtrans "backend_capstone/utils/midtransdriver/dto"
	"errors"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Mailjet interface {
	SendBill(name string, email string, bill dto.BillClient) (err error)
	SendInvoice(name string, email string, bill dto.BillClient) (err error)
}

type Midtrans interface {
	DoPayment(method string, midtranspaymentDTO dtoMidtrans.MidtransPaymentDTO) (data *models.Payment, err error)
}

type Repository interface {
	FindById(id string) (transaction *models.TransactionResponse, err error)
	FindAll() (transactions *[]dto.ClientTransactionsResponse, err error)
	UsersFindAll(uip string, params ...string) (dataCount int64, transactions *[]dto.ClientTransactionsResponse, err error)
	UsersFindById(uid string, tid string) (transaction *dto.ClientTransactionsResponse, err error)
	CheckProductStock(pid string) (product *models.Product, err error)
	ProductReStock(pid string) (err error)
	Insert(data *models.Transaction) (transaction *models.TransactionResponse, err error)
	InsertPayment(data *models.Payment) (transaction *models.Payment, err error)
	Update() (transaction *models.TransactionResponse, err error)
	MidtransUpdate(tid string, status string) (err error)
	GetTransactionProduct(pid string) (product *models.Product, err error)
	GetBillById(tid string) (bill dto.BillClient, err error)
	GetUserInfo(tid string) (user models.UserResponse, err error)
	Delete(id string) (err error)
	AdminDetailTransaction(params ...string) (dashboardData dto.DetailPenjualanDTO, err error)
}

type Service interface {
	GetById(id string) (transaction models.TransactionResponse, err error)
	GetAll() (transactions []dto.ClientTransactionsResponse, err error)
	UsersGetAll(uid string, params ...string) (transactions dto.ResponseUserTransaction, err error)
	UsersGetById(uid string, tid string) (transaction dto.ClientTransactionsResponse, err error)
	Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (bill dto.BillClient, err error)
	GetBill(uid string, tid string) (bills dto.BillClient, err error)
	Modify() (transaction models.TransactionResponse, err error)
	Remove() (err error)
	MidtransAfterPayment(midtransData dto.MidtransAfterPayment) (err error)
	AdminDetailTransaction(params ...string) (dashboardData dto.DetailPenjualanDTO, err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
	midtrans   Midtrans
	mailjet    Mailjet
}

func NewService(repository Repository, midtransApi Midtrans, mailjetApi Mailjet) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
		midtrans:   midtransApi,
		mailjet:    mailjetApi,
	}
}

func (s *service) GetById(id string) (transaction models.TransactionResponse, err error) {

	return
}
func (s *service) GetAll() (transactions []dto.ClientTransactionsResponse, err error) {
	data, err := s.repository.FindAll()
	for in := range *data {
		(*data)[in].CreatedAt = (*data)[in].CreatedAt.Round(60 * time.Minute)
	}
	if err != nil {
		return
	}
	if data == nil {
		transactions = []dto.ClientTransactionsResponse{}
		return
	}
	transactions = *data
	return
}
func (s *service) UsersGetAll(uid string, params ...string) (transactions dto.ResponseUserTransaction, err error) {
	if params[2] != "" {
		regexDateRange := "([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])_([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])"
		if resDR, _ := regexp.MatchString(regexDateRange, params[2]); !resDR {
			return
		}
		date := strings.Split(params[2], "_")
		dateTop, _ := time.Parse("02-01-2006 15:04:05", date[1]+" 08:04:00")
		date[1] = dateTop.AddDate(0, 0, 1).Format("02-01-2006")
		params = append(params, date...)
	} else if params[1] != "" {
		regexDate := "([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])"
		if resR, _ := regexp.MatchString(regexDate, params[1]); !resR {
			return
		}
	}
	for in, el := range params {
		if in == 0 || in == 3 || in == 4 {
			params[in] = "%" + el + "%"
		}
	}
	if params[5] == "" {
		params[5] = "1"
	}
	if params[6] == "" {
		params[6] = "5"
	}
	den, err := strconv.Atoi(params[6])
	if err != nil {
		return
	}
	count, data, err := s.repository.UsersFindAll(uid, params...)
	if err != nil {
		return
	}
	if den < -1 || den == 0 {
		den = 10
	} else if den == 0 {
		den = int(count)
	}
	if data == nil {
		transactions.Data = []dto.ClientTransactionsResponse{}
		return
	}
	transactions.Count = int(math.Ceil(float64(count) / float64(den)))
	transactions.Data = *data
	return
}
func (s *service) UsersGetById(uid string, tid string) (transaction dto.ClientTransactionsResponse, err error) {
	_, err = uuid.Parse(uid)
	if err != nil {
		return
	}
	_, err = uuid.Parse(tid)
	if err != nil {
		return
	}
	data, err := s.repository.UsersFindById(uid, tid)
	if err != nil {
		return
	}
	if data == nil {
		transaction = dto.ClientTransactionsResponse{}
		return
	}
	transaction = *data
	return
}
func (s *service) Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (bill dto.BillClient, err error) {
	if err = s.validate.Struct(createtransactionDTO); err != nil {
		return
	}
	// check stock barang
	dataProduct, err := s.repository.CheckProductStock(createtransactionDTO.ProductId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(userId)
	if err != nil {
		return
	}
	tid := uuid.New().String()
	dataTransaction, err := s.repository.Insert(createtransactionDTO.GenerateTransactionModel(tid, userId))
	if err != nil {
		return
	}
	payment, err := s.midtrans.DoPayment(createtransactionDTO.PaymentMethod, createtransactionDTO.GenerateMidtransPayment(tid, dataProduct.PriceBuy))
	if err != nil {
		s.repository.ProductReStock(createtransactionDTO.ProductId)
		s.repository.Delete(tid)
		// error transaksi gagal
		err = errors.New("Transaksi gagal")
		return
	}
	pid := uuid.New().String()
	payment.Id = pid
	dataPayment, err := s.repository.InsertPayment(payment)
	if err != nil {
		return
	}
	// send bill using mail jet

	bill = dto.BillClient{
		Id:             dataPayment.Id,
		TransactionId:  dataTransaction.Id,
		Status:         dataPayment.Status,
		VaNumber:       dataPayment.Description,
		PaymentDetails: dataPayment.MethodDetails,
		Billed:         dataPayment.Billed,
		Product:        dataProduct.Name,
		ProductPrice:   dataProduct.Price,
		Charger:        dataProduct.Profit,
		Deadline:       dataPayment.CreatedAt.Add(time.Hour * time.Duration(1)),
	}
	userInfo, err := s.repository.GetUserInfo(tid)
	if err != nil {
		return
	}
	err = s.mailjet.SendBill(userInfo.Name, createtransactionDTO.Email, bill)
	if err != nil {
		return
	}
	return
}
func (s *service) GetBill(uid string, tid string) (bills dto.BillClient, err error) {
	_, err = uuid.Parse(tid)
	if err != nil {
		return
	}
	// ngambil data transaksi dari repo pake tid : transaction
	dataTransaction, err := s.repository.FindById(tid)
	if err != nil {
		return
	}
	log.Print(dataTransaction)
	// ngambil data product dari repo pake data `transaction.product_id` : product
	dataProduct, err := s.repository.GetTransactionProduct(dataTransaction.ProductId)
	if err != nil {
		return
	}
	bills = dto.BillClient{
		Id:             dataTransaction.Payment.Id,
		TransactionId:  dataTransaction.Id,
		Status:         dataTransaction.Payment.Status,
		VaNumber:       dataTransaction.Payment.Description,
		PaymentDetails: dataTransaction.Payment.MethodDetails,
		Billed:         dataTransaction.Payment.Billed,
		Product:        dataProduct.Name,
		ProductPrice:   dataProduct.Price,
		Charger:        dataProduct.Profit,
		Deadline:       dataTransaction.Payment.CreatedAt.Add(time.Hour * time.Duration(1)),
	}
	return
}
func (s *service) MidtransAfterPayment(midtransData dto.MidtransAfterPayment) (err error) {
	_, err = uuid.Parse(midtransData.TransactionId)
	if err != nil {
		err = errors.New("Transaction Id " + midtransData.TransactionId + " Is Invalid")
		return
	}
	if midtransData.Code == "201" {
		log.Print("Update transaction skipped")
		err = errors.New("Update transaction skipped")
		return
	}
	switch midtransData.Status {
	case "capture":
		midtransData.Status = "Success"
	case "settlement":
		midtransData.Status = "Success"
	default:
		midtransData.Status = "Cancelled"
	}
	if err = s.repository.MidtransUpdate(midtransData.TransactionId, midtransData.Status); err != nil {
		err = errors.New("Midtrans Transaction Id " + midtransData.TransactionId + " Fail To Update")
	}
	bills, err := s.repository.GetBillById(midtransData.TransactionId)
	if err != nil {
		return
	}
	userInfo, err := s.repository.GetUserInfo(midtransData.TransactionId)
	if err != nil {
		return
	}
	err = s.mailjet.SendInvoice(userInfo.Name, userInfo.Email, bills)
	if err != nil {
		return
	}
	return
}
func (s *service) Modify() (transaction models.TransactionResponse, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
func (s *service) AdminDetailTransaction(params ...string) (dashboardData dto.DetailPenjualanDTO, err error) {
	if params[1] != "" {
		regexDateRange := "([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])_([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])"
		if resDR, _ := regexp.MatchString(regexDateRange, params[2]); !resDR {
			return
		}
		date := strings.Split(params[1], "_")
		dateTop, _ := time.Parse("02-01-2006 15:04:05", date[1]+" 08:04:00")
		date[1] = dateTop.AddDate(0, 0, 1).Format("02-01-2006")
		params = append(params, date...)
	}
	dashboardData, err = s.repository.AdminDetailTransaction(params...)
	if err != nil {
		return
	}
	return
}
