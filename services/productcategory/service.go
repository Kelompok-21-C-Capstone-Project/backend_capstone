package productcategory

import (
	"backend_capstone/models"
	"backend_capstone/services/productcategory/dto"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (productCategory *models.ProductCategory, err error)
	FindByQuery(key string, value interface{}) (productCategories *[]models.ProductCategory, err error)
	FindAll() (productCategories *[]models.ProductCategory, err error)
	Insert(data *models.ProductCategory) (productCategory *models.ProductCategory, err error)
	Update(id string, data *models.ProductCategory) (productCategory *models.ProductCategory, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (productCategory models.ProductCategory, err error)
	GetAll() (productCategories []models.ProductCategory, err error)
	Create(createcategoryDTO dto.CreateCategoryDTO) (productCategory models.ProductCategory, err error)
	Modify(id string, updatecategoryDTO dto.UpdateCategoryDTO) (productCategory models.ProductCategory, err error)
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

func (s *service) GetById(id string) (productCategory models.ProductCategory, err error) {
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	productCategory = *data
	return
}
func (s *service) GetAll() (productCategories []models.ProductCategory, err error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return
	}
	productCategories = *data
	return
}
func (s *service) Create(createcategoryDTO dto.CreateCategoryDTO) (productCategory models.ProductCategory, err error) {
	if err = s.validate.Struct(createcategoryDTO); err != nil {
		return
	}
	id := uuid.New()
	slug := strings.ReplaceAll(strings.ToLower(createcategoryDTO.Name), " ", "-")
	data, err := s.repository.Insert(createcategoryDTO.GenerateModel(id, slug))
	productCategory = *data
	return
}
func (s *service) Modify(id string, updatecategoryDTO dto.UpdateCategoryDTO) (productCategory models.ProductCategory, err error) {
	if err = s.validate.Struct(updatecategoryDTO); err != nil {
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	slug := strings.ReplaceAll(strings.ToLower(updatecategoryDTO.Name), " ", "-")
	data, err := s.repository.Update(id, updatecategoryDTO.GenerateModel(uid, slug))
	productCategory = *data
	return
}
func (s *service) Remove(id string) (err error) {
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
