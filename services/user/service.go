package user

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id int) (transaction *models.User, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.User, err error)
	FindAll() (transactions *[]models.User, err error)
	Insert() (transaction *models.User, err error)
	Update() (transaction *models.User, err error)
	Delete() (transaction *models.User, err error)
}

type Service interface {
	GetById(id int) (transaction models.User, err error)
	GetByUserId(id int) (transaction models.User, err error)
	GetAll() (transactions []models.User, err error)
	Create() (transaction models.User, err error)
	Modify() (transaction models.User, err error)
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

func (s *service) GetById(id int) (transaction models.User, err error) {
	return
}
func (s *service) GetByUserId(id int) (transaction models.User, err error) {
	return
}
func (s *service) GetAll() (transactions []models.User, err error) {
	return
}
func (s *service) Create() (transaction models.User, err error) {
	return
}
func (s *service) Modify() (transaction models.User, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
