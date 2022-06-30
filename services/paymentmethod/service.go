package paymentmethod

import (
	"backend_capstone/models"
	"backend_capstone/services/paymentmethod/dto"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (paymentMethod *models.PaymentMethod, err error)
	FindByQuery(key string, value interface{}) (paymentMethods *[]models.PaymentMethod, err error)
	FindAll() (paymentMethods *[]models.PaymentMethod, err error)
	Insert(data *models.PaymentMethod) (paymentMethod *models.PaymentMethod, err error)
	Update(id string, data *models.PaymentMethod) (paymentMethod *models.PaymentMethod, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (paymentMethod models.PaymentMethod, err error)
	GetAll() (paymentMethods []models.PaymentMethod, err error)
	Create(methodDTO dto.CreateMethodDTO) (paymentMethod models.PaymentMethod, err error)
	Modify(id string, methodDTO dto.CreateMethodDTO) (paymentMethod models.PaymentMethod, err error)
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

func (s *service) GetById(id string) (paymentMethod models.PaymentMethod, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	paymentMethod = *data
	return
}
func (s *service) GetAll() (paymentMethods []models.PaymentMethod, err error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return
	}
	paymentMethods = *data
	return
}
func (s *service) Create(methodDTO dto.CreateMethodDTO) (paymentMethod models.PaymentMethod, err error) {
	if err = s.validate.Struct(methodDTO); err != nil {
		return
	}
	id := uuid.New().String()
	slug := strings.ReplaceAll(strings.ToLower(methodDTO.Name), " ", "-")
	data, err := s.repository.Insert(methodDTO.GenerateModel(id, slug))
	paymentMethod = *data
	return
}
func (s *service) Modify(id string, methodDTO dto.CreateMethodDTO) (paymentMethod models.PaymentMethod, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	slug := strings.ReplaceAll(strings.ToLower(methodDTO.Name), " ", "-")
	data, err := s.repository.Insert(methodDTO.GenerateModel(id, slug))
	paymentMethod = *data
	return
}
func (s *service) Remove(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	err = s.repository.Delete(id)
	if err != nil {
		return
	}
	return
}
