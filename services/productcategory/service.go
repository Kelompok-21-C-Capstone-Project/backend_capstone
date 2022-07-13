package productcategory

import (
	"backend_capstone/models"
	"backend_capstone/services/productcategory/dto"
	"math"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (productCategory *models.ProductCategory, err error)
	FindAll(params ...string) (dataCount int64, productCategories *[]dto.ProductCategory, err error)
	Insert(data *models.ProductCategory) (productCategory *models.ProductCategory, err error)
	Update(id string, data *models.ProductCategory) (productCategory *models.ProductCategory, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (productCategory models.ProductCategory, err error)
	GetAll(params ...string) (data dto.ResponseBodyProductCategory, err error)
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
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	productCategory = *data
	return
}
func (s *service) GetAll(params ...string) (data dto.ResponseBodyProductCategory, err error) {
	if params[1] == "" {
		params[1] = "1"
	}
	if params[2] == "" {
		params[2] = "5"
	}
	nom, err := strconv.Atoi(params[1])
	if err != nil {
		return
	}
	if nom < 0 {
		params[1] = strconv.Itoa(nom)
	}
	den, err := strconv.Atoi(params[2])
	if err != nil {
		return
	}
	if den <= 0 {
		den = 5
	}
	dataCount, datas, err := s.repository.FindAll(params...)
	if err != nil {
		return
	}
	data.PageLength = int(math.Ceil(float64(dataCount) / float64(den)))
	if datas == nil {
		data.Data = []dto.ProductCategory{}
		return
	}
	data.Data = *datas
	return
}
func (s *service) Create(createcategoryDTO dto.CreateCategoryDTO) (productCategory models.ProductCategory, err error) {
	if err = s.validate.Struct(createcategoryDTO); err != nil {
		return
	}
	id := uuid.New()
	slug := strings.ReplaceAll(strings.ToLower(createcategoryDTO.Name), " ", "-")
	data, err := s.repository.Insert(createcategoryDTO.GenerateModel(id.String(), slug))
	productCategory = *data
	return
}
func (s *service) Modify(id string, updatecategoryDTO dto.UpdateCategoryDTO) (productCategory models.ProductCategory, err error) {
	if err = s.validate.Struct(updatecategoryDTO); err != nil {
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	slug := strings.ReplaceAll(strings.ToLower(updatecategoryDTO.Name), " ", "-")
	data, err := s.repository.Update(id, updatecategoryDTO.GenerateModel(id, slug))
	productCategory = *data
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
