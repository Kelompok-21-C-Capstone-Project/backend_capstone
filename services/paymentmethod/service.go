package paymentmethod

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (paymentMethod *models.PaymentMethod, err error)
	FindByQuery(key string, value interface{}) (paymentMethods *[]models.PaymentMethod, err error)
	FindAll() (paymentMethods *[]models.PaymentMethod, err error)
	Insert() (paymentMethod *models.PaymentMethod, err error)
	Update() (paymentMethod *models.PaymentMethod, err error)
	Delete() (paymentMethod *models.PaymentMethod, err error)
}

type Service interface {
	GetById(id string) (paymentMethod models.PaymentMethod, err error)
	GetAll() (paymentMethods []models.PaymentMethod, err error)
	Create() (paymentMethod models.PaymentMethod, err error)
	Modify() (paymentMethod models.PaymentMethod, err error)
	Remove() (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetById(id string) (paymentMethod models.PaymentMethod, err error) {
	return
}
func (s *service) GetAll() (paymentMethods []models.PaymentMethod, err error) {
	return
}
func (s *service) Create() (paymentMethod models.PaymentMethod, err error) {
	return
}
func (s *service) Modify() (paymentMethod models.PaymentMethod, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
