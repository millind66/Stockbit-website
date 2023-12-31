// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../entity/interfaces.go

package repository

import (
	sql "database/sql"

	gomock "github.com/golang/mock/gomock"
)

// Mock of DBClient interface
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *_MockDBClientRecorder
}

// Recorder for MockDBClient (not exported)
type _MockDBClientRecorder struct {
	mock *MockDBClient
}

func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &_MockDBClientRecorder{mock}
	return mock
}

func (_m *MockDBClient) EXPECT() *_MockDBClientRecorder {
	return _m.recorder
}

func (_m *MockDBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	_s := []interface{}{query}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBClientRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", _s...)
}
