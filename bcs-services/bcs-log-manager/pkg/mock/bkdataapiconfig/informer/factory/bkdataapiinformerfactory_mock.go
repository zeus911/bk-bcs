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
// Source: github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/informers/externalversions (interfaces: SharedInformerFactory)

// Package factory is a generated GoMock package.
package factory

import (
	externalversions "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/informers/externalversions"
	bkbcs "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/informers/externalversions/bkbcs.tencent.com"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/informers/externalversions/internalinterfaces"
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	reflect "reflect"
)

// MockSharedInformerFactory is a mock of SharedInformerFactory interface
type MockSharedInformerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSharedInformerFactoryMockRecorder
}

// MockSharedInformerFactoryMockRecorder is the mock recorder for MockSharedInformerFactory
type MockSharedInformerFactoryMockRecorder struct {
	mock *MockSharedInformerFactory
}

// NewMockSharedInformerFactory creates a new mock instance
func NewMockSharedInformerFactory(ctrl *gomock.Controller) *MockSharedInformerFactory {
	mock := &MockSharedInformerFactory{ctrl: ctrl}
	mock.recorder = &MockSharedInformerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSharedInformerFactory) EXPECT() *MockSharedInformerFactoryMockRecorder {
	return m.recorder
}

// Bkbcs mocks base method
func (m *MockSharedInformerFactory) Bkbcs() bkbcs.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bkbcs")
	ret0, _ := ret[0].(bkbcs.Interface)
	return ret0
}

// Bkbcs indicates an expected call of Bkbcs
func (mr *MockSharedInformerFactoryMockRecorder) Bkbcs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bkbcs", reflect.TypeOf((*MockSharedInformerFactory)(nil).Bkbcs))
}

// ForResource mocks base method
func (m *MockSharedInformerFactory) ForResource(arg0 schema.GroupVersionResource) (externalversions.GenericInformer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForResource", arg0)
	ret0, _ := ret[0].(externalversions.GenericInformer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForResource indicates an expected call of ForResource
func (mr *MockSharedInformerFactoryMockRecorder) ForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForResource", reflect.TypeOf((*MockSharedInformerFactory)(nil).ForResource), arg0)
}

// InformerFor mocks base method
func (m *MockSharedInformerFactory) InformerFor(arg0 runtime.Object, arg1 internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InformerFor", arg0, arg1)
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// InformerFor indicates an expected call of InformerFor
func (mr *MockSharedInformerFactoryMockRecorder) InformerFor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InformerFor", reflect.TypeOf((*MockSharedInformerFactory)(nil).InformerFor), arg0, arg1)
}

// Start mocks base method
func (m *MockSharedInformerFactory) Start(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockSharedInformerFactoryMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSharedInformerFactory)(nil).Start), arg0)
}

// WaitForCacheSync mocks base method
func (m *MockSharedInformerFactory) WaitForCacheSync(arg0 <-chan struct{}) map[reflect.Type]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCacheSync", arg0)
	ret0, _ := ret[0].(map[reflect.Type]bool)
	return ret0
}

// WaitForCacheSync indicates an expected call of WaitForCacheSync
func (mr *MockSharedInformerFactoryMockRecorder) WaitForCacheSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCacheSync", reflect.TypeOf((*MockSharedInformerFactory)(nil).WaitForCacheSync), arg0)
}
