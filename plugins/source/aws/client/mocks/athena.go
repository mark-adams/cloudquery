// Code generated by MockGen. DO NOT EDIT.
// Source: athena.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	athena "github.com/aws/aws-sdk-go-v2/service/athena"
	gomock "github.com/golang/mock/gomock"
)

// MockAthenaClient is a mock of AthenaClient interface.
type MockAthenaClient struct {
	ctrl     *gomock.Controller
	recorder *MockAthenaClientMockRecorder
}

// MockAthenaClientMockRecorder is the mock recorder for MockAthenaClient.
type MockAthenaClientMockRecorder struct {
	mock *MockAthenaClient
}

// NewMockAthenaClient creates a new mock instance.
func NewMockAthenaClient(ctrl *gomock.Controller) *MockAthenaClient {
	mock := &MockAthenaClient{ctrl: ctrl}
	mock.recorder = &MockAthenaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAthenaClient) EXPECT() *MockAthenaClientMockRecorder {
	return m.recorder
}

// BatchGetNamedQuery mocks base method.
func (m *MockAthenaClient) BatchGetNamedQuery(arg0 context.Context, arg1 *athena.BatchGetNamedQueryInput, arg2 ...func(*athena.Options)) (*athena.BatchGetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetNamedQuery", varargs...)
	ret0, _ := ret[0].(*athena.BatchGetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetNamedQuery indicates an expected call of BatchGetNamedQuery.
func (mr *MockAthenaClientMockRecorder) BatchGetNamedQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetNamedQuery", reflect.TypeOf((*MockAthenaClient)(nil).BatchGetNamedQuery), varargs...)
}

// BatchGetPreparedStatement mocks base method.
func (m *MockAthenaClient) BatchGetPreparedStatement(arg0 context.Context, arg1 *athena.BatchGetPreparedStatementInput, arg2 ...func(*athena.Options)) (*athena.BatchGetPreparedStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetPreparedStatement", varargs...)
	ret0, _ := ret[0].(*athena.BatchGetPreparedStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetPreparedStatement indicates an expected call of BatchGetPreparedStatement.
func (mr *MockAthenaClientMockRecorder) BatchGetPreparedStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetPreparedStatement", reflect.TypeOf((*MockAthenaClient)(nil).BatchGetPreparedStatement), varargs...)
}

// BatchGetQueryExecution mocks base method.
func (m *MockAthenaClient) BatchGetQueryExecution(arg0 context.Context, arg1 *athena.BatchGetQueryExecutionInput, arg2 ...func(*athena.Options)) (*athena.BatchGetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetQueryExecution", varargs...)
	ret0, _ := ret[0].(*athena.BatchGetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetQueryExecution indicates an expected call of BatchGetQueryExecution.
func (mr *MockAthenaClientMockRecorder) BatchGetQueryExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetQueryExecution", reflect.TypeOf((*MockAthenaClient)(nil).BatchGetQueryExecution), varargs...)
}

// GetDataCatalog mocks base method.
func (m *MockAthenaClient) GetDataCatalog(arg0 context.Context, arg1 *athena.GetDataCatalogInput, arg2 ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataCatalog", varargs...)
	ret0, _ := ret[0].(*athena.GetDataCatalogOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataCatalog indicates an expected call of GetDataCatalog.
func (mr *MockAthenaClientMockRecorder) GetDataCatalog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataCatalog", reflect.TypeOf((*MockAthenaClient)(nil).GetDataCatalog), varargs...)
}

// GetDatabase mocks base method.
func (m *MockAthenaClient) GetDatabase(arg0 context.Context, arg1 *athena.GetDatabaseInput, arg2 ...func(*athena.Options)) (*athena.GetDatabaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDatabase", varargs...)
	ret0, _ := ret[0].(*athena.GetDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockAthenaClientMockRecorder) GetDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockAthenaClient)(nil).GetDatabase), varargs...)
}

// GetNamedQuery mocks base method.
func (m *MockAthenaClient) GetNamedQuery(arg0 context.Context, arg1 *athena.GetNamedQueryInput, arg2 ...func(*athena.Options)) (*athena.GetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamedQuery", varargs...)
	ret0, _ := ret[0].(*athena.GetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamedQuery indicates an expected call of GetNamedQuery.
func (mr *MockAthenaClientMockRecorder) GetNamedQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamedQuery", reflect.TypeOf((*MockAthenaClient)(nil).GetNamedQuery), varargs...)
}

// GetPreparedStatement mocks base method.
func (m *MockAthenaClient) GetPreparedStatement(arg0 context.Context, arg1 *athena.GetPreparedStatementInput, arg2 ...func(*athena.Options)) (*athena.GetPreparedStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPreparedStatement", varargs...)
	ret0, _ := ret[0].(*athena.GetPreparedStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreparedStatement indicates an expected call of GetPreparedStatement.
func (mr *MockAthenaClientMockRecorder) GetPreparedStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreparedStatement", reflect.TypeOf((*MockAthenaClient)(nil).GetPreparedStatement), varargs...)
}

// GetQueryExecution mocks base method.
func (m *MockAthenaClient) GetQueryExecution(arg0 context.Context, arg1 *athena.GetQueryExecutionInput, arg2 ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryExecution", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryExecution indicates an expected call of GetQueryExecution.
func (mr *MockAthenaClientMockRecorder) GetQueryExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryExecution", reflect.TypeOf((*MockAthenaClient)(nil).GetQueryExecution), varargs...)
}

// GetQueryResults mocks base method.
func (m *MockAthenaClient) GetQueryResults(arg0 context.Context, arg1 *athena.GetQueryResultsInput, arg2 ...func(*athena.Options)) (*athena.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryResults", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResults indicates an expected call of GetQueryResults.
func (mr *MockAthenaClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResults", reflect.TypeOf((*MockAthenaClient)(nil).GetQueryResults), varargs...)
}

// GetQueryRuntimeStatistics mocks base method.
func (m *MockAthenaClient) GetQueryRuntimeStatistics(arg0 context.Context, arg1 *athena.GetQueryRuntimeStatisticsInput, arg2 ...func(*athena.Options)) (*athena.GetQueryRuntimeStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryRuntimeStatistics", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryRuntimeStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryRuntimeStatistics indicates an expected call of GetQueryRuntimeStatistics.
func (mr *MockAthenaClientMockRecorder) GetQueryRuntimeStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryRuntimeStatistics", reflect.TypeOf((*MockAthenaClient)(nil).GetQueryRuntimeStatistics), varargs...)
}

// GetTableMetadata mocks base method.
func (m *MockAthenaClient) GetTableMetadata(arg0 context.Context, arg1 *athena.GetTableMetadataInput, arg2 ...func(*athena.Options)) (*athena.GetTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTableMetadata", varargs...)
	ret0, _ := ret[0].(*athena.GetTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTableMetadata indicates an expected call of GetTableMetadata.
func (mr *MockAthenaClientMockRecorder) GetTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTableMetadata", reflect.TypeOf((*MockAthenaClient)(nil).GetTableMetadata), varargs...)
}

// GetWorkGroup mocks base method.
func (m *MockAthenaClient) GetWorkGroup(arg0 context.Context, arg1 *athena.GetWorkGroupInput, arg2 ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkGroup", varargs...)
	ret0, _ := ret[0].(*athena.GetWorkGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkGroup indicates an expected call of GetWorkGroup.
func (mr *MockAthenaClientMockRecorder) GetWorkGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkGroup", reflect.TypeOf((*MockAthenaClient)(nil).GetWorkGroup), varargs...)
}

// ListDataCatalogs mocks base method.
func (m *MockAthenaClient) ListDataCatalogs(arg0 context.Context, arg1 *athena.ListDataCatalogsInput, arg2 ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDataCatalogs", varargs...)
	ret0, _ := ret[0].(*athena.ListDataCatalogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDataCatalogs indicates an expected call of ListDataCatalogs.
func (mr *MockAthenaClientMockRecorder) ListDataCatalogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDataCatalogs", reflect.TypeOf((*MockAthenaClient)(nil).ListDataCatalogs), varargs...)
}

// ListDatabases mocks base method.
func (m *MockAthenaClient) ListDatabases(arg0 context.Context, arg1 *athena.ListDatabasesInput, arg2 ...func(*athena.Options)) (*athena.ListDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatabases", varargs...)
	ret0, _ := ret[0].(*athena.ListDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatabases indicates an expected call of ListDatabases.
func (mr *MockAthenaClientMockRecorder) ListDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabases", reflect.TypeOf((*MockAthenaClient)(nil).ListDatabases), varargs...)
}

// ListEngineVersions mocks base method.
func (m *MockAthenaClient) ListEngineVersions(arg0 context.Context, arg1 *athena.ListEngineVersionsInput, arg2 ...func(*athena.Options)) (*athena.ListEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEngineVersions", varargs...)
	ret0, _ := ret[0].(*athena.ListEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEngineVersions indicates an expected call of ListEngineVersions.
func (mr *MockAthenaClientMockRecorder) ListEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEngineVersions", reflect.TypeOf((*MockAthenaClient)(nil).ListEngineVersions), varargs...)
}

// ListNamedQueries mocks base method.
func (m *MockAthenaClient) ListNamedQueries(arg0 context.Context, arg1 *athena.ListNamedQueriesInput, arg2 ...func(*athena.Options)) (*athena.ListNamedQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamedQueries", varargs...)
	ret0, _ := ret[0].(*athena.ListNamedQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamedQueries indicates an expected call of ListNamedQueries.
func (mr *MockAthenaClientMockRecorder) ListNamedQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueries", reflect.TypeOf((*MockAthenaClient)(nil).ListNamedQueries), varargs...)
}

// ListPreparedStatements mocks base method.
func (m *MockAthenaClient) ListPreparedStatements(arg0 context.Context, arg1 *athena.ListPreparedStatementsInput, arg2 ...func(*athena.Options)) (*athena.ListPreparedStatementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPreparedStatements", varargs...)
	ret0, _ := ret[0].(*athena.ListPreparedStatementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPreparedStatements indicates an expected call of ListPreparedStatements.
func (mr *MockAthenaClientMockRecorder) ListPreparedStatements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPreparedStatements", reflect.TypeOf((*MockAthenaClient)(nil).ListPreparedStatements), varargs...)
}

// ListQueryExecutions mocks base method.
func (m *MockAthenaClient) ListQueryExecutions(arg0 context.Context, arg1 *athena.ListQueryExecutionsInput, arg2 ...func(*athena.Options)) (*athena.ListQueryExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueryExecutions", varargs...)
	ret0, _ := ret[0].(*athena.ListQueryExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueryExecutions indicates an expected call of ListQueryExecutions.
func (mr *MockAthenaClientMockRecorder) ListQueryExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutions", reflect.TypeOf((*MockAthenaClient)(nil).ListQueryExecutions), varargs...)
}

// ListTableMetadata mocks base method.
func (m *MockAthenaClient) ListTableMetadata(arg0 context.Context, arg1 *athena.ListTableMetadataInput, arg2 ...func(*athena.Options)) (*athena.ListTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableMetadata", varargs...)
	ret0, _ := ret[0].(*athena.ListTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableMetadata indicates an expected call of ListTableMetadata.
func (mr *MockAthenaClientMockRecorder) ListTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableMetadata", reflect.TypeOf((*MockAthenaClient)(nil).ListTableMetadata), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAthenaClient) ListTagsForResource(arg0 context.Context, arg1 *athena.ListTagsForResourceInput, arg2 ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*athena.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAthenaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAthenaClient)(nil).ListTagsForResource), varargs...)
}

// ListWorkGroups mocks base method.
func (m *MockAthenaClient) ListWorkGroups(arg0 context.Context, arg1 *athena.ListWorkGroupsInput, arg2 ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkGroups", varargs...)
	ret0, _ := ret[0].(*athena.ListWorkGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkGroups indicates an expected call of ListWorkGroups.
func (mr *MockAthenaClientMockRecorder) ListWorkGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkGroups", reflect.TypeOf((*MockAthenaClient)(nil).ListWorkGroups), varargs...)
}