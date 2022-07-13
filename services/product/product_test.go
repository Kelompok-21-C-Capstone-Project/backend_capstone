package product

import (
	"backend_capstone/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productrepo = &repository.ProductRepoMock{Mock: mock.Mock{}}
var productservice = service{repository: productrepo}

func TestProductService_GetById(t *testing.T) {

	// program mock
	productrepo.Mock.On("FindById", "1").Return(nil)

	product, err := productservice.GetById("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)

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

	product, err := productservice.Modify()

	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestProductService_CreateProduct(t *testing.T) {

	// program mock
	productrepo.Mock.On("Create", "id=?", "name=?", "price=?", "stock=?", "IsAvailable=?", "Description=?").Return(nil)

	product, err := productservice.Create()

	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestProductService_RemoveProduct(t *testing.T) {

	// program mock
	productrepo.Mock.On("Delete", "id=?").Return(nil)

	product, err := productservice.Remove()

	assert.Nil(t, product)
	assert.NotNil(t, err)

}
