// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mymmrac/telego/telegoapi (interfaces: Caller)
//
// Generated by this command:
//
//	mockgen -typed -package mock -destination=mock/caller.go github.com/mymmrac/telego/telegoapi Caller
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	telegoapi "github.com/mymmrac/telego/telegoapi"
	gomock "go.uber.org/mock/gomock"
)

// MockCaller is a mock of Caller interface.
type MockCaller struct {
	ctrl     *gomock.Controller
	recorder *MockCallerMockRecorder
}

// MockCallerMockRecorder is the mock recorder for MockCaller.
type MockCallerMockRecorder struct {
	mock *MockCaller
}

// NewMockCaller creates a new mock instance.
func NewMockCaller(ctrl *gomock.Controller) *MockCaller {
	mock := &MockCaller{ctrl: ctrl}
	mock.recorder = &MockCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaller) EXPECT() *MockCallerMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockCaller) Call(arg0 string, arg1 *telegoapi.RequestData) (*telegoapi.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(*telegoapi.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockCallerMockRecorder) Call(arg0, arg1 any) *CallerCallCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockCaller)(nil).Call), arg0, arg1)
	return &CallerCallCall{Call: call}
}

// CallerCallCall wrap *gomock.Call
type CallerCallCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CallerCallCall) Return(arg0 *telegoapi.Response, arg1 error) *CallerCallCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CallerCallCall) Do(f func(string, *telegoapi.RequestData) (*telegoapi.Response, error)) *CallerCallCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CallerCallCall) DoAndReturn(f func(string, *telegoapi.RequestData) (*telegoapi.Response, error)) *CallerCallCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
