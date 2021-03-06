// Code generated by MockGen. DO NOT EDIT.
// Source: ../generator.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGenerator is a mock of Generator interface
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// RandomTitle mocks base method
func (m *MockGenerator) RandomTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomTitle indicates an expected call of RandomTitle
func (mr *MockGeneratorMockRecorder) RandomTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomTitle", reflect.TypeOf((*MockGenerator)(nil).RandomTitle))
}
