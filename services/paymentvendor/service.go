package paymentvendor

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (paymentVendor *models.PaymentVendor, err error)
	FindByQuery(key string, value interface{}) (paymentVendors *[]models.PaymentVendor, err error)
	FindAll() (paymentVendors *[]models.PaymentVendor, err error)
	Insert() (paymentVendor *models.PaymentVendor, err error)
	Update() (paymentVendor *models.PaymentVendor, err error)
	Delete() (paymentVendor *models.PaymentVendor, err error)
}

type Service interface {
	GetById(id string) (paymentVendor models.PaymentVendor, err error)
	GetAll() (paymentVendors []models.PaymentVendor, err error)
	Create() (paymentVendor models.PaymentVendor, err error)
	Modify() (paymentVendor models.PaymentVendor, err error)
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

func (s *service) GetById(id string) (paymentVendor models.PaymentVendor, err error) {
	return
}
func (s *service) GetAll() (paymentVendors []models.PaymentVendor, err error) {
	return
}
func (s *service) Create() (paymentVendor models.PaymentVendor, err error) {
	return
}
func (s *service) Modify() (paymentVendor models.PaymentVendor, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
