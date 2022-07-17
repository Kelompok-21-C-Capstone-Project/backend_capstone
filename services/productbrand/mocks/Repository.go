// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	dto "backend_capstone/services/productbrand/dto"

	mock "github.com/stretchr/testify/mock"

	models "backend_capstone/models"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckBrandCategory provides a mock function with given fields: brandId, categoryId
func (_m *Repository) CheckBrandCategory(brandId string, categoryId string) (int64, error) {
	ret := _m.Called(brandId, categoryId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(brandId, categoryId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(brandId, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBrandCategory provides a mock function with given fields: brandId, categoryId
func (_m *Repository) DeleteBrandCategory(brandId string, categoryId string) error {
	ret := _m.Called(brandId, categoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(brandId, categoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: params
func (_m *Repository) FindAll(params ...string) (int64, *[]dto.ProductBrand, error) {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(...string) int64); ok {
		r0 = rf(params...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *[]dto.ProductBrand
	if rf, ok := ret.Get(1).(func(...string) *[]dto.ProductBrand); ok {
		r1 = rf(params...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*[]dto.ProductBrand)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(...string) error); ok {
		r2 = rf(params...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: id
func (_m *Repository) FindById(id string) (*models.ProductBrandResponse, error) {
	ret := _m.Called(id)

	var r0 *models.ProductBrandResponse
	if rf, ok := ret.Get(0).(func(string) *models.ProductBrandResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductBrandResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCategoryById provides a mock function with given fields: id
func (_m *Repository) FindCategoryById(id string) (*models.ProductCategory, error) {
	ret := _m.Called(id)

	var r0 *models.ProductCategory
	if rf, ok := ret.Get(0).(func(string) *models.ProductCategory); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductCategory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *Repository) Insert(data *models.ProductBrand) (*models.ProductBrandResponse, error) {
	ret := _m.Called(data)

	var r0 *models.ProductBrandResponse
	if rf, ok := ret.Get(0).(func(*models.ProductBrand) *models.ProductBrandResponse); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductBrandResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ProductBrand) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBrandCategory provides a mock function with given fields: brandId, categoryId, slug
func (_m *Repository) InsertBrandCategory(brandId string, categoryId string, slug string) (*models.ProductBrandCategory, error) {
	ret := _m.Called(brandId, categoryId, slug)

	var r0 *models.ProductBrandCategory
	if rf, ok := ret.Get(0).(func(string, string, string) *models.ProductBrandCategory); ok {
		r0 = rf(brandId, categoryId, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductBrandCategory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(brandId, categoryId, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, data
func (_m *Repository) Update(id string, data *models.ProductBrand) (*models.ProductBrandResponse, error) {
	ret := _m.Called(id, data)

	var r0 *models.ProductBrandResponse
	if rf, ok := ret.Get(0).(func(string, *models.ProductBrand) *models.ProductBrandResponse); ok {
		r0 = rf(id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductBrandResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *models.ProductBrand) error); ok {
		r1 = rf(id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
