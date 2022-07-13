package productbrand

import (
	"backend_capstone/models"
	"backend_capstone/services/productbrand/dto"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (productBrand *models.ProductBrandResponse, err error)
	FindCategoryById(id string) (productCategory *models.ProductCategory, err error)
	FindAll(params ...string) (dataCount int64, productBrands *[]dto.ProductBrand, err error)
	Insert(data *models.ProductBrand) (productBrand *models.ProductBrandResponse, err error)
	Update(id string, data *models.ProductBrand) (productBrand *models.ProductBrandResponse, err error)
	Delete(id string) (err error)
	CheckBrandCategory(brandId string, categoryId string) (rowCount int64, err error)
	InsertBrandCategory(brandId string, categoryId string, slug string) (productBrand *models.ProductBrandCategory, err error)
	DeleteBrandCategory(brandId string, categoryId string) (err error)
}

type Service interface {
	GetById(id string) (productBrand models.ProductBrandResponse, err error)
	GetAll(params ...string) (productBrands dto.ResponseBodyProductBrand, err error)
	Create(createbrandDTO dto.CreateBrandDTO) (productBrand models.ProductBrandResponse, err error)
	Modify(id string, updatebrandDTO dto.UpdateBrandDTO) (productBrand models.ProductBrandResponse, err error)
	Remove(id string) (err error)
	AddBrandCategory(brandId string, categoryId string) (productBrandCategory models.ProductBrandCategory, err error)
	RemoveBrandCategory(brandId string, categoryId string) (err error)
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

func (s *service) GetById(id string) (productBrand models.ProductBrandResponse, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	productBrand = *data
	return
}
func (s *service) GetAll(params ...string) (productBrands dto.ResponseBodyProductBrand, err error) {
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
	productBrands.PageLength = int(math.Ceil(float64(dataCount) / float64(den)))
	if datas == nil {
		productBrands.Data = []dto.ProductBrand{}
		return
	}
	productBrands.Data = *datas
	return
}
func (s *service) Create(createbrandDTO dto.CreateBrandDTO) (productBrand models.ProductBrandResponse, err error) {
	if err = s.validate.Struct(createbrandDTO); err != nil {
		return
	}
	id := uuid.New()
	data, err := s.repository.Insert(createbrandDTO.GenerateModel(id.String()))
	productBrand = *data
	return
}
func (s *service) Modify(id string, updatebrandDTO dto.UpdateBrandDTO) (productBrand models.ProductBrandResponse, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	data, err := s.repository.Update(id, updatebrandDTO.GenerateModel(id))
	productBrand = *data
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
func (s *service) AddBrandCategory(brandId string, categoryId string) (productBrandCategory models.ProductBrandCategory, err error) {
	_, err = uuid.Parse(brandId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(categoryId)
	if err != nil {
		return
	}
	_, err = s.repository.CheckBrandCategory(brandId, categoryId)
	if err != nil {
		return
	}
	dataCategory, err := s.repository.FindCategoryById(categoryId)
	if err != nil {
		return
	}
	dataBrand, err := s.repository.FindById(brandId)
	if err != nil {
		return
	}
	slug := dataBrand.Name + " - " + dataCategory.Name
	data, err := s.repository.InsertBrandCategory(brandId, categoryId, slug)
	productBrandCategory = *data
	return
}
func (s *service) RemoveBrandCategory(brandId string, categoryId string) (err error) {
	_, err = uuid.Parse(brandId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(categoryId)
	if err != nil {
		return
	}
	_, err = s.repository.FindCategoryById(categoryId)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(brandId)
	if err != nil {
		return
	}
	s.repository.DeleteBrandCategory(brandId, categoryId)
	return
}
