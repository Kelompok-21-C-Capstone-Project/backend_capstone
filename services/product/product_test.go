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
	var mockRepo = new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductResponse{}, nil)
		service := product.NewService(mockRepo)
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.NoError(t, err)
	})
}
