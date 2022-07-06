package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	"backend_capstone/utils/midtransdriver"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (transaction *models.Transaction, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error)
	FindAll() (transactions *[]dto.ClientTransactionsResponse, err error)
	UsersFindAll(uip string) (transactions *[]dto.ClientTransactionsResponse, err error)
	CheckProductStock(pid string) (product *models.Product, err error)
	ProductReStock(pid string) (err error)
	Insert(data *models.Transaction) (transaction *models.Transaction, err error)
	InsertPayment(data *models.Payment) (transaction *models.Payment, err error)
	Update() (transaction *models.Transaction, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (transaction models.Transaction, err error)
	GetAll() (transactions []dto.ClientTransactionsResponse, err error)
	UsersGetAll(uid string) (transactions []dto.ClientTransactionsResponse, err error)
	Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (bill dto.BillClient, err error)
	Modify() (transaction models.Transaction, err error)
	Remove() (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
	midtrans   *midtransdriver.MidtransDriver
}

func NewService(repository Repository, midtransApi *midtransdriver.MidtransDriver) Service {
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
func (s *service) Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (bill dto.BillClient, err error) {
	if err = s.validate.Struct(createtransactionDTO); err != nil {
		return
	}
	// check stock barang
	dataProduct, err := s.repository.CheckProductStock(createtransactionDTO.ProductId)
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

	log.Print("Id", dataPayment.Id)
	log.Print("TransactionId", dataTransaction.Id)
	log.Print("VaNumber", dataPayment.Description)
	log.Print("PaymentDetails", dataPayment.MethodDetails)
	log.Print("Billed", dataPayment.Billed)
	log.Print("Product", dataProduct.Name)
	log.Print("ProductPrice", dataProduct.Price)
	log.Print("Charger", dataPayment.Billed)
	log.Print("Deadline", dataPayment.CreatedAt)

	bill = dto.BillClient{
		Id:             dataPayment.Id,
		TransactionId:  dataTransaction.Id,
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
func (s *service) Modify() (transaction models.Transaction, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
