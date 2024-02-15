// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/coordinator/coordinatorclient (interfaces: CoordinatorClient)
//
// Generated by this command:
//
//	mockgen -destination mock_coordinatorclient/mock_coordinatorclient.go github.com/anyproto/any-sync/coordinator/coordinatorclient CoordinatorClient
//

// Package mock_coordinatorclient is a generated GoMock package.
package mock_coordinatorclient

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	coordinatorclient "github.com/anyproto/any-sync/coordinator/coordinatorclient"
	coordinatorproto "github.com/anyproto/any-sync/coordinator/coordinatorproto"
	identityrepoproto "github.com/anyproto/any-sync/identityrepo/identityrepoproto"
	gomock "go.uber.org/mock/gomock"
)

// MockCoordinatorClient is a mock of CoordinatorClient interface.
type MockCoordinatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinatorClientMockRecorder
}

// MockCoordinatorClientMockRecorder is the mock recorder for MockCoordinatorClient.
type MockCoordinatorClientMockRecorder struct {
	mock *MockCoordinatorClient
}

// NewMockCoordinatorClient creates a new mock instance.
func NewMockCoordinatorClient(ctrl *gomock.Controller) *MockCoordinatorClient {
	mock := &MockCoordinatorClient{ctrl: ctrl}
	mock.recorder = &MockCoordinatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoordinatorClient) EXPECT() *MockCoordinatorClientMockRecorder {
	return m.recorder
}

// AccountDelete mocks base method.
func (m *MockCoordinatorClient) AccountDelete(arg0 context.Context, arg1 *coordinatorproto.DeletionConfirmPayloadWithSignature) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountDelete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountDelete indicates an expected call of AccountDelete.
func (mr *MockCoordinatorClientMockRecorder) AccountDelete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountDelete", reflect.TypeOf((*MockCoordinatorClient)(nil).AccountDelete), arg0, arg1)
}

// AccountRevertDeletion mocks base method.
func (m *MockCoordinatorClient) AccountRevertDeletion(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountRevertDeletion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AccountRevertDeletion indicates an expected call of AccountRevertDeletion.
func (mr *MockCoordinatorClientMockRecorder) AccountRevertDeletion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountRevertDeletion", reflect.TypeOf((*MockCoordinatorClient)(nil).AccountRevertDeletion), arg0)
}

// AclAddRecord mocks base method.
func (m *MockCoordinatorClient) AclAddRecord(arg0 context.Context, arg1 string, arg2 *consensusproto.RawRecord) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclAddRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclAddRecord indicates an expected call of AclAddRecord.
func (mr *MockCoordinatorClientMockRecorder) AclAddRecord(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclAddRecord", reflect.TypeOf((*MockCoordinatorClient)(nil).AclAddRecord), arg0, arg1, arg2)
}

// AclGetRecords mocks base method.
func (m *MockCoordinatorClient) AclGetRecords(arg0 context.Context, arg1, arg2 string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclGetRecords", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclGetRecords indicates an expected call of AclGetRecords.
func (mr *MockCoordinatorClientMockRecorder) AclGetRecords(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclGetRecords", reflect.TypeOf((*MockCoordinatorClient)(nil).AclGetRecords), arg0, arg1, arg2)
}

// DeletionLog mocks base method.
func (m *MockCoordinatorClient) DeletionLog(arg0 context.Context, arg1 string, arg2 int) ([]*coordinatorproto.DeletionLogRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletionLog", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*coordinatorproto.DeletionLogRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletionLog indicates an expected call of DeletionLog.
func (mr *MockCoordinatorClientMockRecorder) DeletionLog(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletionLog", reflect.TypeOf((*MockCoordinatorClient)(nil).DeletionLog), arg0, arg1, arg2)
}

// FileLimitCheck mocks base method.
func (m *MockCoordinatorClient) FileLimitCheck(arg0 context.Context, arg1 string, arg2 []byte) (*coordinatorproto.FileLimitCheckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileLimitCheck", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coordinatorproto.FileLimitCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileLimitCheck indicates an expected call of FileLimitCheck.
func (mr *MockCoordinatorClientMockRecorder) FileLimitCheck(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileLimitCheck", reflect.TypeOf((*MockCoordinatorClient)(nil).FileLimitCheck), arg0, arg1, arg2)
}

// IdentityRepoGet mocks base method.
func (m *MockCoordinatorClient) IdentityRepoGet(arg0 context.Context, arg1, arg2 []string) ([]*identityrepoproto.DataWithIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityRepoGet", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*identityrepoproto.DataWithIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdentityRepoGet indicates an expected call of IdentityRepoGet.
func (mr *MockCoordinatorClientMockRecorder) IdentityRepoGet(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityRepoGet", reflect.TypeOf((*MockCoordinatorClient)(nil).IdentityRepoGet), arg0, arg1, arg2)
}

// IdentityRepoPut mocks base method.
func (m *MockCoordinatorClient) IdentityRepoPut(arg0 context.Context, arg1 string, arg2 []*identityrepoproto.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityRepoPut", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// IdentityRepoPut indicates an expected call of IdentityRepoPut.
func (mr *MockCoordinatorClientMockRecorder) IdentityRepoPut(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityRepoPut", reflect.TypeOf((*MockCoordinatorClient)(nil).IdentityRepoPut), arg0, arg1, arg2)
}

// Init mocks base method.
func (m *MockCoordinatorClient) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockCoordinatorClientMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockCoordinatorClient)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockCoordinatorClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCoordinatorClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCoordinatorClient)(nil).Name))
}

// NetworkConfiguration mocks base method.
func (m *MockCoordinatorClient) NetworkConfiguration(arg0 context.Context, arg1 string) (*coordinatorproto.NetworkConfigurationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*coordinatorproto.NetworkConfigurationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkConfiguration indicates an expected call of NetworkConfiguration.
func (mr *MockCoordinatorClientMockRecorder) NetworkConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkConfiguration", reflect.TypeOf((*MockCoordinatorClient)(nil).NetworkConfiguration), arg0, arg1)
}

// SpaceDelete mocks base method.
func (m *MockCoordinatorClient) SpaceDelete(arg0 context.Context, arg1 string, arg2 *coordinatorproto.DeletionConfirmPayloadWithSignature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SpaceDelete indicates an expected call of SpaceDelete.
func (mr *MockCoordinatorClientMockRecorder) SpaceDelete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceDelete", reflect.TypeOf((*MockCoordinatorClient)(nil).SpaceDelete), arg0, arg1, arg2)
}

// SpaceSign mocks base method.
func (m *MockCoordinatorClient) SpaceSign(arg0 context.Context, arg1 coordinatorclient.SpaceSignPayload) (*coordinatorproto.SpaceReceiptWithSignature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceSign", arg0, arg1)
	ret0, _ := ret[0].(*coordinatorproto.SpaceReceiptWithSignature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceSign indicates an expected call of SpaceSign.
func (mr *MockCoordinatorClientMockRecorder) SpaceSign(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceSign", reflect.TypeOf((*MockCoordinatorClient)(nil).SpaceSign), arg0, arg1)
}

// StatusCheck mocks base method.
func (m *MockCoordinatorClient) StatusCheck(arg0 context.Context, arg1 string) (*coordinatorproto.SpaceStatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusCheck", arg0, arg1)
	ret0, _ := ret[0].(*coordinatorproto.SpaceStatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusCheck indicates an expected call of StatusCheck.
func (mr *MockCoordinatorClientMockRecorder) StatusCheck(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusCheck", reflect.TypeOf((*MockCoordinatorClient)(nil).StatusCheck), arg0, arg1)
}

// StatusCheckMany mocks base method.
func (m *MockCoordinatorClient) StatusCheckMany(arg0 context.Context, arg1 []string) ([]*coordinatorproto.SpaceStatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusCheckMany", arg0, arg1)
	ret0, _ := ret[0].([]*coordinatorproto.SpaceStatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusCheckMany indicates an expected call of StatusCheckMany.
func (mr *MockCoordinatorClientMockRecorder) StatusCheckMany(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusCheckMany", reflect.TypeOf((*MockCoordinatorClient)(nil).StatusCheckMany), arg0, arg1)
}
