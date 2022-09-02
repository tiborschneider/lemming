// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openconfig/lemming/dataplane/forwarding/fwdport (interfaces: Port)

// Package mock_fwdport is a generated GoMock package.
package mock_fwdport

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fwdaction "github.com/openconfig/lemming/dataplane/forwarding/fwdaction"
	fwdattribute "github.com/openconfig/lemming/dataplane/forwarding/infra/fwdattribute"
	fwdobject "github.com/openconfig/lemming/dataplane/forwarding/infra/fwdobject"
	fwdpacket "github.com/openconfig/lemming/dataplane/forwarding/infra/fwdpacket"
	forwarding "github.com/openconfig/lemming/proto/forwarding"
)

// MockPort is a mock of Port interface.
type MockPort struct {
	ctrl     *gomock.Controller
	recorder *MockPortMockRecorder
}

// MockPortMockRecorder is the mock recorder for MockPort.
type MockPortMockRecorder struct {
	mock *MockPort
}

// NewMockPort creates a new mock instance.
func NewMockPort(ctrl *gomock.Controller) *MockPort {
	mock := &MockPort{ctrl: ctrl}
	mock.recorder = &MockPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPort) EXPECT() *MockPortMockRecorder {
	return m.recorder
}

// Acquire mocks base method.
func (m *MockPort) Acquire() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acquire")
	ret0, _ := ret[0].(error)
	return ret0
}

// Acquire indicates an expected call of Acquire.
func (mr *MockPortMockRecorder) Acquire() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockPort)(nil).Acquire))
}

// Actions mocks base method.
func (m *MockPort) Actions(arg0 forwarding.PortAction) fwdaction.Actions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions", arg0)
	ret0, _ := ret[0].(fwdaction.Actions)
	return ret0
}

// Actions indicates an expected call of Actions.
func (mr *MockPortMockRecorder) Actions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockPort)(nil).Actions), arg0)
}

// Attributes mocks base method.
func (m *MockPort) Attributes() fwdattribute.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attributes")
	ret0, _ := ret[0].(fwdattribute.Set)
	return ret0
}

// Attributes indicates an expected call of Attributes.
func (mr *MockPortMockRecorder) Attributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attributes", reflect.TypeOf((*MockPort)(nil).Attributes))
}

// Counters mocks base method.
func (m *MockPort) Counters() map[forwarding.CounterId]fwdobject.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Counters")
	ret0, _ := ret[0].(map[forwarding.CounterId]fwdobject.Counter)
	return ret0
}

// Counters indicates an expected call of Counters.
func (mr *MockPortMockRecorder) Counters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counters", reflect.TypeOf((*MockPort)(nil).Counters))
}

// ID mocks base method.
func (m *MockPort) ID() fwdobject.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(fwdobject.ID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockPortMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockPort)(nil).ID))
}

// Increment mocks base method.
func (m *MockPort) Increment(arg0 forwarding.CounterId, arg1 uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Increment", arg0, arg1)
}

// Increment indicates an expected call of Increment.
func (mr *MockPortMockRecorder) Increment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockPort)(nil).Increment), arg0, arg1)
}

// Init mocks base method.
func (m *MockPort) Init(arg0 fwdobject.ID, arg1 fwdobject.NID, arg2 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0, arg1, arg2)
}

// Init indicates an expected call of Init.
func (mr *MockPortMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockPort)(nil).Init), arg0, arg1, arg2)
}

// NID mocks base method.
func (m *MockPort) NID() fwdobject.NID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NID")
	ret0, _ := ret[0].(fwdobject.NID)
	return ret0
}

// NID indicates an expected call of NID.
func (mr *MockPortMockRecorder) NID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NID", reflect.TypeOf((*MockPort)(nil).NID))
}

// Release mocks base method.
func (m *MockPort) Release(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release.
func (mr *MockPortMockRecorder) Release(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockPort)(nil).Release), arg0)
}

// State mocks base method.
func (m *MockPort) State(arg0 *forwarding.PortInfo) (forwarding.PortStateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State", arg0)
	ret0, _ := ret[0].(forwarding.PortStateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockPortMockRecorder) State(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockPort)(nil).State), arg0)
}

// String mocks base method.
func (m *MockPort) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockPortMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockPort)(nil).String))
}

// Update mocks base method.
func (m *MockPort) Update(arg0 *forwarding.PortUpdateDesc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPortMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPort)(nil).Update), arg0)
}

// Write mocks base method.
func (m *MockPort) Write(arg0 fwdpacket.Packet) (fwdaction.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(fwdaction.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockPortMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockPort)(nil).Write), arg0)
}
