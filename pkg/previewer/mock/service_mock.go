// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_previewer is a generated GoMock package.
package mock_previewer

import (
	previewer "image-previewer/pkg/previewer"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Fill mocks base method.
func (m *MockService) Fill(params *previewer.FillParams) (*previewer.FillResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fill", params)
	ret0, _ := ret[0].(*previewer.FillResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fill indicates an expected call of Fill.
func (mr *MockServiceMockRecorder) Fill(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fill", reflect.TypeOf((*MockService)(nil).Fill), params)
}
