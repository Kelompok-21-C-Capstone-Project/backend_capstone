package product

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (transaction *models.Product, err error)
	FindByQuery(key string, value interface{}) (transactions *[]models.Product, err error)
	FindAll() (transactions *[]models.Product, err error)
	Insert() (transaction *models.Product, err error)
	Update() (transaction *models.Product, err error)
	Delete() (transaction *models.Product, err error)
	GetProductBrandCategories() (productBrandCategories []models.ProductBrandCategory, err error)
}

type Service interface {
	GetById(id string) (transaction models.Product, err error)
	GetByUserId(id string) (transaction models.Product, err error)
	GetAll() (transactions []models.Product, err error)
	Create() (transaction models.Product, err error)
	Modify() (transaction models.Product, err error)
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

func (s *service) GetById(id string) (transaction models.Product, err error) {
	return
}
func (s *service) GetByUserId(id string) (transaction models.Product, err error) {
	return
}
func (s *service) GetAll() (transactions []models.Product, err error) {
	return
}
func (s *service) Create() (transaction models.Product, err error) {
	return
}
func (s *service) Modify() (transaction models.Product, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
