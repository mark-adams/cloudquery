// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: MySQLServerClient,MySQLConfigurationClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	gomock "github.com/golang/mock/gomock"
)

// MockMySQLServerClient is a mock of MySQLServerClient interface.
type MockMySQLServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLServerClientMockRecorder
}

// MockMySQLServerClientMockRecorder is the mock recorder for MockMySQLServerClient.
type MockMySQLServerClientMockRecorder struct {
	mock *MockMySQLServerClient
}

// NewMockMySQLServerClient creates a new mock instance.
func NewMockMySQLServerClient(ctrl *gomock.Controller) *MockMySQLServerClient {
	mock := &MockMySQLServerClient{ctrl: ctrl}
	mock.recorder = &MockMySQLServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMySQLServerClient) EXPECT() *MockMySQLServerClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockMySQLServerClient) List(arg0 context.Context) (mysql.ServerListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(mysql.ServerListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMySQLServerClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMySQLServerClient)(nil).List), arg0)
}

// MockMySQLConfigurationClient is a mock of MySQLConfigurationClient interface.
type MockMySQLConfigurationClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLConfigurationClientMockRecorder
}

// MockMySQLConfigurationClientMockRecorder is the mock recorder for MockMySQLConfigurationClient.
type MockMySQLConfigurationClientMockRecorder struct {
	mock *MockMySQLConfigurationClient
}

// NewMockMySQLConfigurationClient creates a new mock instance.
func NewMockMySQLConfigurationClient(ctrl *gomock.Controller) *MockMySQLConfigurationClient {
	mock := &MockMySQLConfigurationClient{ctrl: ctrl}
	mock.recorder = &MockMySQLConfigurationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMySQLConfigurationClient) EXPECT() *MockMySQLConfigurationClientMockRecorder {
	return m.recorder
}

// ListByServer mocks base method.
func (m *MockMySQLConfigurationClient) ListByServer(arg0 context.Context, arg1, arg2 string) (mysql.ConfigurationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(mysql.ConfigurationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByServer indicates an expected call of ListByServer.
func (mr *MockMySQLConfigurationClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockMySQLConfigurationClient)(nil).ListByServer), arg0, arg1, arg2)
}
