package productbrand

import (
	"backend_capstone/models"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id string) (productBrand *models.ProductBrand, err error)
	FindByQuery(key string, value interface{}) (productBrands *[]models.ProductBrand, err error)
	FindAll() (productBrands *[]models.ProductBrand, err error)
	Insert() (productBrand *models.ProductBrand, err error)
	Update() (productBrand *models.ProductBrand, err error)
	Delete() (productBrand *models.ProductBrand, err error)
}

type Service interface {
	GetById(id string) (productBrand models.ProductBrand, err error)
	GetAll() (productBrands []models.ProductBrand, err error)
	Create() (productBrand models.ProductBrand, err error)
	Modify() (productBrand models.ProductBrand, err error)
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

func (s *service) GetById(id string) (productBrand models.ProductBrand, err error) {
	return
}
func (s *service) GetAll() (productBrands []models.ProductBrand, err error) {
	return
}
func (s *service) Create() (productBrand models.ProductBrand, err error) {
	return
}
func (s *service) Modify() (productBrand models.ProductBrand, err error) {
	return
}
func (s *service) Remove() (err error) {
	return
}
