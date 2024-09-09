// Code generated by MockGen. DO NOT EDIT.
// Source: b30northwindapi/services (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	db "b30northwindapi/db/sqlc"
	models "b30northwindapi/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockStore) CreateCart(arg0 context.Context, arg1 db.CreateCartParams) (*db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0, arg1)
	ret0, _ := ret[0].(*db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockStoreMockRecorder) CreateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockStore)(nil).CreateCart), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 *models.CategoryPostReq) (*db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(*db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// // CreateOrder mocks base method.
// func (m *MockStore) CreateOrder(arg0 context.Context, arg1 db.CreateOrderParams) (*db.Order, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Order)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // CreateOrder indicates an expected call of CreateOrder.
// func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
// }

// // CreateOrderTx mocks base method.
// func (m *MockStore) CreateOrderTx(arg0 context.Context, arg1 db.CreateOrderParams) (*db.Order, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "CreateOrderTx", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Order)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // CreateOrderTx indicates an expected call of CreateOrderTx.
// func (mr *MockStoreMockRecorder) CreateOrderTx(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderTx", reflect.TypeOf((*MockStore)(nil).CreateOrderTx), arg0, arg1)
// }

// // CreateProduct mocks base method.
// func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (*db.Product, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Product)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // CreateProduct indicates an expected call of CreateProduct.
// func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
// }

// // CreateUser mocks base method.
// func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (*db.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
// 	ret0, _ := ret[0].(*db.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // CreateUser indicates an expected call of CreateUser.
// func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
// }

// // DeleteCart mocks base method.
// func (m *MockStore) DeleteCart(arg0 context.Context, arg1 int32) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteCart", arg0, arg1)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // DeleteCart indicates an expected call of DeleteCart.
// func (mr *MockStoreMockRecorder) DeleteCart(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockStore)(nil).DeleteCart), arg0, arg1)
// }

// // DeleteCategory mocks base method.
// func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int32) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // DeleteCategory indicates an expected call of DeleteCategory.
// func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
// }

// // DeleteOrder mocks base method.
// func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 int16) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // DeleteOrder indicates an expected call of DeleteOrder.
// func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
// }

// // DeleteProduct mocks base method.
// func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int16) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // DeleteProduct indicates an expected call of DeleteProduct.
// func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
// }

// // DeleteToken mocks base method.
// func (m *MockStore) DeleteToken(arg0 context.Context, arg1 *string) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteToken", arg0, arg1)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // DeleteToken indicates an expected call of DeleteToken.
// func (mr *MockStoreMockRecorder) DeleteToken(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockStore)(nil).DeleteToken), arg0, arg1)
// }

// // FindAllCategory mocks base method.
// func (m *MockStore) FindAllCategory(arg0 context.Context) ([]*db.Category, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindAllCategory", arg0)
// 	ret0, _ := ret[0].([]*db.Category)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindAllCategory indicates an expected call of FindAllCategory.
// func (mr *MockStoreMockRecorder) FindAllCategory(arg0 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCategory", reflect.TypeOf((*MockStore)(nil).FindAllCategory), arg0)
// }

// // FindAllOrder mocks base method.
// func (m *MockStore) FindAllOrder(arg0 context.Context) ([]*db.Order, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindAllOrder", arg0)
// 	ret0, _ := ret[0].([]*db.Order)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindAllOrder indicates an expected call of FindAllOrder.
// func (mr *MockStoreMockRecorder) FindAllOrder(arg0 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllOrder", reflect.TypeOf((*MockStore)(nil).FindAllOrder), arg0)
// }

// // FindAllProduct mocks base method.
// func (m *MockStore) FindAllProduct(arg0 context.Context) ([]*db.Product, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindAllProduct", arg0)
// 	ret0, _ := ret[0].([]*db.Product)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindAllProduct indicates an expected call of FindAllProduct.
// func (mr *MockStoreMockRecorder) FindAllProduct(arg0 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllProduct", reflect.TypeOf((*MockStore)(nil).FindAllProduct), arg0)
// }

// // FindAllProductPaging mocks base method.
// func (m *MockStore) FindAllProductPaging(arg0 context.Context, arg1 db.FindAllProductPagingParams) ([]*db.Product, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindAllProductPaging", arg0, arg1)
// 	ret0, _ := ret[0].([]*db.Product)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindAllProductPaging indicates an expected call of FindAllProductPaging.
// func (mr *MockStoreMockRecorder) FindAllProductPaging(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllProductPaging", reflect.TypeOf((*MockStore)(nil).FindAllProductPaging), arg0, arg1)
// }

// // FindCartByCustomerAndProduct mocks base method.
// func (m *MockStore) FindCartByCustomerAndProduct(arg0 context.Context, arg1 db.FindCartByCustomerAndProductParams) (*db.FindCartByCustomerAndProductRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindCartByCustomerAndProduct", arg0, arg1)
// 	ret0, _ := ret[0].(*db.FindCartByCustomerAndProductRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindCartByCustomerAndProduct indicates an expected call of FindCartByCustomerAndProduct.
// func (mr *MockStoreMockRecorder) FindCartByCustomerAndProduct(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartByCustomerAndProduct", reflect.TypeOf((*MockStore)(nil).FindCartByCustomerAndProduct), arg0, arg1)
// }

// // FindCartByCustomerId mocks base method.
// func (m *MockStore) FindCartByCustomerId(arg0 context.Context, arg1 string) ([]*db.FindCartByCustomerIdRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindCartByCustomerId", arg0, arg1)
// 	ret0, _ := ret[0].([]*db.FindCartByCustomerIdRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindCartByCustomerId indicates an expected call of FindCartByCustomerId.
// func (mr *MockStoreMockRecorder) FindCartByCustomerId(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartByCustomerId", reflect.TypeOf((*MockStore)(nil).FindCartByCustomerId), arg0, arg1)
// }

// // FindCartByCustomerPaging mocks base method.
// func (m *MockStore) FindCartByCustomerPaging(arg0 context.Context, arg1 db.FindCartByCustomerPagingParams) ([]*db.FindCartByCustomerPagingRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindCartByCustomerPaging", arg0, arg1)
// 	ret0, _ := ret[0].([]*db.FindCartByCustomerPagingRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindCartByCustomerPaging indicates an expected call of FindCartByCustomerPaging.
// func (mr *MockStoreMockRecorder) FindCartByCustomerPaging(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartByCustomerPaging", reflect.TypeOf((*MockStore)(nil).FindCartByCustomerPaging), arg0, arg1)
// }

// // FindCategoryById mocks base method.
// func (m *MockStore) FindCategoryById(arg0 context.Context, arg1 int32) (*db.Category, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindCategoryById", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Category)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindCategoryById indicates an expected call of FindCategoryById.
// func (mr *MockStoreMockRecorder) FindCategoryById(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryById", reflect.TypeOf((*MockStore)(nil).FindCategoryById), arg0, arg1)
// }

// // FindOrderById mocks base method.
// func (m *MockStore) FindOrderById(arg0 context.Context, arg1 int16) (*db.Order, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindOrderById", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Order)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindOrderById indicates an expected call of FindOrderById.
// func (mr *MockStoreMockRecorder) FindOrderById(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderById", reflect.TypeOf((*MockStore)(nil).FindOrderById), arg0, arg1)
// }

// // FindProductById mocks base method.
// func (m *MockStore) FindProductById(arg0 context.Context, arg1 int16) (*db.Product, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindProductById", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Product)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindProductById indicates an expected call of FindProductById.
// func (mr *MockStoreMockRecorder) FindProductById(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductById", reflect.TypeOf((*MockStore)(nil).FindProductById), arg0, arg1)
// }

// // FindUserByPhone mocks base method.
// func (m *MockStore) FindUserByPhone(arg0 context.Context, arg1 *string) (*db.FindUserByPhoneRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindUserByPhone", arg0, arg1)
// 	ret0, _ := ret[0].(*db.FindUserByPhoneRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindUserByPhone indicates an expected call of FindUserByPhone.
// func (mr *MockStoreMockRecorder) FindUserByPhone(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByPhone", reflect.TypeOf((*MockStore)(nil).FindUserByPhone), arg0, arg1)
// }

// // FindUserByUserPassword mocks base method.
// func (m *MockStore) FindUserByUserPassword(arg0 context.Context, arg1 db.FindUserByUserPasswordParams) (*db.FindUserByUserPasswordRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindUserByUserPassword", arg0, arg1)
// 	ret0, _ := ret[0].(*db.FindUserByUserPasswordRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindUserByUserPassword indicates an expected call of FindUserByUserPassword.
// func (mr *MockStoreMockRecorder) FindUserByUserPassword(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUserPassword", reflect.TypeOf((*MockStore)(nil).FindUserByUserPassword), arg0, arg1)
// }

// // FindUserByUsername mocks base method.
// func (m *MockStore) FindUserByUsername(arg0 context.Context, arg1 *string) (*db.FindUserByUsernameRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FindUserByUsername", arg0, arg1)
// 	ret0, _ := ret[0].(*db.FindUserByUsernameRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // FindUserByUsername indicates an expected call of FindUserByUsername.
// func (mr *MockStoreMockRecorder) FindUserByUsername(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsername", reflect.TypeOf((*MockStore)(nil).FindUserByUsername), arg0, arg1)
// }

// // GetUserRoles mocks base method.
// func (m *MockStore) GetUserRoles(arg0 context.Context, arg1 db.GetUserRolesParams) ([]*db.GetUserRolesRow, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "GetUserRoles", arg0, arg1)
// 	ret0, _ := ret[0].([]*db.GetUserRolesRow)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // GetUserRoles indicates an expected call of GetUserRoles.
// func (mr *MockStoreMockRecorder) GetUserRoles(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRoles", reflect.TypeOf((*MockStore)(nil).GetUserRoles), arg0, arg1)
// }

// // Signin mocks base method.
// func (m *MockStore) Signin(arg0 context.Context, arg1 models.CreateUserReq) (*models.UserResponse, *models.Error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Signin", arg0, arg1)
// 	ret0, _ := ret[0].(*models.UserResponse)
// 	ret1, _ := ret[1].(*models.Error)
// 	return ret0, ret1
// }

// // Signin indicates an expected call of Signin.
// func (mr *MockStoreMockRecorder) Signin(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signin", reflect.TypeOf((*MockStore)(nil).Signin), arg0, arg1)
// }

// // Signout mocks base method.
// func (m *MockStore) Signout(arg0 context.Context, arg1 string) *models.Error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Signout", arg0, arg1)
// 	ret0, _ := ret[0].(*models.Error)
// 	return ret0
// }

// // Signout indicates an expected call of Signout.
// func (mr *MockStoreMockRecorder) Signout(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signout", reflect.TypeOf((*MockStore)(nil).Signout), arg0, arg1)
// }

// // Signup mocks base method.
// func (m *MockStore) Signup(arg0 context.Context, arg1 models.CreateUserReq) (*models.UserResponse, *models.Error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Signup", arg0, arg1)
// 	ret0, _ := ret[0].(*models.UserResponse)
// 	ret1, _ := ret[1].(*models.Error)
// 	return ret0, ret1
// }

// // Signup indicates an expected call of Signup.
// func (mr *MockStoreMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockStore)(nil).Signup), arg0, arg1)
// }

// // UpdateCartQty mocks base method.
// func (m *MockStore) UpdateCartQty(arg0 context.Context, arg1 db.UpdateCartQtyParams) (*db.Cart, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateCartQty", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Cart)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateCartQty indicates an expected call of UpdateCartQty.
// func (mr *MockStoreMockRecorder) UpdateCartQty(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartQty", reflect.TypeOf((*MockStore)(nil).UpdateCartQty), arg0, arg1)
// }

// // UpdateCategory mocks base method.
// func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (*db.Category, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Category)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateCategory indicates an expected call of UpdateCategory.
// func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
// }

// // UpdateOrderShip mocks base method.
// func (m *MockStore) UpdateOrderShip(arg0 context.Context, arg1 db.UpdateOrderShipParams) (*db.Order, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateOrderShip", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Order)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateOrderShip indicates an expected call of UpdateOrderShip.
// func (mr *MockStoreMockRecorder) UpdateOrderShip(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderShip", reflect.TypeOf((*MockStore)(nil).UpdateOrderShip), arg0, arg1)
// }

// // UpdateProduct mocks base method.
// func (m *MockStore) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (*db.Product, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
// 	ret0, _ := ret[0].(*db.Product)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateProduct indicates an expected call of UpdateProduct.
// func (mr *MockStoreMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), arg0, arg1)
// }

// // UpdateToken mocks base method.
// func (m *MockStore) UpdateToken(arg0 context.Context, arg1 db.UpdateTokenParams) (*db.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1)
// 	ret0, _ := ret[0].(*db.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateToken indicates an expected call of UpdateToken.
// func (mr *MockStoreMockRecorder) UpdateToken(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockStore)(nil).UpdateToken), arg0, arg1)
// }

// // UpdateUserName mocks base method.
// func (m *MockStore) UpdateUserName(arg0 context.Context, arg1 db.UpdateUserNameParams) (*db.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateUserName", arg0, arg1)
// 	ret0, _ := ret[0].(*db.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateUserName indicates an expected call of UpdateUserName.
// func (mr *MockStoreMockRecorder) UpdateUserName(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserName", reflect.TypeOf((*MockStore)(nil).UpdateUserName), arg0, arg1)
// }

// // UpdateUserPhone mocks base method.
// func (m *MockStore) UpdateUserPhone(arg0 context.Context, arg1 db.UpdateUserPhoneParams) (*db.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateUserPhone", arg0, arg1)
// 	ret0, _ := ret[0].(*db.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateUserPhone indicates an expected call of UpdateUserPhone.
// func (mr *MockStoreMockRecorder) UpdateUserPhone(arg0, arg1 interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPhone", reflect.TypeOf((*MockStore)(nil).UpdateUserPhone), arg0, arg1)
// }
