// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/querier (interfaces: AgentNetworkPolicyInfoQuerier)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1beta1"
	reflect "reflect"
)

// MockAgentNetworkPolicyInfoQuerier is a mock of AgentNetworkPolicyInfoQuerier interface
type MockAgentNetworkPolicyInfoQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockAgentNetworkPolicyInfoQuerierMockRecorder
}

// MockAgentNetworkPolicyInfoQuerierMockRecorder is the mock recorder for MockAgentNetworkPolicyInfoQuerier
type MockAgentNetworkPolicyInfoQuerierMockRecorder struct {
	mock *MockAgentNetworkPolicyInfoQuerier
}

// NewMockAgentNetworkPolicyInfoQuerier creates a new mock instance
func NewMockAgentNetworkPolicyInfoQuerier(ctrl *gomock.Controller) *MockAgentNetworkPolicyInfoQuerier {
	mock := &MockAgentNetworkPolicyInfoQuerier{ctrl: ctrl}
	mock.recorder = &MockAgentNetworkPolicyInfoQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentNetworkPolicyInfoQuerier) EXPECT() *MockAgentNetworkPolicyInfoQuerierMockRecorder {
	return m.recorder
}

// GetAddressGroupNum mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetAddressGroupNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressGroupNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAddressGroupNum indicates an expected call of GetAddressGroupNum
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetAddressGroupNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressGroupNum", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetAddressGroupNum))
}

// GetAddressGroups mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetAddressGroups() []v1beta1.AddressGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressGroups")
	ret0, _ := ret[0].([]v1beta1.AddressGroup)
	return ret0
}

// GetAddressGroups indicates an expected call of GetAddressGroups
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetAddressGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressGroups", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetAddressGroups))
}

// GetAppliedNetworkPolicies mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetAppliedNetworkPolicies(arg0, arg1 string) []v1beta1.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppliedNetworkPolicies", arg0, arg1)
	ret0, _ := ret[0].([]v1beta1.NetworkPolicy)
	return ret0
}

// GetAppliedNetworkPolicies indicates an expected call of GetAppliedNetworkPolicies
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetAppliedNetworkPolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppliedNetworkPolicies", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetAppliedNetworkPolicies), arg0, arg1)
}

// GetAppliedToGroupNum mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetAppliedToGroupNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppliedToGroupNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAppliedToGroupNum indicates an expected call of GetAppliedToGroupNum
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetAppliedToGroupNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppliedToGroupNum", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetAppliedToGroupNum))
}

// GetAppliedToGroups mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetAppliedToGroups() []v1beta1.AppliedToGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppliedToGroups")
	ret0, _ := ret[0].([]v1beta1.AppliedToGroup)
	return ret0
}

// GetAppliedToGroups indicates an expected call of GetAppliedToGroups
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetAppliedToGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppliedToGroups", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetAppliedToGroups))
}

// GetControllerConnectionStatus mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetControllerConnectionStatus() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControllerConnectionStatus")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetControllerConnectionStatus indicates an expected call of GetControllerConnectionStatus
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetControllerConnectionStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControllerConnectionStatus", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetControllerConnectionStatus))
}

// GetNetworkPolicies mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetNetworkPolicies(arg0 string) []v1beta1.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkPolicies", arg0)
	ret0, _ := ret[0].([]v1beta1.NetworkPolicy)
	return ret0
}

// GetNetworkPolicies indicates an expected call of GetNetworkPolicies
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetNetworkPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPolicies", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetNetworkPolicies), arg0)
}

// GetNetworkPolicy mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetNetworkPolicy(arg0, arg1 string) *v1beta1.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkPolicy", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.NetworkPolicy)
	return ret0
}

// GetNetworkPolicy indicates an expected call of GetNetworkPolicy
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetNetworkPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPolicy", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetNetworkPolicy), arg0, arg1)
}

// GetNetworkPolicyNum mocks base method
func (m *MockAgentNetworkPolicyInfoQuerier) GetNetworkPolicyNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkPolicyNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNetworkPolicyNum indicates an expected call of GetNetworkPolicyNum
func (mr *MockAgentNetworkPolicyInfoQuerierMockRecorder) GetNetworkPolicyNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPolicyNum", reflect.TypeOf((*MockAgentNetworkPolicyInfoQuerier)(nil).GetNetworkPolicyNum))
}
