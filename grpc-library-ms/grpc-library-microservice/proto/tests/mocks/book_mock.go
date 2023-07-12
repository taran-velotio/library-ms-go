// Code generated by MockGen. DO NOT EDIT.
// Source: book_grpc.pb.go

// Package mock_book is a generated GoMock package.
package mock_book

import (
	context "context"
	book "library-comp/proto/book"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBookServiceClient is a mock of BookServiceClient interface.
type MockBookServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceClientMockRecorder
}

// MockBookServiceClientMockRecorder is the mock recorder for MockBookServiceClient.
type MockBookServiceClientMockRecorder struct {
	mock *MockBookServiceClient
}

// NewMockBookServiceClient creates a new mock instance.
func NewMockBookServiceClient(ctrl *gomock.Controller) *MockBookServiceClient {
	mock := &MockBookServiceClient{ctrl: ctrl}
	mock.recorder = &MockBookServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceClient) EXPECT() *MockBookServiceClientMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServiceClient) CreateBook(ctx context.Context, in *book.CreateBookRequest, opts ...grpc.CallOption) (*book.CreateBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBook", varargs...)
	ret0, _ := ret[0].(*book.CreateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceClientMockRecorder) CreateBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServiceClient)(nil).CreateBook), varargs...)
}

// DeleteBook mocks base method.
func (m *MockBookServiceClient) DeleteBook(ctx context.Context, in *book.DeleteBookRequest, opts ...grpc.CallOption) (*book.DeleteBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBook", varargs...)
	ret0, _ := ret[0].(*book.DeleteBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceClientMockRecorder) DeleteBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookServiceClient)(nil).DeleteBook), varargs...)
}

// GetBook mocks base method.
func (m *MockBookServiceClient) GetBook(ctx context.Context, in *book.GetBookRequest, opts ...grpc.CallOption) (*book.GetBookResonse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBook", varargs...)
	ret0, _ := ret[0].(*book.GetBookResonse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServiceClientMockRecorder) GetBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookServiceClient)(nil).GetBook), varargs...)
}

// GetListOfBooks mocks base method.
func (m *MockBookServiceClient) GetListOfBooks(ctx context.Context, in *book.GetListOfBooksRequest, opts ...grpc.CallOption) (*book.GetListOfBooksResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetListOfBooks", varargs...)
	ret0, _ := ret[0].(*book.GetListOfBooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfBooks indicates an expected call of GetListOfBooks.
func (mr *MockBookServiceClientMockRecorder) GetListOfBooks(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfBooks", reflect.TypeOf((*MockBookServiceClient)(nil).GetListOfBooks), varargs...)
}

// UpdateBook mocks base method.
func (m *MockBookServiceClient) UpdateBook(ctx context.Context, in *book.UpdateBookRequest, opts ...grpc.CallOption) (*book.UpdateBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBook", varargs...)
	ret0, _ := ret[0].(*book.UpdateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceClientMockRecorder) UpdateBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookServiceClient)(nil).UpdateBook), varargs...)
}

// MockBookServiceServer is a mock of BookServiceServer interface.
type MockBookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceServerMockRecorder
}

// MockBookServiceServerMockRecorder is the mock recorder for MockBookServiceServer.
type MockBookServiceServerMockRecorder struct {
	mock *MockBookServiceServer
}

// NewMockBookServiceServer creates a new mock instance.
func NewMockBookServiceServer(ctrl *gomock.Controller) *MockBookServiceServer {
	mock := &MockBookServiceServer{ctrl: ctrl}
	mock.recorder = &MockBookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceServer) EXPECT() *MockBookServiceServerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServiceServer) CreateBook(arg0 context.Context, arg1 *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(*book.CreateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceServerMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServiceServer)(nil).CreateBook), arg0, arg1)
}

// DeleteBook mocks base method.
func (m *MockBookServiceServer) DeleteBook(arg0 context.Context, arg1 *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(*book.DeleteBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceServerMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookServiceServer)(nil).DeleteBook), arg0, arg1)
}

// GetBook mocks base method.
func (m *MockBookServiceServer) GetBook(arg0 context.Context, arg1 *book.GetBookRequest) (*book.GetBookResonse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0, arg1)
	ret0, _ := ret[0].(*book.GetBookResonse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServiceServerMockRecorder) GetBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookServiceServer)(nil).GetBook), arg0, arg1)
}

// GetListOfBooks mocks base method.
func (m *MockBookServiceServer) GetListOfBooks(arg0 context.Context, arg1 *book.GetListOfBooksRequest) (*book.GetListOfBooksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListOfBooks", arg0, arg1)
	ret0, _ := ret[0].(*book.GetListOfBooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfBooks indicates an expected call of GetListOfBooks.
func (mr *MockBookServiceServerMockRecorder) GetListOfBooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfBooks", reflect.TypeOf((*MockBookServiceServer)(nil).GetListOfBooks), arg0, arg1)
}

// UpdateBook mocks base method.
func (m *MockBookServiceServer) UpdateBook(arg0 context.Context, arg1 *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(*book.UpdateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceServerMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookServiceServer)(nil).UpdateBook), arg0, arg1)
}

// mustEmbedUnimplementedBookServiceServer mocks base method.
func (m *MockBookServiceServer) mustEmbedUnimplementedBookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServiceServer")
}

// mustEmbedUnimplementedBookServiceServer indicates an expected call of mustEmbedUnimplementedBookServiceServer.
func (mr *MockBookServiceServerMockRecorder) mustEmbedUnimplementedBookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServiceServer", reflect.TypeOf((*MockBookServiceServer)(nil).mustEmbedUnimplementedBookServiceServer))
}

// MockUnsafeBookServiceServer is a mock of UnsafeBookServiceServer interface.
type MockUnsafeBookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBookServiceServerMockRecorder
}

// MockUnsafeBookServiceServerMockRecorder is the mock recorder for MockUnsafeBookServiceServer.
type MockUnsafeBookServiceServerMockRecorder struct {
	mock *MockUnsafeBookServiceServer
}

// NewMockUnsafeBookServiceServer creates a new mock instance.
func NewMockUnsafeBookServiceServer(ctrl *gomock.Controller) *MockUnsafeBookServiceServer {
	mock := &MockUnsafeBookServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBookServiceServer) EXPECT() *MockUnsafeBookServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBookServiceServer mocks base method.
func (m *MockUnsafeBookServiceServer) mustEmbedUnimplementedBookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServiceServer")
}

// mustEmbedUnimplementedBookServiceServer indicates an expected call of mustEmbedUnimplementedBookServiceServer.
func (mr *MockUnsafeBookServiceServerMockRecorder) mustEmbedUnimplementedBookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServiceServer", reflect.TypeOf((*MockUnsafeBookServiceServer)(nil).mustEmbedUnimplementedBookServiceServer))
}
