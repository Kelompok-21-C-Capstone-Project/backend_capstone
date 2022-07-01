package paymentvendor

import (
	"backend_capstone/models"
	"backend_capstone/services/paymentvendor/dto"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (paymentVendor *models.PaymentVendor, err error)
	FindByQuery(key string, value interface{}) (paymentVendors *[]models.PaymentVendor, err error)
	FindAll() (paymentVendors *[]models.PaymentVendor, err error)
	ValidatePaymentMethod(MethodId string) (err error)
	Insert(data *models.PaymentVendor) (paymentVendor *models.PaymentVendor, err error)
	Update(id string, data *models.PaymentVendor) (paymentVendor *models.PaymentVendor, err error)
	Delete(id string) (paymentVendor *models.PaymentVendor, err error)
}

type Service interface {
	GetById(id string) (paymentVendor models.PaymentVendor, err error)
	GetAll() (paymentVendors []models.PaymentVendor, err error)
	Create(createvendorDTO dto.CreateVendorDTO) (paymentVendor models.PaymentVendor, err error)
	Modify(id string, updatevendorDTO dto.UpdateVendorDTO) (paymentVendor models.PaymentVendor, err error)
	Remove(id string) (err error)
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
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	paymentVendor = *data
	return
}
func (s *service) GetAll() (paymentVendors []models.PaymentVendor, err error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return
	}
	paymentVendors = *data
	return
}
func (s *service) Create(createvendorDTO dto.CreateVendorDTO) (paymentVendor models.PaymentVendor, err error) {
	err = s.repository.ValidatePaymentMethod(createvendorDTO.PaymentMethodId)
	if err != nil {
		return
	}
	id := uuid.New()
	data, err := s.repository.Insert(createvendorDTO.GenerateModel(id.String()))
	if err != nil {
		return
	}
	paymentVendor = *data
	return
}
func (s *service) Modify(id string, updatevendorDTO dto.UpdateVendorDTO) (paymentVendor models.PaymentVendor, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	err = s.repository.ValidatePaymentMethod(updatevendorDTO.PaymentMethodId)
	if err != nil {
		return
	}
	data, err := s.repository.Insert(updatevendorDTO.GenerateModel(id))
	if err != nil {
		return
	}
	paymentVendor = *data
	return
}
func (s *service) Remove(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	s.repository.Delete(id)
	return
}
