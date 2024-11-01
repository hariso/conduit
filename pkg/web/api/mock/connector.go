// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/web/api (interfaces: ConnectorOrchestrator)
//
// Generated by this command:
//
//	mockgen -typed -destination=mock/connector.go -package=mock -mock_names=ConnectorOrchestrator=ConnectorOrchestrator . ConnectorOrchestrator
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	connector "github.com/conduitio/conduit/pkg/connector"
	inspector "github.com/conduitio/conduit/pkg/inspector"
	gomock "go.uber.org/mock/gomock"
)

// ConnectorOrchestrator is a mock of ConnectorOrchestrator interface.
type ConnectorOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *ConnectorOrchestratorMockRecorder
	isgomock struct{}
}

// ConnectorOrchestratorMockRecorder is the mock recorder for ConnectorOrchestrator.
type ConnectorOrchestratorMockRecorder struct {
	mock *ConnectorOrchestrator
}

// NewConnectorOrchestrator creates a new mock instance.
func NewConnectorOrchestrator(ctrl *gomock.Controller) *ConnectorOrchestrator {
	mock := &ConnectorOrchestrator{ctrl: ctrl}
	mock.recorder = &ConnectorOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ConnectorOrchestrator) EXPECT() *ConnectorOrchestratorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *ConnectorOrchestrator) Create(ctx context.Context, t connector.Type, plugin, pipelineID string, config connector.Config) (*connector.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, t, plugin, pipelineID, config)
	ret0, _ := ret[0].(*connector.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *ConnectorOrchestratorMockRecorder) Create(ctx, t, plugin, pipelineID, config any) *ConnectorOrchestratorCreateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*ConnectorOrchestrator)(nil).Create), ctx, t, plugin, pipelineID, config)
	return &ConnectorOrchestratorCreateCall{Call: call}
}

// ConnectorOrchestratorCreateCall wrap *gomock.Call
type ConnectorOrchestratorCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorCreateCall) Return(arg0 *connector.Instance, arg1 error) *ConnectorOrchestratorCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorCreateCall) Do(f func(context.Context, connector.Type, string, string, connector.Config) (*connector.Instance, error)) *ConnectorOrchestratorCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorCreateCall) DoAndReturn(f func(context.Context, connector.Type, string, string, connector.Config) (*connector.Instance, error)) *ConnectorOrchestratorCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *ConnectorOrchestrator) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *ConnectorOrchestratorMockRecorder) Delete(ctx, id any) *ConnectorOrchestratorDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*ConnectorOrchestrator)(nil).Delete), ctx, id)
	return &ConnectorOrchestratorDeleteCall{Call: call}
}

// ConnectorOrchestratorDeleteCall wrap *gomock.Call
type ConnectorOrchestratorDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorDeleteCall) Return(arg0 error) *ConnectorOrchestratorDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorDeleteCall) Do(f func(context.Context, string) error) *ConnectorOrchestratorDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorDeleteCall) DoAndReturn(f func(context.Context, string) error) *ConnectorOrchestratorDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *ConnectorOrchestrator) Get(ctx context.Context, id string) (*connector.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*connector.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *ConnectorOrchestratorMockRecorder) Get(ctx, id any) *ConnectorOrchestratorGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*ConnectorOrchestrator)(nil).Get), ctx, id)
	return &ConnectorOrchestratorGetCall{Call: call}
}

// ConnectorOrchestratorGetCall wrap *gomock.Call
type ConnectorOrchestratorGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorGetCall) Return(arg0 *connector.Instance, arg1 error) *ConnectorOrchestratorGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorGetCall) Do(f func(context.Context, string) (*connector.Instance, error)) *ConnectorOrchestratorGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorGetCall) DoAndReturn(f func(context.Context, string) (*connector.Instance, error)) *ConnectorOrchestratorGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Inspect mocks base method.
func (m *ConnectorOrchestrator) Inspect(ctx context.Context, id string) (*inspector.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", ctx, id)
	ret0, _ := ret[0].(*inspector.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect.
func (mr *ConnectorOrchestratorMockRecorder) Inspect(ctx, id any) *ConnectorOrchestratorInspectCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*ConnectorOrchestrator)(nil).Inspect), ctx, id)
	return &ConnectorOrchestratorInspectCall{Call: call}
}

// ConnectorOrchestratorInspectCall wrap *gomock.Call
type ConnectorOrchestratorInspectCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorInspectCall) Return(arg0 *inspector.Session, arg1 error) *ConnectorOrchestratorInspectCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorInspectCall) Do(f func(context.Context, string) (*inspector.Session, error)) *ConnectorOrchestratorInspectCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorInspectCall) DoAndReturn(f func(context.Context, string) (*inspector.Session, error)) *ConnectorOrchestratorInspectCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *ConnectorOrchestrator) List(ctx context.Context) map[string]*connector.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(map[string]*connector.Instance)
	return ret0
}

// List indicates an expected call of List.
func (mr *ConnectorOrchestratorMockRecorder) List(ctx any) *ConnectorOrchestratorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*ConnectorOrchestrator)(nil).List), ctx)
	return &ConnectorOrchestratorListCall{Call: call}
}

// ConnectorOrchestratorListCall wrap *gomock.Call
type ConnectorOrchestratorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorListCall) Return(arg0 map[string]*connector.Instance) *ConnectorOrchestratorListCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorListCall) Do(f func(context.Context) map[string]*connector.Instance) *ConnectorOrchestratorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorListCall) DoAndReturn(f func(context.Context) map[string]*connector.Instance) *ConnectorOrchestratorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *ConnectorOrchestrator) Update(ctx context.Context, id, plugin string, config connector.Config) (*connector.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, plugin, config)
	ret0, _ := ret[0].(*connector.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *ConnectorOrchestratorMockRecorder) Update(ctx, id, plugin, config any) *ConnectorOrchestratorUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*ConnectorOrchestrator)(nil).Update), ctx, id, plugin, config)
	return &ConnectorOrchestratorUpdateCall{Call: call}
}

// ConnectorOrchestratorUpdateCall wrap *gomock.Call
type ConnectorOrchestratorUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorUpdateCall) Return(arg0 *connector.Instance, arg1 error) *ConnectorOrchestratorUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorUpdateCall) Do(f func(context.Context, string, string, connector.Config) (*connector.Instance, error)) *ConnectorOrchestratorUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorUpdateCall) DoAndReturn(f func(context.Context, string, string, connector.Config) (*connector.Instance, error)) *ConnectorOrchestratorUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Validate mocks base method.
func (m *ConnectorOrchestrator) Validate(ctx context.Context, t connector.Type, plugin string, config connector.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, t, plugin, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *ConnectorOrchestratorMockRecorder) Validate(ctx, t, plugin, config any) *ConnectorOrchestratorValidateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*ConnectorOrchestrator)(nil).Validate), ctx, t, plugin, config)
	return &ConnectorOrchestratorValidateCall{Call: call}
}

// ConnectorOrchestratorValidateCall wrap *gomock.Call
type ConnectorOrchestratorValidateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnectorOrchestratorValidateCall) Return(arg0 error) *ConnectorOrchestratorValidateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnectorOrchestratorValidateCall) Do(f func(context.Context, connector.Type, string, connector.Config) error) *ConnectorOrchestratorValidateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnectorOrchestratorValidateCall) DoAndReturn(f func(context.Context, connector.Type, string, connector.Config) error) *ConnectorOrchestratorValidateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
