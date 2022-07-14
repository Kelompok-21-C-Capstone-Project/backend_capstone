package product

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"
	"log"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (product *models.ProductResponse, err error)
	FindAll(params ...string) (dataCount int64, products *[]dto.Product, err error)
	ClientFindAll() (products *[]dto.ProductCategory, err error)
	ClientFindAllBySlug(slug string) (products *dto.ProductCategory, err error)
	Insert(data *models.Product) (product *models.ProductResponse, err error)
	Update(id string, data *models.Product) (product *models.ProductResponse, err error)
	Delete(id string) (err error)
	ValidateProductBrandCategories(brandId string, categoryId string) (productBrandCategoriesId string, err error)
}

type Service interface {
	GetById(id string) (product models.ProductResponse, err error)
	GetAll(params ...string) (products dto.ResponseBodyProduct, err error)
	ClientGetAll() (products []dto.ProductCategory, err error)
	ClientGetAllBySlug(slug string) (products dto.ProductCategory, err error)
	GetAllByCategory(categoryId string) (products []models.ProductResponse, err error)
	Create(createproductDTO dto.CraeteProductDTO) (product models.ProductResponse, err error)
	Modify(id string, updateproductDTO dto.UpdateProductDTO) (product models.ProductResponse, err error)
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

func (s *service) GetById(id string) (product models.ProductResponse, err error) {
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
func (s *service) GetAll(params ...string) (products dto.ResponseBodyProduct, err error) {
	log.Print("enter service.GetAll")
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
	dataCount, data, err := s.repository.FindAll(params...)
	log.Print(dataCount)
	if err != nil {
		return
	}
	if den < -1 || den == 0 {
		den = 10
	} else if den == -1 {
		den = int(dataCount)
	}
	products.PageLength = int(math.Ceil(float64(dataCount) / float64(den)))
	if data == nil {
		products.Data = []dto.Product{}
		return
	}
	products.Data = *data
	return
}
func (s *service) ClientGetAllBySlug(slug string) (category dto.ProductCategory, err error) {
	data, err := s.repository.ClientFindAllBySlug(slug)
	if err != nil {
		return
	}
	category = *data
	for ip := range category.Products {
		if category.Products[ip].Id == "" {
			category.Products[ip] = nil
		}
	}
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
func (s *service) GetAllByCategory(categoryId string) (products []models.ProductResponse, err error) {
	return
}
func (s *service) Create(createproductDTO dto.CraeteProductDTO) (product models.ProductResponse, err error) {
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
func (s *service) Modify(id string, updateproductDTO dto.UpdateProductDTO) (product models.ProductResponse, err error) {
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
