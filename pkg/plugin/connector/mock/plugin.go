// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/plugin/connector (interfaces: Dispenser,DestinationPlugin,SourcePlugin,SpecifierPlugin)
//
// Generated by this command:
//
//	mockgen -typed -destination=mock/plugin.go -package=mock -mock_names=Dispenser=Dispenser,SourcePlugin=SourcePlugin,DestinationPlugin=DestinationPlugin,SpecifierPlugin=SpecifierPlugin . Dispenser,DestinationPlugin,SourcePlugin,SpecifierPlugin
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pconnector "github.com/conduitio/conduit-connector-protocol/pconnector"
	connector "github.com/conduitio/conduit/pkg/plugin/connector"
	gomock "go.uber.org/mock/gomock"
)

// Dispenser is a mock of Dispenser interface.
type Dispenser struct {
	ctrl     *gomock.Controller
	recorder *DispenserMockRecorder
}

// DispenserMockRecorder is the mock recorder for Dispenser.
type DispenserMockRecorder struct {
	mock *Dispenser
}

// NewDispenser creates a new mock instance.
func NewDispenser(ctrl *gomock.Controller) *Dispenser {
	mock := &Dispenser{ctrl: ctrl}
	mock.recorder = &DispenserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Dispenser) EXPECT() *DispenserMockRecorder {
	return m.recorder
}

// DispenseDestination mocks base method.
func (m *Dispenser) DispenseDestination() (connector.DestinationPlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispenseDestination")
	ret0, _ := ret[0].(connector.DestinationPlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispenseDestination indicates an expected call of DispenseDestination.
func (mr *DispenserMockRecorder) DispenseDestination() *DispenserDispenseDestinationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispenseDestination", reflect.TypeOf((*Dispenser)(nil).DispenseDestination))
	return &DispenserDispenseDestinationCall{Call: call}
}

// DispenserDispenseDestinationCall wrap *gomock.Call
type DispenserDispenseDestinationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DispenserDispenseDestinationCall) Return(arg0 connector.DestinationPlugin, arg1 error) *DispenserDispenseDestinationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DispenserDispenseDestinationCall) Do(f func() (connector.DestinationPlugin, error)) *DispenserDispenseDestinationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DispenserDispenseDestinationCall) DoAndReturn(f func() (connector.DestinationPlugin, error)) *DispenserDispenseDestinationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DispenseSource mocks base method.
func (m *Dispenser) DispenseSource() (connector.SourcePlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispenseSource")
	ret0, _ := ret[0].(connector.SourcePlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispenseSource indicates an expected call of DispenseSource.
func (mr *DispenserMockRecorder) DispenseSource() *DispenserDispenseSourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispenseSource", reflect.TypeOf((*Dispenser)(nil).DispenseSource))
	return &DispenserDispenseSourceCall{Call: call}
}

// DispenserDispenseSourceCall wrap *gomock.Call
type DispenserDispenseSourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DispenserDispenseSourceCall) Return(arg0 connector.SourcePlugin, arg1 error) *DispenserDispenseSourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DispenserDispenseSourceCall) Do(f func() (connector.SourcePlugin, error)) *DispenserDispenseSourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DispenserDispenseSourceCall) DoAndReturn(f func() (connector.SourcePlugin, error)) *DispenserDispenseSourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DispenseSpecifier mocks base method.
func (m *Dispenser) DispenseSpecifier() (connector.SpecifierPlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispenseSpecifier")
	ret0, _ := ret[0].(connector.SpecifierPlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispenseSpecifier indicates an expected call of DispenseSpecifier.
func (mr *DispenserMockRecorder) DispenseSpecifier() *DispenserDispenseSpecifierCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispenseSpecifier", reflect.TypeOf((*Dispenser)(nil).DispenseSpecifier))
	return &DispenserDispenseSpecifierCall{Call: call}
}

// DispenserDispenseSpecifierCall wrap *gomock.Call
type DispenserDispenseSpecifierCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DispenserDispenseSpecifierCall) Return(arg0 connector.SpecifierPlugin, arg1 error) *DispenserDispenseSpecifierCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DispenserDispenseSpecifierCall) Do(f func() (connector.SpecifierPlugin, error)) *DispenserDispenseSpecifierCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DispenserDispenseSpecifierCall) DoAndReturn(f func() (connector.SpecifierPlugin, error)) *DispenserDispenseSpecifierCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DestinationPlugin is a mock of DestinationPlugin interface.
type DestinationPlugin struct {
	ctrl     *gomock.Controller
	recorder *DestinationPluginMockRecorder
}

// DestinationPluginMockRecorder is the mock recorder for DestinationPlugin.
type DestinationPluginMockRecorder struct {
	mock *DestinationPlugin
}

// NewDestinationPlugin creates a new mock instance.
func NewDestinationPlugin(ctrl *gomock.Controller) *DestinationPlugin {
	mock := &DestinationPlugin{ctrl: ctrl}
	mock.recorder = &DestinationPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *DestinationPlugin) EXPECT() *DestinationPluginMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *DestinationPlugin) Configure(arg0 context.Context, arg1 pconnector.DestinationConfigureRequest) (pconnector.DestinationConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure.
func (mr *DestinationPluginMockRecorder) Configure(arg0, arg1 any) *DestinationPluginConfigureCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*DestinationPlugin)(nil).Configure), arg0, arg1)
	return &DestinationPluginConfigureCall{Call: call}
}

// DestinationPluginConfigureCall wrap *gomock.Call
type DestinationPluginConfigureCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginConfigureCall) Return(arg0 pconnector.DestinationConfigureResponse, arg1 error) *DestinationPluginConfigureCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginConfigureCall) Do(f func(context.Context, pconnector.DestinationConfigureRequest) (pconnector.DestinationConfigureResponse, error)) *DestinationPluginConfigureCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginConfigureCall) DoAndReturn(f func(context.Context, pconnector.DestinationConfigureRequest) (pconnector.DestinationConfigureResponse, error)) *DestinationPluginConfigureCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnCreated mocks base method.
func (m *DestinationPlugin) LifecycleOnCreated(arg0 context.Context, arg1 pconnector.DestinationLifecycleOnCreatedRequest) (pconnector.DestinationLifecycleOnCreatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnCreated", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationLifecycleOnCreatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnCreated indicates an expected call of LifecycleOnCreated.
func (mr *DestinationPluginMockRecorder) LifecycleOnCreated(arg0, arg1 any) *DestinationPluginLifecycleOnCreatedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnCreated", reflect.TypeOf((*DestinationPlugin)(nil).LifecycleOnCreated), arg0, arg1)
	return &DestinationPluginLifecycleOnCreatedCall{Call: call}
}

// DestinationPluginLifecycleOnCreatedCall wrap *gomock.Call
type DestinationPluginLifecycleOnCreatedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginLifecycleOnCreatedCall) Return(arg0 pconnector.DestinationLifecycleOnCreatedResponse, arg1 error) *DestinationPluginLifecycleOnCreatedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginLifecycleOnCreatedCall) Do(f func(context.Context, pconnector.DestinationLifecycleOnCreatedRequest) (pconnector.DestinationLifecycleOnCreatedResponse, error)) *DestinationPluginLifecycleOnCreatedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginLifecycleOnCreatedCall) DoAndReturn(f func(context.Context, pconnector.DestinationLifecycleOnCreatedRequest) (pconnector.DestinationLifecycleOnCreatedResponse, error)) *DestinationPluginLifecycleOnCreatedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnDeleted mocks base method.
func (m *DestinationPlugin) LifecycleOnDeleted(arg0 context.Context, arg1 pconnector.DestinationLifecycleOnDeletedRequest) (pconnector.DestinationLifecycleOnDeletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnDeleted", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationLifecycleOnDeletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnDeleted indicates an expected call of LifecycleOnDeleted.
func (mr *DestinationPluginMockRecorder) LifecycleOnDeleted(arg0, arg1 any) *DestinationPluginLifecycleOnDeletedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnDeleted", reflect.TypeOf((*DestinationPlugin)(nil).LifecycleOnDeleted), arg0, arg1)
	return &DestinationPluginLifecycleOnDeletedCall{Call: call}
}

// DestinationPluginLifecycleOnDeletedCall wrap *gomock.Call
type DestinationPluginLifecycleOnDeletedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginLifecycleOnDeletedCall) Return(arg0 pconnector.DestinationLifecycleOnDeletedResponse, arg1 error) *DestinationPluginLifecycleOnDeletedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginLifecycleOnDeletedCall) Do(f func(context.Context, pconnector.DestinationLifecycleOnDeletedRequest) (pconnector.DestinationLifecycleOnDeletedResponse, error)) *DestinationPluginLifecycleOnDeletedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginLifecycleOnDeletedCall) DoAndReturn(f func(context.Context, pconnector.DestinationLifecycleOnDeletedRequest) (pconnector.DestinationLifecycleOnDeletedResponse, error)) *DestinationPluginLifecycleOnDeletedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnUpdated mocks base method.
func (m *DestinationPlugin) LifecycleOnUpdated(arg0 context.Context, arg1 pconnector.DestinationLifecycleOnUpdatedRequest) (pconnector.DestinationLifecycleOnUpdatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnUpdated", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationLifecycleOnUpdatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnUpdated indicates an expected call of LifecycleOnUpdated.
func (mr *DestinationPluginMockRecorder) LifecycleOnUpdated(arg0, arg1 any) *DestinationPluginLifecycleOnUpdatedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnUpdated", reflect.TypeOf((*DestinationPlugin)(nil).LifecycleOnUpdated), arg0, arg1)
	return &DestinationPluginLifecycleOnUpdatedCall{Call: call}
}

// DestinationPluginLifecycleOnUpdatedCall wrap *gomock.Call
type DestinationPluginLifecycleOnUpdatedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginLifecycleOnUpdatedCall) Return(arg0 pconnector.DestinationLifecycleOnUpdatedResponse, arg1 error) *DestinationPluginLifecycleOnUpdatedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginLifecycleOnUpdatedCall) Do(f func(context.Context, pconnector.DestinationLifecycleOnUpdatedRequest) (pconnector.DestinationLifecycleOnUpdatedResponse, error)) *DestinationPluginLifecycleOnUpdatedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginLifecycleOnUpdatedCall) DoAndReturn(f func(context.Context, pconnector.DestinationLifecycleOnUpdatedRequest) (pconnector.DestinationLifecycleOnUpdatedResponse, error)) *DestinationPluginLifecycleOnUpdatedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewStream mocks base method.
func (m *DestinationPlugin) NewStream() pconnector.DestinationRunStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream")
	ret0, _ := ret[0].(pconnector.DestinationRunStream)
	return ret0
}

// NewStream indicates an expected call of NewStream.
func (mr *DestinationPluginMockRecorder) NewStream() *DestinationPluginNewStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*DestinationPlugin)(nil).NewStream))
	return &DestinationPluginNewStreamCall{Call: call}
}

// DestinationPluginNewStreamCall wrap *gomock.Call
type DestinationPluginNewStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginNewStreamCall) Return(arg0 pconnector.DestinationRunStream) *DestinationPluginNewStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginNewStreamCall) Do(f func() pconnector.DestinationRunStream) *DestinationPluginNewStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginNewStreamCall) DoAndReturn(f func() pconnector.DestinationRunStream) *DestinationPluginNewStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Open mocks base method.
func (m *DestinationPlugin) Open(arg0 context.Context, arg1 pconnector.DestinationOpenRequest) (pconnector.DestinationOpenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationOpenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *DestinationPluginMockRecorder) Open(arg0, arg1 any) *DestinationPluginOpenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*DestinationPlugin)(nil).Open), arg0, arg1)
	return &DestinationPluginOpenCall{Call: call}
}

// DestinationPluginOpenCall wrap *gomock.Call
type DestinationPluginOpenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginOpenCall) Return(arg0 pconnector.DestinationOpenResponse, arg1 error) *DestinationPluginOpenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginOpenCall) Do(f func(context.Context, pconnector.DestinationOpenRequest) (pconnector.DestinationOpenResponse, error)) *DestinationPluginOpenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginOpenCall) DoAndReturn(f func(context.Context, pconnector.DestinationOpenRequest) (pconnector.DestinationOpenResponse, error)) *DestinationPluginOpenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *DestinationPlugin) Run(arg0 context.Context, arg1 pconnector.DestinationRunStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *DestinationPluginMockRecorder) Run(arg0, arg1 any) *DestinationPluginRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*DestinationPlugin)(nil).Run), arg0, arg1)
	return &DestinationPluginRunCall{Call: call}
}

// DestinationPluginRunCall wrap *gomock.Call
type DestinationPluginRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginRunCall) Return(arg0 error) *DestinationPluginRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginRunCall) Do(f func(context.Context, pconnector.DestinationRunStream) error) *DestinationPluginRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginRunCall) DoAndReturn(f func(context.Context, pconnector.DestinationRunStream) error) *DestinationPluginRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *DestinationPlugin) Stop(arg0 context.Context, arg1 pconnector.DestinationStopRequest) (pconnector.DestinationStopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *DestinationPluginMockRecorder) Stop(arg0, arg1 any) *DestinationPluginStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*DestinationPlugin)(nil).Stop), arg0, arg1)
	return &DestinationPluginStopCall{Call: call}
}

// DestinationPluginStopCall wrap *gomock.Call
type DestinationPluginStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginStopCall) Return(arg0 pconnector.DestinationStopResponse, arg1 error) *DestinationPluginStopCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginStopCall) Do(f func(context.Context, pconnector.DestinationStopRequest) (pconnector.DestinationStopResponse, error)) *DestinationPluginStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginStopCall) DoAndReturn(f func(context.Context, pconnector.DestinationStopRequest) (pconnector.DestinationStopResponse, error)) *DestinationPluginStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Teardown mocks base method.
func (m *DestinationPlugin) Teardown(arg0 context.Context, arg1 pconnector.DestinationTeardownRequest) (pconnector.DestinationTeardownResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0, arg1)
	ret0, _ := ret[0].(pconnector.DestinationTeardownResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teardown indicates an expected call of Teardown.
func (mr *DestinationPluginMockRecorder) Teardown(arg0, arg1 any) *DestinationPluginTeardownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*DestinationPlugin)(nil).Teardown), arg0, arg1)
	return &DestinationPluginTeardownCall{Call: call}
}

// DestinationPluginTeardownCall wrap *gomock.Call
type DestinationPluginTeardownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *DestinationPluginTeardownCall) Return(arg0 pconnector.DestinationTeardownResponse, arg1 error) *DestinationPluginTeardownCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *DestinationPluginTeardownCall) Do(f func(context.Context, pconnector.DestinationTeardownRequest) (pconnector.DestinationTeardownResponse, error)) *DestinationPluginTeardownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *DestinationPluginTeardownCall) DoAndReturn(f func(context.Context, pconnector.DestinationTeardownRequest) (pconnector.DestinationTeardownResponse, error)) *DestinationPluginTeardownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SourcePlugin is a mock of SourcePlugin interface.
type SourcePlugin struct {
	ctrl     *gomock.Controller
	recorder *SourcePluginMockRecorder
}

// SourcePluginMockRecorder is the mock recorder for SourcePlugin.
type SourcePluginMockRecorder struct {
	mock *SourcePlugin
}

// NewSourcePlugin creates a new mock instance.
func NewSourcePlugin(ctrl *gomock.Controller) *SourcePlugin {
	mock := &SourcePlugin{ctrl: ctrl}
	mock.recorder = &SourcePluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *SourcePlugin) EXPECT() *SourcePluginMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *SourcePlugin) Configure(arg0 context.Context, arg1 pconnector.SourceConfigureRequest) (pconnector.SourceConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure.
func (mr *SourcePluginMockRecorder) Configure(arg0, arg1 any) *SourcePluginConfigureCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*SourcePlugin)(nil).Configure), arg0, arg1)
	return &SourcePluginConfigureCall{Call: call}
}

// SourcePluginConfigureCall wrap *gomock.Call
type SourcePluginConfigureCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginConfigureCall) Return(arg0 pconnector.SourceConfigureResponse, arg1 error) *SourcePluginConfigureCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginConfigureCall) Do(f func(context.Context, pconnector.SourceConfigureRequest) (pconnector.SourceConfigureResponse, error)) *SourcePluginConfigureCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginConfigureCall) DoAndReturn(f func(context.Context, pconnector.SourceConfigureRequest) (pconnector.SourceConfigureResponse, error)) *SourcePluginConfigureCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnCreated mocks base method.
func (m *SourcePlugin) LifecycleOnCreated(arg0 context.Context, arg1 pconnector.SourceLifecycleOnCreatedRequest) (pconnector.SourceLifecycleOnCreatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnCreated", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceLifecycleOnCreatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnCreated indicates an expected call of LifecycleOnCreated.
func (mr *SourcePluginMockRecorder) LifecycleOnCreated(arg0, arg1 any) *SourcePluginLifecycleOnCreatedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnCreated", reflect.TypeOf((*SourcePlugin)(nil).LifecycleOnCreated), arg0, arg1)
	return &SourcePluginLifecycleOnCreatedCall{Call: call}
}

// SourcePluginLifecycleOnCreatedCall wrap *gomock.Call
type SourcePluginLifecycleOnCreatedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginLifecycleOnCreatedCall) Return(arg0 pconnector.SourceLifecycleOnCreatedResponse, arg1 error) *SourcePluginLifecycleOnCreatedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginLifecycleOnCreatedCall) Do(f func(context.Context, pconnector.SourceLifecycleOnCreatedRequest) (pconnector.SourceLifecycleOnCreatedResponse, error)) *SourcePluginLifecycleOnCreatedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginLifecycleOnCreatedCall) DoAndReturn(f func(context.Context, pconnector.SourceLifecycleOnCreatedRequest) (pconnector.SourceLifecycleOnCreatedResponse, error)) *SourcePluginLifecycleOnCreatedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnDeleted mocks base method.
func (m *SourcePlugin) LifecycleOnDeleted(arg0 context.Context, arg1 pconnector.SourceLifecycleOnDeletedRequest) (pconnector.SourceLifecycleOnDeletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnDeleted", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceLifecycleOnDeletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnDeleted indicates an expected call of LifecycleOnDeleted.
func (mr *SourcePluginMockRecorder) LifecycleOnDeleted(arg0, arg1 any) *SourcePluginLifecycleOnDeletedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnDeleted", reflect.TypeOf((*SourcePlugin)(nil).LifecycleOnDeleted), arg0, arg1)
	return &SourcePluginLifecycleOnDeletedCall{Call: call}
}

// SourcePluginLifecycleOnDeletedCall wrap *gomock.Call
type SourcePluginLifecycleOnDeletedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginLifecycleOnDeletedCall) Return(arg0 pconnector.SourceLifecycleOnDeletedResponse, arg1 error) *SourcePluginLifecycleOnDeletedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginLifecycleOnDeletedCall) Do(f func(context.Context, pconnector.SourceLifecycleOnDeletedRequest) (pconnector.SourceLifecycleOnDeletedResponse, error)) *SourcePluginLifecycleOnDeletedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginLifecycleOnDeletedCall) DoAndReturn(f func(context.Context, pconnector.SourceLifecycleOnDeletedRequest) (pconnector.SourceLifecycleOnDeletedResponse, error)) *SourcePluginLifecycleOnDeletedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LifecycleOnUpdated mocks base method.
func (m *SourcePlugin) LifecycleOnUpdated(arg0 context.Context, arg1 pconnector.SourceLifecycleOnUpdatedRequest) (pconnector.SourceLifecycleOnUpdatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LifecycleOnUpdated", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceLifecycleOnUpdatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LifecycleOnUpdated indicates an expected call of LifecycleOnUpdated.
func (mr *SourcePluginMockRecorder) LifecycleOnUpdated(arg0, arg1 any) *SourcePluginLifecycleOnUpdatedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LifecycleOnUpdated", reflect.TypeOf((*SourcePlugin)(nil).LifecycleOnUpdated), arg0, arg1)
	return &SourcePluginLifecycleOnUpdatedCall{Call: call}
}

// SourcePluginLifecycleOnUpdatedCall wrap *gomock.Call
type SourcePluginLifecycleOnUpdatedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginLifecycleOnUpdatedCall) Return(arg0 pconnector.SourceLifecycleOnUpdatedResponse, arg1 error) *SourcePluginLifecycleOnUpdatedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginLifecycleOnUpdatedCall) Do(f func(context.Context, pconnector.SourceLifecycleOnUpdatedRequest) (pconnector.SourceLifecycleOnUpdatedResponse, error)) *SourcePluginLifecycleOnUpdatedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginLifecycleOnUpdatedCall) DoAndReturn(f func(context.Context, pconnector.SourceLifecycleOnUpdatedRequest) (pconnector.SourceLifecycleOnUpdatedResponse, error)) *SourcePluginLifecycleOnUpdatedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewStream mocks base method.
func (m *SourcePlugin) NewStream() pconnector.SourceRunStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream")
	ret0, _ := ret[0].(pconnector.SourceRunStream)
	return ret0
}

// NewStream indicates an expected call of NewStream.
func (mr *SourcePluginMockRecorder) NewStream() *SourcePluginNewStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*SourcePlugin)(nil).NewStream))
	return &SourcePluginNewStreamCall{Call: call}
}

// SourcePluginNewStreamCall wrap *gomock.Call
type SourcePluginNewStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginNewStreamCall) Return(arg0 pconnector.SourceRunStream) *SourcePluginNewStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginNewStreamCall) Do(f func() pconnector.SourceRunStream) *SourcePluginNewStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginNewStreamCall) DoAndReturn(f func() pconnector.SourceRunStream) *SourcePluginNewStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Open mocks base method.
func (m *SourcePlugin) Open(arg0 context.Context, arg1 pconnector.SourceOpenRequest) (pconnector.SourceOpenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceOpenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *SourcePluginMockRecorder) Open(arg0, arg1 any) *SourcePluginOpenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*SourcePlugin)(nil).Open), arg0, arg1)
	return &SourcePluginOpenCall{Call: call}
}

// SourcePluginOpenCall wrap *gomock.Call
type SourcePluginOpenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginOpenCall) Return(arg0 pconnector.SourceOpenResponse, arg1 error) *SourcePluginOpenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginOpenCall) Do(f func(context.Context, pconnector.SourceOpenRequest) (pconnector.SourceOpenResponse, error)) *SourcePluginOpenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginOpenCall) DoAndReturn(f func(context.Context, pconnector.SourceOpenRequest) (pconnector.SourceOpenResponse, error)) *SourcePluginOpenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *SourcePlugin) Run(arg0 context.Context, arg1 pconnector.SourceRunStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *SourcePluginMockRecorder) Run(arg0, arg1 any) *SourcePluginRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*SourcePlugin)(nil).Run), arg0, arg1)
	return &SourcePluginRunCall{Call: call}
}

// SourcePluginRunCall wrap *gomock.Call
type SourcePluginRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginRunCall) Return(arg0 error) *SourcePluginRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginRunCall) Do(f func(context.Context, pconnector.SourceRunStream) error) *SourcePluginRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginRunCall) DoAndReturn(f func(context.Context, pconnector.SourceRunStream) error) *SourcePluginRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *SourcePlugin) Stop(arg0 context.Context, arg1 pconnector.SourceStopRequest) (pconnector.SourceStopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *SourcePluginMockRecorder) Stop(arg0, arg1 any) *SourcePluginStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*SourcePlugin)(nil).Stop), arg0, arg1)
	return &SourcePluginStopCall{Call: call}
}

// SourcePluginStopCall wrap *gomock.Call
type SourcePluginStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginStopCall) Return(arg0 pconnector.SourceStopResponse, arg1 error) *SourcePluginStopCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginStopCall) Do(f func(context.Context, pconnector.SourceStopRequest) (pconnector.SourceStopResponse, error)) *SourcePluginStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginStopCall) DoAndReturn(f func(context.Context, pconnector.SourceStopRequest) (pconnector.SourceStopResponse, error)) *SourcePluginStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Teardown mocks base method.
func (m *SourcePlugin) Teardown(arg0 context.Context, arg1 pconnector.SourceTeardownRequest) (pconnector.SourceTeardownResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SourceTeardownResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teardown indicates an expected call of Teardown.
func (mr *SourcePluginMockRecorder) Teardown(arg0, arg1 any) *SourcePluginTeardownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*SourcePlugin)(nil).Teardown), arg0, arg1)
	return &SourcePluginTeardownCall{Call: call}
}

// SourcePluginTeardownCall wrap *gomock.Call
type SourcePluginTeardownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SourcePluginTeardownCall) Return(arg0 pconnector.SourceTeardownResponse, arg1 error) *SourcePluginTeardownCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SourcePluginTeardownCall) Do(f func(context.Context, pconnector.SourceTeardownRequest) (pconnector.SourceTeardownResponse, error)) *SourcePluginTeardownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SourcePluginTeardownCall) DoAndReturn(f func(context.Context, pconnector.SourceTeardownRequest) (pconnector.SourceTeardownResponse, error)) *SourcePluginTeardownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SpecifierPlugin is a mock of SpecifierPlugin interface.
type SpecifierPlugin struct {
	ctrl     *gomock.Controller
	recorder *SpecifierPluginMockRecorder
}

// SpecifierPluginMockRecorder is the mock recorder for SpecifierPlugin.
type SpecifierPluginMockRecorder struct {
	mock *SpecifierPlugin
}

// NewSpecifierPlugin creates a new mock instance.
func NewSpecifierPlugin(ctrl *gomock.Controller) *SpecifierPlugin {
	mock := &SpecifierPlugin{ctrl: ctrl}
	mock.recorder = &SpecifierPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *SpecifierPlugin) EXPECT() *SpecifierPluginMockRecorder {
	return m.recorder
}

// Specify mocks base method.
func (m *SpecifierPlugin) Specify(arg0 context.Context, arg1 pconnector.SpecifierSpecifyRequest) (pconnector.SpecifierSpecifyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Specify", arg0, arg1)
	ret0, _ := ret[0].(pconnector.SpecifierSpecifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Specify indicates an expected call of Specify.
func (mr *SpecifierPluginMockRecorder) Specify(arg0, arg1 any) *SpecifierPluginSpecifyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Specify", reflect.TypeOf((*SpecifierPlugin)(nil).Specify), arg0, arg1)
	return &SpecifierPluginSpecifyCall{Call: call}
}

// SpecifierPluginSpecifyCall wrap *gomock.Call
type SpecifierPluginSpecifyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SpecifierPluginSpecifyCall) Return(arg0 pconnector.SpecifierSpecifyResponse, arg1 error) *SpecifierPluginSpecifyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SpecifierPluginSpecifyCall) Do(f func(context.Context, pconnector.SpecifierSpecifyRequest) (pconnector.SpecifierSpecifyResponse, error)) *SpecifierPluginSpecifyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SpecifierPluginSpecifyCall) DoAndReturn(f func(context.Context, pconnector.SpecifierSpecifyRequest) (pconnector.SpecifierSpecifyResponse, error)) *SpecifierPluginSpecifyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
