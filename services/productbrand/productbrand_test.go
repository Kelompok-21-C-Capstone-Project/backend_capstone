package productbrand_test

import (
	"backend_capstone/models"
	"backend_capstone/services/productbrand"
	"backend_capstone/services/productbrand/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductBrand_GetById(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductBrandResponse{}, nil).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.NoError(t, err)
	})
	t.Run("fail: error not found in repository", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("some error")).Once()
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9-a58b7c0b6bf5")
		assert.Error(t, err)
	})
	t.Run("fail: error id format", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		_, err := service.GetById("5f8fe693-fbee-438d-9ca9")
		assert.Error(t, err)
	})
}

func TestProductBrand_Remove(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		err := service.Remove("06086f3c-7d3c-4d83")
		assert.Error(t, err)
	})
	t.Run("fail : error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(nil, errors.New("error find by id")).Once()
		err := service.Remove("06086f3c-7d3c-4d83-a0fa-9f4a287dfbdc")
		assert.Error(t, err)
	})
	t.Run("fail : error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		mockRepo.On("FindById", mock.Anything).Return(&models.ProductBrandResponse{}, nil).Once()
		mockRepo.On("Delete", mock.Anything).Return(errors.New("error id")).Once()
		err := service.Remove("06086f3c-7d3c-4d83-a0fa-9f4a287dfbdc")
		assert.Error(t, err)
	})

}

func TestProductBrand_AddBrandCategory(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		_, err := service.AddBrandCategory("0ec4be23-a140-4df0-", "14cb58f4-be6f-46fb-9096-bf5d57e4a0df")
		assert.Error(t, err)
	})
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		_, err := service.AddBrandCategory("0ec4be23-a140-4df0-8096-1e5b0181391f", "14cb58f4-be6f-46fb")
		assert.Error(t, err)
	})
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		mockRepo.On("CheckBrandCategory", mock.Anything, mock.Anything).Return(int64(0), errors.New("error id")).Once()
		_, err := service.AddBrandCategory("0ec4be23-a140-4df0-8096-1e5b0181391f", "14cb58f4-be6f-46fb-9096-bf5d57e4a0df")
		assert.Error(t, err)
	})

}

func TestProductBrand_RemoveBrandCategory(t *testing.T) {
	var mockRepo = new(mocks.Repository)
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		err := service.RemoveBrandCategory("0ec4be23-a140-4df0", "14cb58f4-be6f-46fb-9096-bf5d57e4a0df")
		assert.Error(t, err)
	})
	t.Run("fail: error in id", func(t *testing.T) {
		service := productbrand.NewService(mockRepo)
		err := service.RemoveBrandCategory("0ec4be23-a140-4df0-8096-1e5b0181391f", "bf5d57e4a0df")
		assert.Error(t, err)
	})

}
