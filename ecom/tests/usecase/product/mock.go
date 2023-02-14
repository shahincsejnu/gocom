// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../usecase/product/usecase.go

// Package mock_product is a generated GoMock package.
package mock_product

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	product "github.com/shahincsejnu/gocom/ecom/domain/product"
	db "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(ctx context.Context, opts *product.CreationOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepository)(nil).Create), ctx, opts)
}

// DeleteOne mocks base method.
func (m *MockProductRepository) DeleteOne(ctx context.Context, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockProductRepositoryMockRecorder) DeleteOne(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockProductRepository)(nil).DeleteOne), ctx, productID)
}

// GetAll mocks base method.
func (m *MockProductRepository) GetAll(ctx context.Context) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductRepository)(nil).GetAll), ctx)
}

// GetOne mocks base method.
func (m *MockProductRepository) GetOne(ctx context.Context, productID string) (*db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", ctx, productID)
	ret0, _ := ret[0].(*db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockProductRepositoryMockRecorder) GetOne(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockProductRepository)(nil).GetOne), ctx, productID)
}

// Update mocks base method.
func (m *MockProductRepository) Update(ctx context.Context, opts *product.UpdateOptions, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, opts, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(ctx, opts, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), ctx, opts, productID)
}