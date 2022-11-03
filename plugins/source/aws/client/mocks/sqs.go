// Code generated by MockGen. DO NOT EDIT.
// Source: sqs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	gomock "github.com/golang/mock/gomock"
)

// MockSqsClient is a mock of SqsClient interface.
type MockSqsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSqsClientMockRecorder
}

// MockSqsClientMockRecorder is the mock recorder for MockSqsClient.
type MockSqsClientMockRecorder struct {
	mock *MockSqsClient
}

// NewMockSqsClient creates a new mock instance.
func NewMockSqsClient(ctrl *gomock.Controller) *MockSqsClient {
	mock := &MockSqsClient{ctrl: ctrl}
	mock.recorder = &MockSqsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqsClient) EXPECT() *MockSqsClientMockRecorder {
	return m.recorder
}

// GetQueueAttributes mocks base method.
func (m *MockSqsClient) GetQueueAttributes(arg0 context.Context, arg1 *sqs.GetQueueAttributesInput, arg2 ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueueAttributes", varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueAttributes indicates an expected call of GetQueueAttributes.
func (mr *MockSqsClientMockRecorder) GetQueueAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAttributes", reflect.TypeOf((*MockSqsClient)(nil).GetQueueAttributes), varargs...)
}

// GetQueueUrl mocks base method.
func (m *MockSqsClient) GetQueueUrl(arg0 context.Context, arg1 *sqs.GetQueueUrlInput, arg2 ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueueUrl", varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueUrl indicates an expected call of GetQueueUrl.
func (mr *MockSqsClientMockRecorder) GetQueueUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueUrl", reflect.TypeOf((*MockSqsClient)(nil).GetQueueUrl), varargs...)
}

// ListDeadLetterSourceQueues mocks base method.
func (m *MockSqsClient) ListDeadLetterSourceQueues(arg0 context.Context, arg1 *sqs.ListDeadLetterSourceQueuesInput, arg2 ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeadLetterSourceQueues", varargs...)
	ret0, _ := ret[0].(*sqs.ListDeadLetterSourceQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeadLetterSourceQueues indicates an expected call of ListDeadLetterSourceQueues.
func (mr *MockSqsClientMockRecorder) ListDeadLetterSourceQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeadLetterSourceQueues", reflect.TypeOf((*MockSqsClient)(nil).ListDeadLetterSourceQueues), varargs...)
}

// ListQueueTags mocks base method.
func (m *MockSqsClient) ListQueueTags(arg0 context.Context, arg1 *sqs.ListQueueTagsInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueueTags", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueueTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueueTags indicates an expected call of ListQueueTags.
func (mr *MockSqsClientMockRecorder) ListQueueTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueueTags", reflect.TypeOf((*MockSqsClient)(nil).ListQueueTags), varargs...)
}

// ListQueues mocks base method.
func (m *MockSqsClient) ListQueues(arg0 context.Context, arg1 *sqs.ListQueuesInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueues", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueues indicates an expected call of ListQueues.
func (mr *MockSqsClientMockRecorder) ListQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueues", reflect.TypeOf((*MockSqsClient)(nil).ListQueues), varargs...)
}