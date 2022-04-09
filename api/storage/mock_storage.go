package storage

import (
	"reflect"

	"github.com/getzion/relay/api"
	"github.com/golang/mock/gomock"
)

type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *mockStorageMockRecorder
}

// mockStorageMockRecorder is the mock recorder for MockStorage.
type mockStorageMockRecorder struct {
	mock *MockStorage
}

func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &mockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *mockStorageMockRecorder {
	return m.recorder
}

func (m *MockStorage) GetCommunities() ([]api.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommunities")
	ret0, _ := ret[0].([]api.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) GetCommunityByZid(arg0 string) (*api.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommunityByZid", arg0)
	ret0, _ := ret[0].(*api.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) InsertCommunity(arg0 *api.Community) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCommunity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) AddUserToCommunity(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToCommunity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) RemoveUserToCommunity(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserToCommunity", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) GetUsers() ([]api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) GetUserByDid(arg0 string) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByDid", arg0)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) GetUserByUsername(arg0 string) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) InsertUser(arg0 *api.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) GetConversations() ([]api.Conversation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversations")
	ret0, _ := ret[0].([]api.Conversation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) InsertConversation(arg0 *api.Conversation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertConversation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) GetComments() ([]api.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments")
	ret0, _ := ret[0].([]api.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) InsertComment(arg0 *api.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertComment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockStorage) GetPayments() ([]api.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayments")
	ret0, _ := ret[0].([]api.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockStorage) InsertPayment(arg0 *api.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPayment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *mockStorageMockRecorder) GetCommunities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommunities", reflect.TypeOf((*MockStorage)(nil).GetCommunities))
}

func (mr *mockStorageMockRecorder) GetCommunityByZid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommunityByZid", reflect.TypeOf((*MockStorage)(nil).GetCommunityByZid), arg0)
}

func (mr *mockStorageMockRecorder) InsertCommunity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCommunity", reflect.TypeOf((*MockStorage)(nil).InsertCommunity), arg0)
}

func (mr *mockStorageMockRecorder) AddUserToCommunity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToCommunity", reflect.TypeOf((*MockStorage)(nil).AddUserToCommunity), arg0, arg1)
}

func (mr *mockStorageMockRecorder) RemoveUserToCommunity(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserToCommunity", reflect.TypeOf((*MockStorage)(nil).RemoveUserToCommunity), arg0, arg1, arg2)
}

func (mr *mockStorageMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStorage)(nil).GetUsers))
}

func (mr *mockStorageMockRecorder) GetUserByDid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByDid", reflect.TypeOf((*MockStorage)(nil).GetUserByDid), arg0)
}

func (mr *mockStorageMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStorage)(nil).GetUserByUsername), arg0)
}

func (mr *mockStorageMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockStorage)(nil).InsertUser), arg0)
}

func (mr *mockStorageMockRecorder) GetConversations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversations", reflect.TypeOf((*MockStorage)(nil).GetConversations))
}

func (mr *mockStorageMockRecorder) InsertConversation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertConversation", reflect.TypeOf((*MockStorage)(nil).InsertConversation), arg0)
}

func (mr *mockStorageMockRecorder) GetComments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockStorage)(nil).GetComments))
}

func (mr *mockStorageMockRecorder) InsertComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertComment", reflect.TypeOf((*MockStorage)(nil).InsertComment), arg0)
}

func (mr *mockStorageMockRecorder) GetPayments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayments", reflect.TypeOf((*MockStorage)(nil).GetPayments))
}

func (mr *mockStorageMockRecorder) InsertPayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPayment", reflect.TypeOf((*MockStorage)(nil).InsertPayment), arg0)
}
