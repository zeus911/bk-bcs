// Tencent is pleased to support the open source community by making Blueking Container Service available.
// Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except
// in compliance with the License. You may obtain a copy of the License at
// http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under
// the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/clientset/versioned (interfaces: Interface)

// Package clientset is a generated GoMock package.
package clientset

import (
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/clientset/versioned/typed/bkbcs.tencent.com/v1"
	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Bkbcs mocks base method
func (m *MockInterface) Bkbcs() v1.BkbcsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bkbcs")
	ret0, _ := ret[0].(v1.BkbcsV1Interface)
	return ret0
}

// Bkbcs indicates an expected call of Bkbcs
func (mr *MockInterfaceMockRecorder) Bkbcs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bkbcs", reflect.TypeOf((*MockInterface)(nil).Bkbcs))
}

// BkbcsV1 mocks base method
func (m *MockInterface) BkbcsV1() v1.BkbcsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BkbcsV1")
	ret0, _ := ret[0].(v1.BkbcsV1Interface)
	return ret0
}

// BkbcsV1 indicates an expected call of BkbcsV1
func (mr *MockInterfaceMockRecorder) BkbcsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BkbcsV1", reflect.TypeOf((*MockInterface)(nil).BkbcsV1))
}

// Discovery mocks base method
func (m *MockInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery
func (mr *MockInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockInterface)(nil).Discovery))
}
