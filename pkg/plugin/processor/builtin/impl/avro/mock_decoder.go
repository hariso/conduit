// Code generated by MockGen. DO NOT EDIT.
// Source: decode.go
//
// Generated by this command:
//
//	mockgen -typed -source decode.go -destination=mock_decoder.go -package=avro -mock_names=decoder=MockDecoder . decoder
//

// Package avro is a generated GoMock package.
package avro

import (
	context "context"
	reflect "reflect"

	opencdc "github.com/conduitio/conduit-commons/opencdc"
	gomock "go.uber.org/mock/gomock"
)

// MockDecoder is a mock of decoder interface.
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
	isgomock struct{}
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder.
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance.
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoder) Decode(ctx context.Context, b opencdc.RawData) (opencdc.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", ctx, b)
	ret0, _ := ret[0].(opencdc.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderMockRecorder) Decode(ctx, b any) *MockDecoderDecodeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), ctx, b)
	return &MockDecoderDecodeCall{Call: call}
}

// MockDecoderDecodeCall wrap *gomock.Call
type MockDecoderDecodeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDecoderDecodeCall) Return(arg0 opencdc.StructuredData, arg1 error) *MockDecoderDecodeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDecoderDecodeCall) Do(f func(context.Context, opencdc.RawData) (opencdc.StructuredData, error)) *MockDecoderDecodeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDecoderDecodeCall) DoAndReturn(f func(context.Context, opencdc.RawData) (opencdc.StructuredData, error)) *MockDecoderDecodeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
