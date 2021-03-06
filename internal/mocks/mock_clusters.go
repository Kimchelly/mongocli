// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ClusterLister,ClusterDescriber,ClusterCreator,ClusterDeleter,ClusterUpdater,ClusterStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockClusterLister is a mock of ClusterLister interface
type MockClusterLister struct {
	ctrl     *gomock.Controller
	recorder *MockClusterListerMockRecorder
}

// MockClusterListerMockRecorder is the mock recorder for MockClusterLister
type MockClusterListerMockRecorder struct {
	mock *MockClusterLister
}

// NewMockClusterLister creates a new mock instance
func NewMockClusterLister(ctrl *gomock.Controller) *MockClusterLister {
	mock := &MockClusterLister{ctrl: ctrl}
	mock.recorder = &MockClusterListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterLister) EXPECT() *MockClusterListerMockRecorder {
	return m.recorder
}

// ProjectClusters mocks base method
func (m *MockClusterLister) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters
func (mr *MockClusterListerMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockClusterLister)(nil).ProjectClusters), arg0, arg1)
}

// MockClusterDescriber is a mock of ClusterDescriber interface
type MockClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDescriberMockRecorder
}

// MockClusterDescriberMockRecorder is the mock recorder for MockClusterDescriber
type MockClusterDescriberMockRecorder struct {
	mock *MockClusterDescriber
}

// NewMockClusterDescriber creates a new mock instance
func NewMockClusterDescriber(ctrl *gomock.Controller) *MockClusterDescriber {
	mock := &MockClusterDescriber{ctrl: ctrl}
	mock.recorder = &MockClusterDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterDescriber) EXPECT() *MockClusterDescriberMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockClusterDescriber) Cluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockClusterDescriberMockRecorder) Cluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockClusterDescriber)(nil).Cluster), arg0, arg1)
}

// MockClusterCreator is a mock of ClusterCreator interface
type MockClusterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCreatorMockRecorder
}

// MockClusterCreatorMockRecorder is the mock recorder for MockClusterCreator
type MockClusterCreatorMockRecorder struct {
	mock *MockClusterCreator
}

// NewMockClusterCreator creates a new mock instance
func NewMockClusterCreator(ctrl *gomock.Controller) *MockClusterCreator {
	mock := &MockClusterCreator{ctrl: ctrl}
	mock.recorder = &MockClusterCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterCreator) EXPECT() *MockClusterCreatorMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockClusterCreator) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockClusterCreatorMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterCreator)(nil).CreateCluster), arg0)
}

// MockClusterDeleter is a mock of ClusterDeleter interface
type MockClusterDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDeleterMockRecorder
}

// MockClusterDeleterMockRecorder is the mock recorder for MockClusterDeleter
type MockClusterDeleterMockRecorder struct {
	mock *MockClusterDeleter
}

// NewMockClusterDeleter creates a new mock instance
func NewMockClusterDeleter(ctrl *gomock.Controller) *MockClusterDeleter {
	mock := &MockClusterDeleter{ctrl: ctrl}
	mock.recorder = &MockClusterDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterDeleter) EXPECT() *MockClusterDeleterMockRecorder {
	return m.recorder
}

// DeleteCluster mocks base method
func (m *MockClusterDeleter) DeleteCluster(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockClusterDeleterMockRecorder) DeleteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterDeleter)(nil).DeleteCluster), arg0, arg1)
}

// MockClusterUpdater is a mock of ClusterUpdater interface
type MockClusterUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockClusterUpdaterMockRecorder
}

// MockClusterUpdaterMockRecorder is the mock recorder for MockClusterUpdater
type MockClusterUpdaterMockRecorder struct {
	mock *MockClusterUpdater
}

// NewMockClusterUpdater creates a new mock instance
func NewMockClusterUpdater(ctrl *gomock.Controller) *MockClusterUpdater {
	mock := &MockClusterUpdater{ctrl: ctrl}
	mock.recorder = &MockClusterUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterUpdater) EXPECT() *MockClusterUpdaterMockRecorder {
	return m.recorder
}

// UpdateCluster mocks base method
func (m *MockClusterUpdater) UpdateCluster(arg0, arg1 string, arg2 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockClusterUpdaterMockRecorder) UpdateCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterUpdater)(nil).UpdateCluster), arg0, arg1, arg2)
}

// MockClusterStore is a mock of ClusterStore interface
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockClusterStore) Cluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockClusterStoreMockRecorder) Cluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockClusterStore)(nil).Cluster), arg0, arg1)
}

// CreateCluster mocks base method
func (m *MockClusterStore) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockClusterStoreMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterStore)(nil).CreateCluster), arg0)
}

// DeleteCluster mocks base method
func (m *MockClusterStore) DeleteCluster(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockClusterStoreMockRecorder) DeleteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterStore)(nil).DeleteCluster), arg0, arg1)
}

// ProjectClusters mocks base method
func (m *MockClusterStore) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters
func (mr *MockClusterStoreMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockClusterStore)(nil).ProjectClusters), arg0, arg1)
}

// UpdateCluster mocks base method
func (m *MockClusterStore) UpdateCluster(arg0, arg1 string, arg2 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockClusterStoreMockRecorder) UpdateCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterStore)(nil).UpdateCluster), arg0, arg1, arg2)
}
