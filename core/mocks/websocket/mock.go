// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/nitric/core/pkg/plugins/websocket (interfaces: WebsocketService)

// Package mock_websocket is a generated GoMock package.
package mock_websocket

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWebsocketService is a mock of WebsocketService interface.
type MockWebsocketService struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketServiceMockRecorder
}

// MockWebsocketServiceMockRecorder is the mock recorder for MockWebsocketService.
type MockWebsocketServiceMockRecorder struct {
	mock *MockWebsocketService
}

// NewMockWebsocketService creates a new mock instance.
func NewMockWebsocketService(ctrl *gomock.Controller) *MockWebsocketService {
	mock := &MockWebsocketService{ctrl: ctrl}
	mock.recorder = &MockWebsocketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebsocketService) EXPECT() *MockWebsocketServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockWebsocketService) Close(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockWebsocketServiceMockRecorder) Close(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWebsocketService)(nil).Close), arg0, arg1, arg2)
}

// Send mocks base method.
func (m *MockWebsocketService) Send(arg0 context.Context, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockWebsocketServiceMockRecorder) Send(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockWebsocketService)(nil).Send), arg0, arg1, arg2, arg3)
}
