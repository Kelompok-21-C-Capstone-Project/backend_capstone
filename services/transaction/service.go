package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	"backend_capstone/utils/midtransdriver"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (transaction *models.Transaction, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error)
	FindAll() (transactions *[]models.Transaction, err error)
	Insert(data *models.Transaction) (transaction *models.Transaction, err error)
	InsertPayment(data *models.Payment) (transaction *models.Payment, err error)
	Update() (transaction *models.Transaction, err error)
	Delete() (transaction *models.Transaction, err error)
}

type Service interface {
	GetById(id string) (transaction models.Transaction, err error)
	GetByUserId(id string) (transaction models.Transaction, err error)
	GetAll() (transactions []models.Transaction, err error)
	Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (transaction models.Transaction, err error)
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
func (s *service) GetByUserId(id string) (transaction models.Transaction, err error) {
	return
}
func (s *service) GetAll() (transactions []models.Transaction, err error) {
	return
}
func (s *service) Create(userId string, createtransactionDTO dto.CreateTransactionDTO) (transaction models.Transaction, err error) {
	if err = s.validate.Struct(createtransactionDTO); err != nil {
		return
	}
	_, err = uuid.Parse(userId)
	if err != nil {
		return
	}
	tid := uuid.New().String()
	s.repository.Insert(createtransactionDTO.GenerateTransactionModel(tid, userId))

	pid := uuid.New().String()
	s.repository.InsertPayment(createtransactionDTO.GeneratePaymentModel(pid, tid, "pending", "", "", ""))
	return
}
func (s *service) Modify() (transaction models.Transaction, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
