// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/listers/core/v1 (interfaces: SecretLister,SecretNamespaceLister)
//
// Generated by this command:
//
//	mockgen k8s.io/client-go/listers/core/v1 SecretLister,SecretNamespaceLister
//

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	v10 "k8s.io/client-go/listers/core/v1"
)

// MockSecretLister is a mock of SecretLister interface.
type MockSecretLister struct {
	ctrl     *gomock.Controller
	recorder *MockSecretListerMockRecorder
}

// MockSecretListerMockRecorder is the mock recorder for MockSecretLister.
type MockSecretListerMockRecorder struct {
	mock *MockSecretLister
}

// NewMockSecretLister creates a new mock instance.
func NewMockSecretLister(ctrl *gomock.Controller) *MockSecretLister {
	mock := &MockSecretLister{ctrl: ctrl}
	mock.recorder = &MockSecretListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretLister) EXPECT() *MockSecretListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSecretLister) List(arg0 labels.Selector) ([]*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretListerMockRecorder) List(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretLister)(nil).List), arg0)
}

// Secrets mocks base method.
func (m *MockSecretLister) Secrets(arg0 string) v10.SecretNamespaceLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets", arg0)
	ret0, _ := ret[0].(v10.SecretNamespaceLister)
	return ret0
}

// Secrets indicates an expected call of Secrets.
func (mr *MockSecretListerMockRecorder) Secrets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockSecretLister)(nil).Secrets), arg0)
}

// MockSecretNamespaceLister is a mock of SecretNamespaceLister interface.
type MockSecretNamespaceLister struct {
	ctrl     *gomock.Controller
	recorder *MockSecretNamespaceListerMockRecorder
}

// MockSecretNamespaceListerMockRecorder is the mock recorder for MockSecretNamespaceLister.
type MockSecretNamespaceListerMockRecorder struct {
	mock *MockSecretNamespaceLister
}

// NewMockSecretNamespaceLister creates a new mock instance.
func NewMockSecretNamespaceLister(ctrl *gomock.Controller) *MockSecretNamespaceLister {
	mock := &MockSecretNamespaceLister{ctrl: ctrl}
	mock.recorder = &MockSecretNamespaceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretNamespaceLister) EXPECT() *MockSecretNamespaceListerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSecretNamespaceLister) Get(arg0 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSecretNamespaceListerMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretNamespaceLister)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockSecretNamespaceLister) List(arg0 labels.Selector) ([]*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretNamespaceListerMockRecorder) List(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretNamespaceLister)(nil).List), arg0)
}
