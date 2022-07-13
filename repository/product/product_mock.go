package products

import (
	"backend_capstone/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (repository *ProductRepoMock) FindById(id string) *models.Product {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		product := arguments.Get(0).(models.Product)
		return &product
	}
}
