package transaction

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (transaction *models.Transaction, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error)
	FindAll() (transactions *[]models.Transaction, err error)
	Insert() (transaction *models.Transaction, err error)
	Update() (transaction *models.Transaction, err error)
	Delete() (transaction *models.Transaction, err error)
}

type Service interface {
	GetById(id string) (transaction models.Transaction, err error)
	GetByUserId(id string) (transaction models.Transaction, err error)
	GetAll() (transactions []models.Transaction, err error)
	Create() (transaction models.Transaction, err error)
	Modify() (transaction models.Transaction, err error)
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

func (s *service) GetById(id string) (transaction models.Transaction, err error) {
	return
}
func (s *service) GetByUserId(id string) (transaction models.Transaction, err error) {
	return
}
func (s *service) GetAll() (transactions []models.Transaction, err error) {
	return
}
func (s *service) Create() (transaction models.Transaction, err error) {
	return
}
func (s *service) Modify() (transaction models.Transaction, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
