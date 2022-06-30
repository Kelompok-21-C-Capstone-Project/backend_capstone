package product

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (product *models.Product, err error)
	FindByQuery(key string, value interface{}) (products *[]models.Product, err error)
	FindAll() (products *[]models.Product, err error)
	ClientFindAll() (products *[]dto.ProductCategory, err error)
	Insert(data *models.Product) (product *models.Product, err error)
	Update(id string, data *models.Product) (product *models.Product, err error)
	Delete(id string) (err error)
	ValidateProductBrandCategories(brandId string, categoryId string) (productBrandCategoriesId string, err error)
}

type Service interface {
	GetById(id string) (product models.Product, err error)
	GetAll() (products []models.Product, err error)
	ClientGetAll() (products []dto.ProductCategory, err error)
	GetAllByCategory(categoryId string) (products []models.Product, err error)
	Create(createproductDTO dto.CraeteProductDTO) (product models.Product, err error)
	Modify(id string, updateproductDTO dto.UpdateProductDTO) (product models.Product, err error)
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

func (s *service) GetById(id string) (product models.Product, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	product = *data
	return
}
func (s *service) GetAll() (products []models.Product, err error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return
	}
	products = *data
	return
}
func (s *service) ClientGetAll() (products []dto.ProductCategory, err error) {
	data, err := s.repository.ClientFindAll()
	if err != nil {
		return
	}
	products = *data
	for ic, category := range products {
		for ip := range category.Products {
			if products[ic].Products[ip].Id == "" {
				products[ic].Products[ip] = nil
			}
		}
	}
	return
}
func (s *service) GetAllByCategory(categoryId string) (products []models.Product, err error) {
	return
}
func (s *service) Create(createproductDTO dto.CraeteProductDTO) (product models.Product, err error) {
	if err = s.validate.Struct(createproductDTO); err != nil {
		return
	}
	_, err = uuid.Parse(createproductDTO.BrandId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(createproductDTO.CategoryId)
	if err != nil {
		return
	}
	pbcid, err := s.repository.ValidateProductBrandCategories(createproductDTO.BrandId, createproductDTO.CategoryId)
	if err != nil {
		return
	}
	id := uuid.New()
	data, err := s.repository.Insert(createproductDTO.GenerateModel(id.String(), pbcid))
	if err != nil {
		return
	}
	product = *data
	return
}
func (s *service) Modify(id string, updateproductDTO dto.UpdateProductDTO) (product models.Product, err error) {
	if err = s.validate.Struct(updateproductDTO); err != nil {
		return
	}
	_, err = uuid.Parse(updateproductDTO.BrandId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(updateproductDTO.CategoryId)
	if err != nil {
		return
	}
	pbcid, err := s.repository.ValidateProductBrandCategories(updateproductDTO.BrandId, updateproductDTO.CategoryId)
	if err != nil {
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.Update(id, updateproductDTO.GenerateModel(id, pbcid))
	if err != nil {
		return
	}
	product = *data
	return
}
func (s *service) Remove(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	s.repository.Delete(id)
	return
}
