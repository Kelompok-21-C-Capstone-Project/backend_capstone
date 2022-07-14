package product_test

import (
	"backend_capstone/models"
	"backend_capstone/services/product"
	"backend_capstone/services/product/dto"
	"backend_capstone/services/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productrepo = &mocks.Repository{Mock: mock.Mock{}}
var productservice = product.NewService(productrepo)

func TestProductService_GetById(t *testing.T) {

	// program mock
	productrepo.Mock.On("FindById", "1").Return(nil)

	product, err := productservice.GetById("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
}

func TestGetById(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductResponse{}, nil)
		service := product.NewService(mockRepo)
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.NoError(t, err)
	})
}

func TestProductService_GetAll(t *testing.T) {

	// program mock
	productrepo.Mock.On("FindAll").Return(nil)

	product, err := productservice.GetAll()

	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestProductService_ModifyProduct(t *testing.T) {

	// program mock
	productrepo.Mock.On("Update", "1").Return(nil)

	product, err := productservice.Modify("x", dto.UpdateProductDTO{})

	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestProductService_CreateProduct(t *testing.T) {

	// program mock
	productrepo.Mock.On("Create", "id=?", "name=?", "price=?", "stock=?", "IsAvailable=?", "Description=?").Return(nil)

	product, err := productservice.Create(dto.CraeteProductDTO{})

	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestProductService_RemoveProduct(t *testing.T) {

	// program mock
	productrepo.Mock.On("Delete", "id=?").Return(nil)

	err := productservice.Remove("x")

	assert.NotNil(t, err)

}
