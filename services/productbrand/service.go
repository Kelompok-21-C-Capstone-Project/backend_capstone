package productbrand

import (
	"backend_capstone/models"
	"backend_capstone/services/productbrand/dto"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (productBrand *models.ProductBrand, err error)
	FindCategoryById(id string) (productCategory *models.ProductCategory, err error)
	FindByQuery(key string, value interface{}) (productBrands *[]models.ProductBrand, err error)
	FindAll() (productBrands *[]models.ProductBrand, err error)
	Insert(data *models.ProductBrand) (productBrand *models.ProductBrand, err error)
	Update(id string, data *models.ProductBrand) (productBrand *models.ProductBrand, err error)
	Delete(id string) (err error)
	CheckBrandCategory(brandId string, categoryId string) (rowCount int64, err error)
	InsertBrandCategory(brandId string, categoryId string, slug string) (productBrand *models.ProductBrandCategory, err error)
	DeleteBrandCategory(brandId string, categoryId string) (err error)
}

type Service interface {
	GetById(id string) (productBrand models.ProductBrand, err error)
	GetAll() (productBrands []models.ProductBrand, err error)
	Create(createbrandDTO dto.CreateBrandDTO) (productBrand models.ProductBrand, err error)
	Modify(id string, updatebrandDTO dto.UpdateBrandDTO) (productBrand models.ProductBrand, err error)
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

func (s *service) GetById(id string) (productBrand models.ProductBrand, err error) {
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
func (s *service) GetAll() (productBrands []models.ProductBrand, err error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return
	}
	productBrands = *data
	return
}
func (s *service) Create(createbrandDTO dto.CreateBrandDTO) (productBrand models.ProductBrand, err error) {
	if err = s.validate.Struct(createbrandDTO); err != nil {
		return
	}
	id := uuid.New()
	data, err := s.repository.Insert(createbrandDTO.GenerateModel(id.String()))
	productBrand = *data
	return
}
func (s *service) Modify(id string, updatebrandDTO dto.UpdateBrandDTO) (productBrand models.ProductBrand, err error) {
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
