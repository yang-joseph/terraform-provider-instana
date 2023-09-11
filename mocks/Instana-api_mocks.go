// Code generated by MockGen. DO NOT EDIT.
// Source: instana/restapi/Instana-api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	restapi "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	gomock "go.uber.org/mock/gomock"
)

// MockInstanaAPI is a mock of InstanaAPI interface.
type MockInstanaAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInstanaAPIMockRecorder
}

// MockInstanaAPIMockRecorder is the mock recorder for MockInstanaAPI.
type MockInstanaAPIMockRecorder struct {
	mock *MockInstanaAPI
}

// NewMockInstanaAPI creates a new mock instance.
func NewMockInstanaAPI(ctrl *gomock.Controller) *MockInstanaAPI {
	mock := &MockInstanaAPI{ctrl: ctrl}
	mock.recorder = &MockInstanaAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanaAPI) EXPECT() *MockInstanaAPIMockRecorder {
	return m.recorder
}

// APITokens mocks base method.
func (m *MockInstanaAPI) APITokens() restapi.RestResource[*restapi.APIToken] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APITokens")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.APIToken])
	return ret0
}

// APITokens indicates an expected call of APITokens.
func (mr *MockInstanaAPIMockRecorder) APITokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APITokens", reflect.TypeOf((*MockInstanaAPI)(nil).APITokens))
}

// AlertingChannels mocks base method.
func (m *MockInstanaAPI) AlertingChannels() restapi.RestResource[*restapi.AlertingChannel] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertingChannels")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.AlertingChannel])
	return ret0
}

// AlertingChannels indicates an expected call of AlertingChannels.
func (mr *MockInstanaAPIMockRecorder) AlertingChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingChannels", reflect.TypeOf((*MockInstanaAPI)(nil).AlertingChannels))
}

// AlertingChannels indicates an expected call of AlertingChannels.
func (mr *MockInstanaAPIMockRecorder) AlertingChannelsDS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingChannelsDS", reflect.TypeOf((*MockInstanaAPI)(nil).AlertingChannelsDS))
}

// AlertingConfigurations mocks base method.
func (m *MockInstanaAPI) AlertingConfigurations() restapi.RestResource[*restapi.AlertingConfiguration] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertingConfigurations")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.AlertingConfiguration])
	return ret0
}

// AlertingConfigurations indicates an expected call of AlertingConfigurations.
func (mr *MockInstanaAPIMockRecorder) AlertingConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingConfigurations", reflect.TypeOf((*MockInstanaAPI)(nil).AlertingConfigurations))
}

// ApplicationAlertConfigs mocks base method.
func (m *MockInstanaAPI) ApplicationAlertConfigs() restapi.RestResource[*restapi.ApplicationAlertConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationAlertConfigs")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.ApplicationAlertConfig])
	return ret0
}

// ApplicationAlertConfigs indicates an expected call of ApplicationAlertConfigs.
func (mr *MockInstanaAPIMockRecorder) ApplicationAlertConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationAlertConfigs", reflect.TypeOf((*MockInstanaAPI)(nil).ApplicationAlertConfigs))
}

// ApplicationConfigs mocks base method.
func (m *MockInstanaAPI) ApplicationConfigs() restapi.RestResource[*restapi.ApplicationConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationConfigs")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.ApplicationConfig])
	return ret0
}

// ApplicationConfigs indicates an expected call of ApplicationConfigs.
func (mr *MockInstanaAPIMockRecorder) ApplicationConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationConfigs", reflect.TypeOf((*MockInstanaAPI)(nil).ApplicationConfigs))
}

// BuiltinEventSpecifications mocks base method.
func (m *MockInstanaAPI) BuiltinEventSpecifications() restapi.ReadOnlyRestResource[*restapi.BuiltinEventSpecification] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuiltinEventSpecifications")
	ret0, _ := ret[0].(restapi.ReadOnlyRestResource[*restapi.BuiltinEventSpecification])
	return ret0
}

// BuiltinEventSpecifications indicates an expected call of BuiltinEventSpecifications.
func (mr *MockInstanaAPIMockRecorder) BuiltinEventSpecifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuiltinEventSpecifications", reflect.TypeOf((*MockInstanaAPI)(nil).BuiltinEventSpecifications))
}

// CustomDashboards mocks base method.
func (m *MockInstanaAPI) CustomDashboards() restapi.RestResource[*restapi.CustomDashboard] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomDashboards")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.CustomDashboard])
	return ret0
}

// CustomDashboards indicates an expected call of CustomDashboards.
func (mr *MockInstanaAPIMockRecorder) CustomDashboards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomDashboards", reflect.TypeOf((*MockInstanaAPI)(nil).CustomDashboards))
}

// CustomEventSpecifications mocks base method.
func (m *MockInstanaAPI) CustomEventSpecifications() restapi.RestResource[*restapi.CustomEventSpecification] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomEventSpecifications")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.CustomEventSpecification])
	return ret0
}

// CustomEventSpecifications indicates an expected call of CustomEventSpecifications.
func (mr *MockInstanaAPIMockRecorder) CustomEventSpecifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomEventSpecifications", reflect.TypeOf((*MockInstanaAPI)(nil).CustomEventSpecifications))
}

// GlobalApplicationAlertConfigs mocks base method.
func (m *MockInstanaAPI) GlobalApplicationAlertConfigs() restapi.RestResource[*restapi.ApplicationAlertConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalApplicationAlertConfigs")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.ApplicationAlertConfig])
	return ret0
}

// GlobalApplicationAlertConfigs indicates an expected call of GlobalApplicationAlertConfigs.
func (mr *MockInstanaAPIMockRecorder) GlobalApplicationAlertConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalApplicationAlertConfigs", reflect.TypeOf((*MockInstanaAPI)(nil).GlobalApplicationAlertConfigs))
}

// Groups mocks base method.
func (m *MockInstanaAPI) Groups() restapi.RestResource[*restapi.Group] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Groups")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.Group])
	return ret0
}

// Groups indicates an expected call of Groups.
func (mr *MockInstanaAPIMockRecorder) Groups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Groups", reflect.TypeOf((*MockInstanaAPI)(nil).Groups))
}

// SliConfigs mocks base method.
func (m *MockInstanaAPI) SliConfigs() restapi.RestResource[*restapi.SliConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliConfigs")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.SliConfig])
	return ret0
}

// SliConfigs indicates an expected call of SliConfigs.
func (mr *MockInstanaAPIMockRecorder) SliConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliConfigs", reflect.TypeOf((*MockInstanaAPI)(nil).SliConfigs))
}

// SyntheticLocation mocks base method.
func (m *MockInstanaAPI) SyntheticLocation() restapi.ReadOnlyRestResource[*restapi.SyntheticLocation] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyntheticLocation")
	ret0, _ := ret[0].(restapi.ReadOnlyRestResource[*restapi.SyntheticLocation])
	return ret0
}

// SyntheticLocation indicates an expected call of SyntheticLocation.
func (mr *MockInstanaAPIMockRecorder) SyntheticLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntheticLocation", reflect.TypeOf((*MockInstanaAPI)(nil).SyntheticLocation))
}

// SyntheticTest mocks base method.
func (m *MockInstanaAPI) SyntheticTest() restapi.RestResource[*restapi.SyntheticTest] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyntheticTest")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.SyntheticTest])
	return ret0
}

// SyntheticTest indicates an expected call of SyntheticTest.
func (mr *MockInstanaAPIMockRecorder) SyntheticTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntheticTest", reflect.TypeOf((*MockInstanaAPI)(nil).SyntheticTest))
}

// WebsiteAlertConfig mocks base method.
func (m *MockInstanaAPI) WebsiteAlertConfig() restapi.RestResource[*restapi.WebsiteAlertConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebsiteAlertConfig")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.WebsiteAlertConfig])
	return ret0
}

// WebsiteAlertConfig indicates an expected call of WebsiteAlertConfig.
func (mr *MockInstanaAPIMockRecorder) WebsiteAlertConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebsiteAlertConfig", reflect.TypeOf((*MockInstanaAPI)(nil).WebsiteAlertConfig))
}

// WebsiteMonitoringConfig mocks base method.
func (m *MockInstanaAPI) WebsiteMonitoringConfig() restapi.RestResource[*restapi.WebsiteMonitoringConfig] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebsiteMonitoringConfig")
	ret0, _ := ret[0].(restapi.RestResource[*restapi.WebsiteMonitoringConfig])
	return ret0
}

// WebsiteMonitoringConfig indicates an expected call of WebsiteMonitoringConfig.
func (mr *MockInstanaAPIMockRecorder) WebsiteMonitoringConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebsiteMonitoringConfig", reflect.TypeOf((*MockInstanaAPI)(nil).WebsiteMonitoringConfig))
}
