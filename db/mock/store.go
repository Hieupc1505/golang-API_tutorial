// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hieupc09/simple_api/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/hieupc09/simple_api/db/sqlc"
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

// CreateTodoTask mocks base method.
func (m *MockStore) CreateTodoTask(arg0 context.Context, arg1 db.CreateTodoTaskParams) (db.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodoTask", arg0, arg1)
	ret0, _ := ret[0].(db.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodoTask indicates an expected call of CreateTodoTask.
func (mr *MockStoreMockRecorder) CreateTodoTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodoTask", reflect.TypeOf((*MockStore)(nil).CreateTodoTask), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteTodoTask mocks base method.
func (m *MockStore) DeleteTodoTask(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodoTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodoTask indicates an expected call of DeleteTodoTask.
func (mr *MockStoreMockRecorder) DeleteTodoTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodoTask", reflect.TypeOf((*MockStore)(nil).DeleteTodoTask), arg0, arg1)
}

// GetTodoTask mocks base method.
func (m *MockStore) GetTodoTask(arg0 context.Context, arg1 int64) (db.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoTask", arg0, arg1)
	ret0, _ := ret[0].(db.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoTask indicates an expected call of GetTodoTask.
func (mr *MockStoreMockRecorder) GetTodoTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoTask", reflect.TypeOf((*MockStore)(nil).GetTodoTask), arg0, arg1)
}

// GetTodoTaskList mocks base method.
func (m *MockStore) GetTodoTaskList(arg0 context.Context, arg1 db.GetTodoTaskListParams) ([]db.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoTaskList", arg0, arg1)
	ret0, _ := ret[0].([]db.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoTaskList indicates an expected call of GetTodoTaskList.
func (mr *MockStoreMockRecorder) GetTodoTaskList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoTaskList", reflect.TypeOf((*MockStore)(nil).GetTodoTaskList), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// UpdateTodoTask mocks base method.
func (m *MockStore) UpdateTodoTask(arg0 context.Context, arg1 db.UpdateTodoTaskParams) (db.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodoTask", arg0, arg1)
	ret0, _ := ret[0].(db.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodoTask indicates an expected call of UpdateTodoTask.
func (mr *MockStoreMockRecorder) UpdateTodoTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodoTask", reflect.TypeOf((*MockStore)(nil).UpdateTodoTask), arg0, arg1)
}
