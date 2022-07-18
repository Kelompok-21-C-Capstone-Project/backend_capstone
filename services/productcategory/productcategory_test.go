package productcategory_test

import (
	"backend_capstone/models"
	"backend_capstone/services/productcategory"
	"backend_capstone/services/productcategory/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductCategory_GetById(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductCategoryResponse{}, nil).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.NoError(t, err)
	})
	t.Run("fail: error not found in repository", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("some error")).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.Error(t, err)
	})
	t.Run("fail: error id format", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9")
		assert.Error(t, err)
	})
}

func TestProductCategory_Remove(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: error in id", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		err := service.Remove("06086f3c-7d3c-4d83")
		assert.Error(t, err)
	})
	t.Run("fail : error in id", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("error find by id")).Once()
		err := service.Remove("06086f3c-7d3c-4d83-a0fa-9f4a287dfbdc")
		assert.Error(t, err)
	})
	t.Run("fail : error in id", func(t *testing.T) {
		service := productcategory.NewService(mockRepo)
		mockRepo.On("Delete", mock.Anything).Return(nil, errors.New("error id")).Once()
		err := service.Remove("06086f3c-7d3c-4d83-a0fa-9f4a287dfbdc")
		assert.Error(t, err)
	})

}
