// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/config (interfaces: Dash)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	cluster "github.com/vmware-tanzu/octant/internal/cluster"
	config "github.com/vmware-tanzu/octant/internal/config"
	errors "github.com/vmware-tanzu/octant/internal/errors"
	module "github.com/vmware-tanzu/octant/internal/module"
	portforward "github.com/vmware-tanzu/octant/internal/portforward"
	log "github.com/vmware-tanzu/octant/pkg/log"
	plugin "github.com/vmware-tanzu/octant/pkg/plugin"
	store "github.com/vmware-tanzu/octant/pkg/store"
)

// MockDash is a mock of Dash interface
type MockDash struct {
	ctrl     *gomock.Controller
	recorder *MockDashMockRecorder
}

// MockDashMockRecorder is the mock recorder for MockDash
type MockDashMockRecorder struct {
	mock *MockDash
}

// NewMockDash creates a new mock instance
func NewMockDash(ctrl *gomock.Controller) *MockDash {
	mock := &MockDash{ctrl: ctrl}
	mock.recorder = &MockDashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDash) EXPECT() *MockDashMockRecorder {
	return m.recorder
}

// BuildInfo mocks base method
func (m *MockDash) BuildInfo() (string, string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildInfo")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	return ret0, ret1, ret2
}

// BuildInfo indicates an expected call of BuildInfo
func (mr *MockDashMockRecorder) BuildInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildInfo", reflect.TypeOf((*MockDash)(nil).BuildInfo))
}

// CRDWatcher mocks base method
func (m *MockDash) CRDWatcher() config.CRDWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CRDWatcher")
	ret0, _ := ret[0].(config.CRDWatcher)
	return ret0
}

// CRDWatcher indicates an expected call of CRDWatcher
func (mr *MockDashMockRecorder) CRDWatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CRDWatcher", reflect.TypeOf((*MockDash)(nil).CRDWatcher))
}

// ClusterClient mocks base method
func (m *MockDash) ClusterClient() cluster.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterClient")
	ret0, _ := ret[0].(cluster.ClientInterface)
	return ret0
}

// ClusterClient indicates an expected call of ClusterClient
func (mr *MockDashMockRecorder) ClusterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterClient", reflect.TypeOf((*MockDash)(nil).ClusterClient))
}

// ContextName mocks base method
func (m *MockDash) ContextName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContextName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContextName indicates an expected call of ContextName
func (mr *MockDashMockRecorder) ContextName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextName", reflect.TypeOf((*MockDash)(nil).ContextName))
}

// DefaultNamespace mocks base method
func (m *MockDash) DefaultNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultNamespace indicates an expected call of DefaultNamespace
func (mr *MockDashMockRecorder) DefaultNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultNamespace", reflect.TypeOf((*MockDash)(nil).DefaultNamespace))
}

// ErrorStore mocks base method
func (m *MockDash) ErrorStore() errors.ErrorStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorStore")
	ret0, _ := ret[0].(errors.ErrorStore)
	return ret0
}

// ErrorStore indicates an expected call of ErrorStore
func (mr *MockDashMockRecorder) ErrorStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorStore", reflect.TypeOf((*MockDash)(nil).ErrorStore))
}

// KubeConfigPath mocks base method
func (m *MockDash) KubeConfigPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeConfigPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// KubeConfigPath indicates an expected call of KubeConfigPath
func (mr *MockDashMockRecorder) KubeConfigPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeConfigPath", reflect.TypeOf((*MockDash)(nil).KubeConfigPath))
}

// Logger mocks base method
func (m *MockDash) Logger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockDashMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockDash)(nil).Logger))
}

// ModuleManager mocks base method
func (m *MockDash) ModuleManager() module.ManagerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModuleManager")
	ret0, _ := ret[0].(module.ManagerInterface)
	return ret0
}

// ModuleManager indicates an expected call of ModuleManager
func (mr *MockDashMockRecorder) ModuleManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModuleManager", reflect.TypeOf((*MockDash)(nil).ModuleManager))
}

// ObjectPath mocks base method
func (m *MockDash) ObjectPath(arg0, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectPath", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPath indicates an expected call of ObjectPath
func (mr *MockDashMockRecorder) ObjectPath(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPath", reflect.TypeOf((*MockDash)(nil).ObjectPath), arg0, arg1, arg2, arg3)
}

// ObjectStore mocks base method
func (m *MockDash) ObjectStore() store.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStore")
	ret0, _ := ret[0].(store.Store)
	return ret0
}

// ObjectStore indicates an expected call of ObjectStore
func (mr *MockDashMockRecorder) ObjectStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStore", reflect.TypeOf((*MockDash)(nil).ObjectStore))
}

// PluginManager mocks base method
func (m *MockDash) PluginManager() plugin.ManagerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluginManager")
	ret0, _ := ret[0].(plugin.ManagerInterface)
	return ret0
}

// PluginManager indicates an expected call of PluginManager
func (mr *MockDashMockRecorder) PluginManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluginManager", reflect.TypeOf((*MockDash)(nil).PluginManager))
}

// PortForwarder mocks base method
func (m *MockDash) PortForwarder() portforward.PortForwarder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PortForwarder")
	ret0, _ := ret[0].(portforward.PortForwarder)
	return ret0
}

// PortForwarder indicates an expected call of PortForwarder
func (mr *MockDashMockRecorder) PortForwarder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortForwarder", reflect.TypeOf((*MockDash)(nil).PortForwarder))
}

// UseContext mocks base method
func (m *MockDash) UseContext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseContext indicates an expected call of UseContext
func (mr *MockDashMockRecorder) UseContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseContext", reflect.TypeOf((*MockDash)(nil).UseContext), arg0, arg1)
}

// Validate mocks base method
func (m *MockDash) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockDashMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDash)(nil).Validate))
}
