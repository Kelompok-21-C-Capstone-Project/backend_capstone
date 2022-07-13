package product_test

import (
	"backend_capstone/models"
	"backend_capstone/services/product"
	"backend_capstone/services/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	var mockRepo = mocks.Repository{Mock: mock.Mock{}}
	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductResponse{}, nil)
		service := product.NewService(mockRepo)
		_, err := service.GetById("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
		assert.NoError(t, err)
	})
}
