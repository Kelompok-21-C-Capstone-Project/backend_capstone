package payment

import (
	"backend_capstone/models"
	"log"

	"github.com/go-playground/validator/v10"
)

type PaymentRepository interface {
	FindById(id string) (payment *models.Payment, err error)
	FindByQuery(key string, value interface{}) (payments *[]models.Payment, err error)
	FindAll() (payments *[]models.Payment, err error)
	Insert() (payment *models.Payment, err error)
	Update() (payment *models.Payment, err error)
	Delete() (payment *models.Payment, err error)
}

type TransactionRepository interface {
	FindById(id string) (payment *models.Transaction, err error)
	Insert() (payment *models.Transaction, err error)
	Update() (payment *models.Transaction, err error)
}

type Service interface {
	GetById(id string) (payment models.Payment, err error)
	GetAll() (payments []models.Payment, err error)
	Remove() (err error)
}

type service struct {
	paymentRepository PaymentRepository
	validate          *validator.Validate
}

func NewService(paymentRepository PaymentRepository) Service {
	log.Print("Enter NewService")
	return &service{
		paymentRepository: paymentRepository,
		validate:          validator.New(),
	}
}

func (s *service) GetById(id string) (payment models.Payment, err error) {
	return
}
func (s *service) GetAll() (payments []models.Payment, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
