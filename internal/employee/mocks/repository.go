// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	employee "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: id, cardNum, firstName, lastName, warehouseId
func (_m *Repository) Create(id int, cardNum int, firstName string, lastName string, warehouseId int) (employee.Employee, error) {
	ret := _m.Called(id, cardNum, firstName, lastName, warehouseId)

	var r0 employee.Employee
	if rf, ok := ret.Get(0).(func(int, int, string, string, int) employee.Employee); ok {
		r0 = rf(id, cardNum, firstName, lastName, warehouseId)
	} else {
		r0 = ret.Get(0).(employee.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string, string, int) error); ok {
		r1 = rf(id, cardNum, firstName, lastName, warehouseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() []employee.Employee {
	ret := _m.Called()

	var r0 []employee.Employee
	if rf, ok := ret.Get(0).(func() []employee.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]employee.Employee)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int) (employee.Employee, error) {
	ret := _m.Called(id)

	var r0 employee.Employee
	if rf, ok := ret.Get(0).(func(int) employee.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(employee.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields:
func (_m *Repository) LastID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Update provides a mock function with given fields: emp, id
func (_m *Repository) Update(emp employee.Employee, id int) (employee.Employee, error) {
	ret := _m.Called(emp, id)

	var r0 employee.Employee
	if rf, ok := ret.Get(0).(func(employee.Employee, int) employee.Employee); ok {
		r0 = rf(emp, id)
	} else {
		r0 = ret.Get(0).(employee.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(employee.Employee, int) error); ok {
		r1 = rf(emp, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
