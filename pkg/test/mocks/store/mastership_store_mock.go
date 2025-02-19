// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/store/mastership/store.go

// Package store is a generated GoMock package.
package store

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/onosproject/onos-config/pkg/device"
	mastership "github.com/onosproject/onos-config/pkg/store/mastership"
	cluster "github.com/onosproject/onos-lib-go/pkg/cluster"
	reflect "reflect"
)

// MockMastershipStore is a mock of Store interface
type MockMastershipStore struct {
	ctrl     *gomock.Controller
	recorder *MockMastershipStoreMockRecorder
}

// MockMastershipStoreMockRecorder is the mock recorder for MockMastershipStore
type MockMastershipStoreMockRecorder struct {
	mock *MockMastershipStore
}

// NewMockMastershipStore creates a new mock instance
func NewMockMastershipStore(ctrl *gomock.Controller) *MockMastershipStore {
	mock := &MockMastershipStore{ctrl: ctrl}
	mock.recorder = &MockMastershipStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMastershipStore) EXPECT() *MockMastershipStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMastershipStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMastershipStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMastershipStore)(nil).Close))
}

// NodeID mocks base method
func (m *MockMastershipStore) NodeID() cluster.NodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(cluster.NodeID)
	return ret0
}

// NodeID indicates an expected call of NodeID
func (mr *MockMastershipStoreMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockMastershipStore)(nil).NodeID))
}

// GetMastership mocks base method
func (m *MockMastershipStore) GetMastership(id device.ID) (*mastership.Mastership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMastership", id)
	ret0, _ := ret[0].(*mastership.Mastership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMastership indicates an expected call of GetMastership
func (mr *MockMastershipStoreMockRecorder) GetMastership(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMastership", reflect.TypeOf((*MockMastershipStore)(nil).GetMastership), id)
}

// Watch mocks base method
func (m *MockMastershipStore) Watch(arg0 device.ID, arg1 chan<- mastership.Mastership) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockMastershipStoreMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMastershipStore)(nil).Watch), arg0, arg1)
}
