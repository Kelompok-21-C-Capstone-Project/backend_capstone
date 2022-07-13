package product

import (
	"backend_capstone/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productrepo = &repository.ProductRepoMock{Mock: mock.Mock{}}
var productservice = service{repository: productrepo}

func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	productrepo.Mock.On("FindById", "1").Return(nil)

	product, err := productservice.GetById("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)

}
