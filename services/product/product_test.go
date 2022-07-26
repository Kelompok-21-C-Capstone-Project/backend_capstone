package product_test

import (
	"backend_capstone/models"
	"backend_capstone/services/product"
	"backend_capstone/services/product/dto"
	"backend_capstone/services/product/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_GetById(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductResponse{}, nil).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.NoError(t, err)
	})
	t.Run("fail: error not found in repository", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("some error")).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.Error(t, err)
	})
	t.Run("fail: error id format", func(t *testing.T) {
		service := product.NewService(mockRepo)
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9")
		assert.Error(t, err)
	})
}

func TestProductService_GetAll(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: parsing parameter page", func(t *testing.T) {
		service := product.NewService(mockRepo)
		_, err := service.GetAll("", "&$!^@#", "12")
		assert.Error(t, err)
	})
	t.Run("fail: error parameter", func(t *testing.T) {
		service := product.NewService(mockRepo)
		_, err := service.GetAll("", "-1", "satu")
		assert.Error(t, err)
	})
	t.Run("fail: error for datacount", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(dto.ResponseBodyProduct{}, errors.New("error datacount")).Once()
		_, err := service.GetAll("key", "1", "2")
		assert.Error(t, err)
	})
	t.Run("success", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything).Return(dto.ResponseBodyProduct{}, nil).Once()
		_, err := service.GetAll("key", "1", "2")
		assert.NoError(t, err)
	})

}

func TestProductService_ClientGetAllBySlug(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: error parameter slug", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ClientFindAllBySlug", mock.Anything).Return(nil, errors.New("there is error")).Once()
		_, err := service.ClientGetAllBySlug("slug")
		assert.Error(t, err)
	})
	t.Run("success", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ClientFindAllBySlug", mock.Anything).Return(&dto.ProductCategory{}, nil).Once()
		_, err := service.ClientGetAllBySlug("orang-tua")
		assert.NoError(t, err)
	})

}

func TestProductService_ClientAll(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: fetch all client", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ClientFindAll", mock.Anything).Return(nil, errors.New("there is error"))
		_, err := service.ClientGetAll()
		assert.Error(t, err)
	})
}

func TestProductService_Create(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: brand id is invalid", func(t *testing.T) {
		service := product.NewService(mockRepo)
		dtoData := dto.CraeteProductDTO{
			Name:        "some product",
			Price:       3000,
			PriceBuy:    1000,
			Profit:      2000,
			Stock:       10,
			Description: "some product description",
			BrandId:     "5f8fe693-fbee-438d-9ca9-",
			CategoryId:  "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			IsAvailable: false,
		}
		_, err := service.Create(dtoData)
		assert.Error(t, err)
	})
	t.Run("fail: category id is invalid", func(t *testing.T) {
		service := product.NewService(mockRepo)
		dtoData := dto.CraeteProductDTO{
			Name:        "some product",
			Price:       3000,
			PriceBuy:    1000,
			Profit:      2000,
			Stock:       10,
			Description: "some product description",
			BrandId:     "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			CategoryId:  "5f8fe693-fbee-438d-9ca9-",
			IsAvailable: false,
		}
		_, err := service.Create(dtoData)
		assert.Error(t, err)
	})
	t.Run("fail: category/brand id is not found", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ValidateProductBrandCategories", mock.Anything, mock.Anything).Return("", errors.New("category not found")).Once()
		dtoData := dto.CraeteProductDTO{
			Name:        "some product",
			Price:       3000,
			PriceBuy:    1000,
			Profit:      2000,
			Stock:       10,
			Description: "some product description",
			BrandId:     "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			CategoryId:  "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			IsAvailable: false,
		}
		_, err := service.Create(dtoData)
		assert.Error(t, err)
	})
	t.Run("fail: insert fail", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ValidateProductBrandCategories", mock.Anything, mock.Anything).Return("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5", nil).Once()
		mockRepo.On("Insert", mock.Anything).Return(nil, errors.New("insert fail")).Once()
		dtoData := dto.CraeteProductDTO{
			Name:        "some product",
			Price:       3000,
			PriceBuy:    1000,
			Profit:      2000,
			Stock:       10,
			Description: "some product description",
			BrandId:     "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			CategoryId:  "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			IsAvailable: false,
		}
		_, err := service.Create(dtoData)
		assert.Error(t, err)
	})
	t.Run("success", func(t *testing.T) {
		service := product.NewService(mockRepo)
		mockRepo.On("ValidateProductBrandCategories", mock.Anything, mock.Anything).Return("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5", nil).Once()
		mockRepo.On("Insert", mock.Anything).Return(&models.ProductResponse{}, nil).Once()
		dtoData := dto.CraeteProductDTO{
			Name:        "some product",
			Price:       3000,
			PriceBuy:    1000,
			Profit:      2000,
			Stock:       10,
			Description: "some product description",
			BrandId:     "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			CategoryId:  "5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5",
			IsAvailable: false,
		}
		_, err := service.Create(dtoData)
		assert.NoError(t, err)
	})
}

func TestProductService_Remove(t *testing.T) {
	// var mockRepo = new(mocks.Repository)
	// t.Run("fail: error id", func(t *testing.T) {
	// 	service := product.NewService(mockRepo)
	// 	err := service.Remove("06086f3c-7d3c-4d83")
	// 	assert.Error(t, err)
	// })
	// t.Run("success", func(t *testing.T) {
	// 	service := product.NewService(mockRepo)
	// 	mockRepo.On("Delete", mock.Anything).Return(mockRepo.Delete, nil).Once()
	// 	err := service.Remove("06086f3c-7d3c-4d83-a0fa-9f4a287dfbdc")
	// 	assert.NoError(t, err)
	// })
}

func TestProductService_CreateProduct(t *testing.T) {

	// program mock
	// productrepo.Mock.On("Create", "id=?", "name=?", "price=?", "stock=?", "IsAvailable=?", "Description=?").Return(nil)

	// product, err := productservice.Create(dto.CraeteProductDTO{})

	// assert.Nil(t, product)
	// assert.NotNil(t, err)

}

func TestProductService_RemoveProduct(t *testing.T) {

	// program mock
	// productrepo.Mock.On("Delete", "id=?").Return(nil)

	// err := productservice.Remove("x")

	// assert.NotNil(t, err)

}
