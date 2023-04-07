// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	module "github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi3/module"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: _a0, _a1
func (_m *ProductRepository) DeleteById(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: _a0, _a1
func (_m *ProductRepository) FindAll(_a0 context.Context, _a1 *[]module.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]module.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: _a0, _a1
func (_m *ProductRepository) FindById(_a0 context.Context, _a1 string) (module.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 module.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (module.Product, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) module.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(module.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *ProductRepository) Insert(_a0 context.Context, _a1 *module.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *module.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *ProductRepository) Update(_a0 context.Context, _a1 module.Product, _a2 string) (module.Product, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 module.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, module.Product, string) (module.Product, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, module.Product, string) module.Product); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(module.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, module.Product, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
