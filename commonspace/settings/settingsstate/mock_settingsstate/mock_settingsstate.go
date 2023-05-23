// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/any-sync/commonspace/settings/settingsstate (interfaces: ObjectDeletionState,StateBuilder,ChangeFactory)

// Package mock_settingsstate is a generated GoMock package.
package mock_settingsstate

import (
	reflect "reflect"

	objecttree "github.com/anytypeio/any-sync/commonspace/object/tree/objecttree"
	settingsstate "github.com/anytypeio/any-sync/commonspace/settings/settingsstate"
	gomock "github.com/golang/mock/gomock"
)

// MockObjectDeletionState is a mock of ObjectDeletionState interface.
type MockObjectDeletionState struct {
	ctrl     *gomock.Controller
	recorder *MockObjectDeletionStateMockRecorder
}

// MockObjectDeletionStateMockRecorder is the mock recorder for MockObjectDeletionState.
type MockObjectDeletionStateMockRecorder struct {
	mock *MockObjectDeletionState
}

// NewMockObjectDeletionState creates a new mock instance.
func NewMockObjectDeletionState(ctrl *gomock.Controller) *MockObjectDeletionState {
	mock := &MockObjectDeletionState{ctrl: ctrl}
	mock.recorder = &MockObjectDeletionStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectDeletionState) EXPECT() *MockObjectDeletionStateMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockObjectDeletionState) Add(arg0 map[string]struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockObjectDeletionStateMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockObjectDeletionState)(nil).Add), arg0)
}

// AddObserver mocks base method.
func (m *MockObjectDeletionState) AddObserver(arg0 settingsstate.StateUpdateObserver) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddObserver", arg0)
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockObjectDeletionStateMockRecorder) AddObserver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockObjectDeletionState)(nil).AddObserver), arg0)
}

// Delete mocks base method.
func (m *MockObjectDeletionState) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockObjectDeletionStateMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockObjectDeletionState)(nil).Delete), arg0)
}

// Exists mocks base method.
func (m *MockObjectDeletionState) Exists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockObjectDeletionStateMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockObjectDeletionState)(nil).Exists), arg0)
}

// FilterJoin mocks base method.
func (m *MockObjectDeletionState) FilterJoin(arg0 ...[]string) []string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilterJoin", varargs...)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FilterJoin indicates an expected call of FilterJoin.
func (mr *MockObjectDeletionStateMockRecorder) FilterJoin(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterJoin", reflect.TypeOf((*MockObjectDeletionState)(nil).FilterJoin), arg0...)
}

// GetQueued mocks base method.
func (m *MockObjectDeletionState) GetQueued() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueued")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetQueued indicates an expected call of GetQueued.
func (mr *MockObjectDeletionStateMockRecorder) GetQueued() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueued", reflect.TypeOf((*MockObjectDeletionState)(nil).GetQueued))
}

// MockStateBuilder is a mock of StateBuilder interface.
type MockStateBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockStateBuilderMockRecorder
}

// MockStateBuilderMockRecorder is the mock recorder for MockStateBuilder.
type MockStateBuilderMockRecorder struct {
	mock *MockStateBuilder
}

// NewMockStateBuilder creates a new mock instance.
func NewMockStateBuilder(ctrl *gomock.Controller) *MockStateBuilder {
	mock := &MockStateBuilder{ctrl: ctrl}
	mock.recorder = &MockStateBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateBuilder) EXPECT() *MockStateBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockStateBuilder) Build(arg0 objecttree.ReadableObjectTree, arg1 *settingsstate.State) (*settingsstate.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(*settingsstate.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockStateBuilderMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockStateBuilder)(nil).Build), arg0, arg1)
}

// MockChangeFactory is a mock of ChangeFactory interface.
type MockChangeFactory struct {
	ctrl     *gomock.Controller
	recorder *MockChangeFactoryMockRecorder
}

// MockChangeFactoryMockRecorder is the mock recorder for MockChangeFactory.
type MockChangeFactoryMockRecorder struct {
	mock *MockChangeFactory
}

// NewMockChangeFactory creates a new mock instance.
func NewMockChangeFactory(ctrl *gomock.Controller) *MockChangeFactory {
	mock := &MockChangeFactory{ctrl: ctrl}
	mock.recorder = &MockChangeFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeFactory) EXPECT() *MockChangeFactoryMockRecorder {
	return m.recorder
}

// CreateObjectDeleteChange mocks base method.
func (m *MockChangeFactory) CreateObjectDeleteChange(arg0 string, arg1 *settingsstate.State, arg2 bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjectDeleteChange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjectDeleteChange indicates an expected call of CreateObjectDeleteChange.
func (mr *MockChangeFactoryMockRecorder) CreateObjectDeleteChange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjectDeleteChange", reflect.TypeOf((*MockChangeFactory)(nil).CreateObjectDeleteChange), arg0, arg1, arg2)
}

// CreateSpaceDeleteChange mocks base method.
func (m *MockChangeFactory) CreateSpaceDeleteChange(arg0 string, arg1 *settingsstate.State, arg2 bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpaceDeleteChange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpaceDeleteChange indicates an expected call of CreateSpaceDeleteChange.
func (mr *MockChangeFactoryMockRecorder) CreateSpaceDeleteChange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpaceDeleteChange", reflect.TypeOf((*MockChangeFactory)(nil).CreateSpaceDeleteChange), arg0, arg1, arg2)
}
