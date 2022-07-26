package products

import (
	"backend_capstone/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (repository *ProductRepoMock) FindById(id string) *models.Product {
	// arguments := repository.Mock.Called(id)
	// if arguments.GetById(0) == nil {
	// 	return nil
	// } else {
	// 	product := arguments.GetById(0).(models.Product)
	// 	return &product
	// }
	return nil
}

func (repository *ProductRepoMock) FindAll() *models.Product {
	// arguments := repository.Mock.Called()
	// if arguments.GetAll(0) == nil {
	// 	return nil
	// } else {
	// 	product := arguments.GetAll(0).(models.Product)
	// 	return &product
	// }
	return nil
}

func (repository *ProductRepoMock) Insert(data []string) *models.Product {
	// arguments := repository.Mock.Insert(data[])
	// if arguments.Create() == nil {
	// 	return nil
	// } else {
	// 	product := arguments.Create(data[]).(models.Product)
	// 	return &product
	// }
	return nil
}

func (repository *ProductRepoMock) Update(id string) *models.Product {
	// arguments := repository.Mock.Update(Id)
	// if arguments.Modify(0) == nil {
	// 	return nil
	// } else {
	// 	product := arguments.Modify(0).(models.Product)
	// 	return &product
	// }
	return nil
}

func (repository *ProductRepoMock) Delete(id string) *models.Product {
	// arguments := repository.Mock.Delete(Id)
	// if arguments.Remove(0) == nil {
	// 	return nil
	// } else {
	// 	product := arguments.Remove(0).(models.Product)
	// 	return &product
	// }
	return nil
}
