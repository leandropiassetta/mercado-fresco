// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	seller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, cid, companyName, address, telephone
func (_m *Repository) Create(ctx context.Context, cid int, companyName string, address string, telephone string) (seller.Seller, error) {
	ret := _m.Called(ctx, cid, companyName, address, telephone)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) seller.Seller); ok {
		r0 = rf(ctx, cid, companyName, address, telephone)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string) error); ok {
		r1 = rf(ctx, cid, companyName, address, telephone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]seller.Seller, error) {
	ret := _m.Called(ctx)

	var r0 []seller.Seller
	if rf, ok := ret.Get(0).(func(context.Context) []seller.Seller); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]seller.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: ctx, id
func (_m *Repository) GetOne(ctx context.Context, id int) (seller.Seller, error) {
	ret := _m.Called(ctx, id)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(context.Context, int) seller.Seller); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cid, companyName, address, telephone, _a5
func (_m *Repository) Update(ctx context.Context, cid int, companyName string, address string, telephone string, _a5 seller.Seller) (seller.Seller, error) {
	ret := _m.Called(ctx, cid, companyName, address, telephone, _a5)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string, seller.Seller) seller.Seller); ok {
		r0 = rf(ctx, cid, companyName, address, telephone, _a5)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string, seller.Seller) error); ok {
		r1 = rf(ctx, cid, companyName, address, telephone, _a5)
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