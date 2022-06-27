package productcategory

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (productCategory *models.ProductCategory, err error)
	FindByQuery(key string, value interface{}) (productCategorys *[]models.ProductCategory, err error)
	FindAll() (productCategorys *[]models.ProductCategory, err error)
	Insert() (productCategory *models.ProductCategory, err error)
	Update() (productCategory *models.ProductCategory, err error)
	Delete() (productCategory *models.ProductCategory, err error)
}

type Service interface {
	GetById(id string) (productCategory models.ProductCategory, err error)
	GetAll() (productCategorys []models.ProductCategory, err error)
	Create() (productCategory models.ProductCategory, err error)
	Modify() (productCategory models.ProductCategory, err error)
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

func (s *service) GetById(id string) (productCategory models.ProductCategory, err error) {
	return
}
func (s *service) GetAll() (productCategorys []models.ProductCategory, err error) {
	return
}
func (s *service) Create() (productCategory models.ProductCategory, err error) {
	return
}
func (s *service) Modify() (productCategory models.ProductCategory, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
