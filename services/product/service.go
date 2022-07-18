package product

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"
	"errors"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (product *models.ProductResponse, err error)
	FindAll(params ...string) (products dto.ResponseBodyProduct, err error)
	ClientFindAll() (products *[]dto.ProductCategory, err error)
	ClientFindAllBySlug(slug string) (products *dto.ProductCategory, err error)
	Insert(data *models.Product) (product *models.ProductResponse, err error)
	Update(id string, data *models.Product) (product *models.ProductResponse, err error)
	Delete(id string) (err error)
	ValidateProductBrandCategories(brandId string, categoryId string) (productBrandCategoriesId string, err error)
	UpdateStock(data *dto.UpdateStockDTO) (err error)
	CreateSupplyProduct(dataSupply models.Supply, dataProducts []models.SupplyProduct) (err error)
	GetSupplyInvocie(params ...string) (data dto.DataSupplyDTO, err error)
}

type Service interface {
	GetById(id string) (product models.ProductResponse, err error)
	GetAll(params ...string) (products dto.ResponseBodyProduct, err error)
	ClientGetAll() (products []dto.ProductCategory, err error)
	ClientGetAllBySlug(slug string) (products dto.ProductCategory, err error)
	Create(createproductDTO dto.CraeteProductDTO) (product models.ProductResponse, err error)
	Modify(id string, updateproductDTO dto.UpdateProductDTO) (product models.ProductResponse, err error)
	ModifyStock(data *dto.UpdateStockDTO) (err error)
	GetSupplyInvocie(params ...string) (data dto.DataSupplyDTO, err error)
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
	products, err = s.repository.FindAll(params...)
	if err != nil {
		return
	}
	if den < -1 {
		products.PageLength = int64(math.Ceil(float64(products.PageLength) / float64(10)))
	} else if den == -1 || den == 0 {
		products.PageLength = 1
	} else {
		products.PageLength = int64(math.Ceil(float64(products.PageLength) / float64(den)))
	}
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
			category.Products = append(category.Products[:ip], category.Products[ip+1:]...)
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
				products[ic].Products = append(products[ic].Products[:ip], products[ic].Products[ip+1:]...)
			}
		}
	}
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
func (s *service) GetSupplyInvocie(params ...string) (data dto.DataSupplyDTO, err error) {
	log.Print("enter service.GetAll")
	if params[3] == "" {
		params[3] = time.Now().Format("02-01-2006") + "_" + time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	}
	if params[3] != "" {
		regexDateRange := "([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])_([0-9])([0-9])-([0-9])([0-9])-([0-9])([0-9])([0-9])([0-9])"
		if resDR, _ := regexp.MatchString(regexDateRange, params[3]); !resDR {
			return
		}
		date := strings.Split(params[3], "_")
		dateTop, _ := time.Parse("02-01-2006 15:04:05", date[1]+" 08:04:00")
		date[1] = dateTop.AddDate(0, 0, 1).Format("02-01-2006")
		params = append(params, date...)
	}
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
	data, err = s.repository.GetSupplyInvocie(params...)
	if err != nil {
		return
	}
	if den < -1 {
		data.PageLength = int64(math.Ceil(float64(data.PageLength) / float64(10)))
	} else if den == -1 || den == 0 {
		data.PageLength = 1
	} else {
		data.PageLength = int64(math.Ceil(float64(data.PageLength) / float64(den)))
	}
	if data.Data == nil {
		data.Data = []dto.SupplyProductDTO{}
		return
	}
	return
}
func (s *service) ModifyStock(data *dto.UpdateStockDTO) (err error) {
	log.Print("service.product.ModifyStock")
	if err = s.validate.Struct(data); err != nil {
		return
	}
	if err = s.validate.Struct(data.Datas); err != nil {
		return
	}
	_, err = uuid.Parse((*data).AdminId)
	if err != nil {
		return
	}
	if len(data.Datas) == 0 {
		err = errors.New("Data tidak boleh kosong")
		return
	}
	if err = s.repository.UpdateStock(data); err != nil {
		log.Print("error di sini")
		return
	}
	id := uuid.New().String()
	dataSupply := models.Supply{
		Id:      id,
		CodeNo:  "INV/" + id[:3] + "/" + id[len(id)-3:],
		AdminId: data.AdminId,
	}
	dataProducts := []models.SupplyProduct{}
	for _, el := range data.Datas {
		dataSupply.SumStock = dataSupply.SumStock + el.Stock
		dataSupply.Name = dataSupply.Name + ", " + el.Name
		dataProducts = append(dataProducts, models.SupplyProduct{
			SupplyId:  id,
			ProductId: el.Id,
			Amount:    el.Stock,
			Cost:      el.Price,
		})
	}
	dataSupply.Name = dataSupply.Name[2:]
	if err = s.repository.CreateSupplyProduct(dataSupply, dataProducts); err != nil {
		return
	}
	return
}
