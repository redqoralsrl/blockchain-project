package mock

import (
	"go.uber.org/mock/gomock"
	"reflect"
	"time"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// AddOrUpdateItem mocks base method.
func (m *MockUseCase) AddOrUpdateItem(key string, value any, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateItem", key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateItem indicates an expected call of AddOrUpdateItem.
func (mr *MockUseCaseMockRecorder) AddOrUpdateItem(key, value, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateItem", reflect.TypeOf((*MockUseCase)(nil).AddOrUpdateItem), key, value, expiration)
}

// ExistsItem mocks base method.
func (m *MockUseCase) ExistsItem(key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsItem", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsItem indicates an expected call of ExistsItem.
func (mr *MockUseCaseMockRecorder) ExistsItem(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsItem", reflect.TypeOf((*MockUseCase)(nil).ExistsItem), key)
}

// FindKeys mocks base method.
func (m *MockUseCase) FindKeys(key string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKeys", key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKeys indicates an expected call of FindKeys.
func (mr *MockUseCaseMockRecorder) FindKeys(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKeys", reflect.TypeOf((*MockUseCase)(nil).FindKeys), key)
}

// InvalidateAll mocks base method.
func (m *MockUseCase) InvalidateAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateAll indicates an expected call of InvalidateAll.
func (mr *MockUseCaseMockRecorder) InvalidateAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateAll", reflect.TypeOf((*MockUseCase)(nil).InvalidateAll))
}

// InvalidateItem mocks base method.
func (m *MockUseCase) InvalidateItem(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateItem", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateItem indicates an expected call of InvalidateItem.
func (mr *MockUseCaseMockRecorder) InvalidateItem(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateItem", reflect.TypeOf((*MockUseCase)(nil).InvalidateItem), key)
}

// RetrieveItem mocks base method.
func (m *MockUseCase) RetrieveItem(key string, dest any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveItem", key, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetrieveItem indicates an expected call of RetrieveItem.
func (mr *MockUseCaseMockRecorder) RetrieveItem(key, dest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveItem", reflect.TypeOf((*MockUseCase)(nil).RetrieveItem), key, dest)
}

// RetrieveMultiItem mocks base method.
func (m *MockUseCase) RetrieveMultiItem(keys []string, dest any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMultiItem", keys, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetrieveMultiItem indicates an expected call of RetrieveMultiItem.
func (mr *MockUseCaseMockRecorder) RetrieveMultiItem(keys, dest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMultiItem", reflect.TypeOf((*MockUseCase)(nil).RetrieveMultiItem), keys, dest)
}
