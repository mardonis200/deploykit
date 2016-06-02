// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/storage (interfaces: Keys)

package storage

import (
	ssh "github.com/docker/libmachete/ssh"
	storage "github.com/docker/libmachete/storage"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Keys interface
type MockKeys struct {
	ctrl     *gomock.Controller
	recorder *_MockKeysRecorder
}

// Recorder for MockKeys (not exported)
type _MockKeysRecorder struct {
	mock *MockKeys
}

func NewMockKeys(ctrl *gomock.Controller) *MockKeys {
	mock := &MockKeys{ctrl: ctrl}
	mock.recorder = &_MockKeysRecorder{mock}
	return mock
}

func (_m *MockKeys) EXPECT() *_MockKeysRecorder {
	return _m.recorder
}

func (_m *MockKeys) Delete(_param0 storage.KeyID) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKeysRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockKeys) GetEncodedPublicKey(_param0 storage.KeyID) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "GetEncodedPublicKey", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKeysRecorder) GetEncodedPublicKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEncodedPublicKey", arg0)
}

func (_m *MockKeys) List() ([]storage.KeyID, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]storage.KeyID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKeysRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockKeys) Save(_param0 storage.KeyID, _param1 *ssh.KeyPair) error {
	ret := _m.ctrl.Call(_m, "Save", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKeysRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0, arg1)
}
