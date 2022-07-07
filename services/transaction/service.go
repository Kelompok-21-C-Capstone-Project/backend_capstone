package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	dtoMidtrans "backend_capstone/utils/midtransdriver/dto"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Mailjet interface {
	SendBill() (err error)
	SendInvoice() (err error)
}

type Midtrans interface {
	DoPayment(method string, midtranspaymentDTO dtoMidtrans.MidtransPaymentDTO) (data *models.Payment, err error)
}

type Repository interface {
	FindById(id string) (transaction *models.Transaction, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error)
	FindAll() (transactions *[]dto.ClientTransactionsResponse, err error)
	UsersFindAll(uip string) (transactions *[]dto.ClientTransactionsResponse, err error)
	UsersFindById(uid string, tid string) (transaction *dto.ClientTransactionsResponse, err error)
	CheckProductStock(pid string) (product *models.Product, err error)
	ProductReStock(pid string) (err error)
	Insert(data *models.Transaction) (transaction *models.Transaction, err error)
	InsertPayment(data *models.Payment) (transaction *models.Payment, err error)
	Update() (transaction *models.Transaction, err error)
	MidtransUpdate(tid string, status string) (err error)
	GetTransactionProduct(pid string) (product *models.Product, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (transaction models.Transaction, err error)
	GetAll() (transactions []dto.ClientTransactionsResponse, err error)
	UsersGetAll(uid string) (transactions []dto.ClientTransactionsResponse, err error)
	UsersGetById(uid string, tid string) (transaction dto.ClientTransactionsResponse, err error)
	Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (bill dto.BillClient, err error)
	GetBill(uid string, tid string) (bills dto.BillClient, err error)
	Modify() (transaction models.Transaction, err error)
	Remove() (err error)
	MidtransAfterPayment(midtransData dto.MidtransAfterPayment) (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
	midtrans   Midtrans
}

func NewService(repository Repository, midtransApi Midtrans) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
		midtrans:   midtransApi,
	}
}

func (s *service) GetById(id string) (transaction models.Transaction, err error) {

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
	transactions = *data
	return
}
func (s *service) UsersGetAll(uid string) (transactions []dto.ClientTransactionsResponse, err error) {
	data, err := s.repository.UsersFindAll(uid)
	if err != nil {
		return
	}
	transactions = *data
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
	payment, err := s.midtrans.DoPayment(createtransactionDTO.PaymentMethod, createtransactionDTO.GenerateMidtransPayment(tid))
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
		Charger:        dataPayment.Billed - dataProduct.Price,
		Deadline:       dataPayment.CreatedAt.Add(time.Hour * time.Duration(1)),
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
		Charger:        dataTransaction.Payment.Billed - dataProduct.Price,
		Deadline:       dataTransaction.Payment.CreatedAt.Add(time.Hour * time.Duration(1)),
	}
	return
}
func (s *service) MidtransAfterPayment(midtransData dto.MidtransAfterPayment) (err error) {
	_, err = uuid.Parse(midtransData.TransactionId)
	if err != nil {
		log.Print("Transaction Id ", midtransData.TransactionId, " Is Invalid")
		return
	}
	if midtransData.Code == "200" {
		log.Print("Update transaction skipped")
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
		log.Print("Midtrans Transaction Id ", midtransData.TransactionId, " Fail To Update")
	}
	return
}
func (s *service) Modify() (transaction models.Transaction, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
