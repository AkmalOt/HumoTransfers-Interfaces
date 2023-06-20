// Code generated by MockGen. DO NOT EDIT.
// Source: Agent.go

// Package mock_Agent is a generated GoMock package.
package mock_Agent

import (
	models "awesomeProject2/models"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockPaginate is a mock of Paginate interface.
type MockPaginate struct {
	ctrl     *gomock.Controller
	recorder *MockPaginateMockRecorder
}

// MockPaginateMockRecorder is the mock recorder for MockPaginate.
type MockPaginateMockRecorder struct {
	mock *MockPaginate
}

// NewMockPaginate creates a new mock instance.
func NewMockPaginate(ctrl *gomock.Controller) *MockPaginate {
	mock := &MockPaginate{ctrl: ctrl}
	mock.recorder = &MockPaginateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaginate) EXPECT() *MockPaginateMockRecorder {
	return m.recorder
}

// GeneratePaginationFromRequest mocks base method.
func (m *MockPaginate) GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePaginationFromRequest", ctx)
	ret0, _ := ret[0].(models.Pagination)
	return ret0
}

// GeneratePaginationFromRequest indicates an expected call of GeneratePaginationFromRequest.
func (mr *MockPaginateMockRecorder) GeneratePaginationFromRequest(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePaginationFromRequest", reflect.TypeOf((*MockPaginate)(nil).GeneratePaginationFromRequest), ctx)
}

// MockAgentInterface is a mock of AgentInterface interface.
type MockAgentInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAgentInterfaceMockRecorder
}

// MockAgentInterfaceMockRecorder is the mock recorder for MockAgentInterface.
type MockAgentInterfaceMockRecorder struct {
	mock *MockAgentInterface
}

// NewMockAgentInterface creates a new mock instance.
func NewMockAgentInterface(ctrl *gomock.Controller) *MockAgentInterface {
	mock := &MockAgentInterface{ctrl: ctrl}
	mock.recorder = &MockAgentInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentInterface) EXPECT() *MockAgentInterfaceMockRecorder {
	return m.recorder
}

// AddAgent mocks base method.
func (m *MockAgentInterface) AddAgent(agents *models.Agents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgent", agents)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAgent indicates an expected call of AddAgent.
func (mr *MockAgentInterfaceMockRecorder) AddAgent(agents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgent", reflect.TypeOf((*MockAgentInterface)(nil).AddAgent), agents)
}

// AgentStatus mocks base method.
func (m *MockAgentInterface) AgentStatus(agent *models.Agents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentStatus", agent)
	ret0, _ := ret[0].(error)
	return ret0
}

// AgentStatus indicates an expected call of AgentStatus.
func (mr *MockAgentInterfaceMockRecorder) AgentStatus(agent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentStatus", reflect.TypeOf((*MockAgentInterface)(nil).AgentStatus), agent)
}

// DeleteAgents mocks base method.
func (m *MockAgentInterface) DeleteAgents(agent *models.Agents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgents", agent)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgents indicates an expected call of DeleteAgents.
func (mr *MockAgentInterfaceMockRecorder) DeleteAgents(agent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgents", reflect.TypeOf((*MockAgentInterface)(nil).DeleteAgents), agent)
}

// GetAgent mocks base method.
func (m *MockAgentInterface) GetAgent(pagination *models.Pagination) ([]models.Agents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", pagination)
	ret0, _ := ret[0].([]models.Agents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockAgentInterfaceMockRecorder) GetAgent(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockAgentInterface)(nil).GetAgent), pagination)
}

// TotalPageAgents mocks base method.
func (m *MockAgentInterface) TotalPageAgents(limit int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalPageAgents", limit)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalPageAgents indicates an expected call of TotalPageAgents.
func (mr *MockAgentInterfaceMockRecorder) TotalPageAgents(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalPageAgents", reflect.TypeOf((*MockAgentInterface)(nil).TotalPageAgents), limit)
}

// UpdateAgents mocks base method.
func (m *MockAgentInterface) UpdateAgents(agent *models.Agents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgents", agent)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgents indicates an expected call of UpdateAgents.
func (mr *MockAgentInterfaceMockRecorder) UpdateAgents(agent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgents", reflect.TypeOf((*MockAgentInterface)(nil).UpdateAgents), agent)
}