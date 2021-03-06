// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/izaakdale/goBank/domain (interfaces: CustomerRepo)

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/izaakdale/goBank/domain"
	errs "github.com/izaakdale/goBank/errs"
)

// MockCustomerRepo is a mock of CustomerRepo interface.
type MockCustomerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepoMockRecorder
}

// MockCustomerRepoMockRecorder is the mock recorder for MockCustomerRepo.
type MockCustomerRepoMockRecorder struct {
	mock *MockCustomerRepo
}

// NewMockCustomerRepo creates a new mock instance.
func NewMockCustomerRepo(ctrl *gomock.Controller) *MockCustomerRepo {
	mock := &MockCustomerRepo{ctrl: ctrl}
	mock.recorder = &MockCustomerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepo) EXPECT() *MockCustomerRepoMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockCustomerRepo) FindAll(arg0 string) ([]domain.Customer, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]domain.Customer)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCustomerRepoMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCustomerRepo)(nil).FindAll), arg0)
}

// FindById mocks base method.
func (m *MockCustomerRepo) FindById(arg0 string) (*domain.Customer, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*domain.Customer)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCustomerRepoMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCustomerRepo)(nil).FindById), arg0)
}
