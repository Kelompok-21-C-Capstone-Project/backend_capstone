package product

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id uuid.UUID) (product *models.Product, err error)
	FindByQuery(key string, value interface{}) (products *[]models.Product, err error)
	FindAll() (products *[]models.Product, err error)
	Insert(data *models.Product) (product *models.Product, err error)
	Update(id uuid.UUID, data *models.Product) (product *models.Product, err error)
	Delete(id uuid.UUID) (product *models.Product, err error)
	ValidateProductBrandCategories(brandId uuid.UUID, categoryId uuid.UUID) (productBrandCategoriesId uuid.UUID, err error)
}

type Service interface {
	GetById(id string) (product models.Product, err error)
	GetAll() (products []models.Product, err error)
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
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.FindById(uid)
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
func (s *service) GetAllByCategory(categoryId string) (products []models.Product, err error) {
	return
}
func (s *service) Create(createproductDTO dto.CraeteProductDTO) (product models.Product, err error) {
	if err = s.validate.Struct(createproductDTO); err != nil {
		return
	}
	bid, err := uuid.Parse(createproductDTO.BrandId)
	if err != nil {
		return
	}
	cid, err := uuid.Parse(createproductDTO.CategoryId)
	if err != nil {
		return
	}
	pbcid, err := s.repository.ValidateProductBrandCategories(bid, cid)
	if err != nil {
		return
	}
	id := uuid.New()
	data, err := s.repository.Insert(createproductDTO.GenerateModel(id, pbcid))
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
	bid, err := uuid.Parse(updateproductDTO.BrandId)
	if err != nil {
		return
	}
	cid, err := uuid.Parse(updateproductDTO.CategoryId)
	if err != nil {
		return
	}
	pbcid, err := s.repository.ValidateProductBrandCategories(bid, cid)
	if err != nil {
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	data, err := s.repository.Update(uid, updateproductDTO.GenerateModel(uid, pbcid))
	if err != nil {
		return
	}
	product = *data
	return
}
func (s *service) Remove(id string) (err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	s.repository.Delete(uid)
	return
}
