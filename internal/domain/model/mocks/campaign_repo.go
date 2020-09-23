// Code generated by MockGen. DO NOT EDIT.
// Source: campaign_repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "blue-discount/internal/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCampaignRepo is a mock of CampaignRepo interface
type MockCampaignRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignRepoMockRecorder
}

// MockCampaignRepoMockRecorder is the mock recorder for MockCampaignRepo
type MockCampaignRepoMockRecorder struct {
	mock *MockCampaignRepo
}

// NewMockCampaignRepo creates a new mock instance
func NewMockCampaignRepo(ctrl *gomock.Controller) *MockCampaignRepo {
	mock := &MockCampaignRepo{ctrl: ctrl}
	mock.recorder = &MockCampaignRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCampaignRepo) EXPECT() *MockCampaignRepoMockRecorder {
	return m.recorder
}

// FindByActive mocks base method
func (m *MockCampaignRepo) FindByActive(arg0 bool) ([]model.Campaign, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActive", arg0)
	ret0, _ := ret[0].([]model.Campaign)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive
func (mr *MockCampaignRepoMockRecorder) FindByActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockCampaignRepo)(nil).FindByActive), arg0)
}
