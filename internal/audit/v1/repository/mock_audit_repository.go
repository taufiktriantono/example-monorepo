// Code generated by MockGen. DO NOT EDIT.
// Source: audit_repository.go
//
// Generated by this command:
//
//	mockgen -source=audit_repository.go -destination=mock_audit_repository.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/taufiktriantono/api-first-monorepo/internal/audit/v1/domain"
	option "github.com/taufiktriantono/api-first-monorepo/pkg/db/option"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockAuditRepository is a mock of AuditRepository interface.
type MockAuditRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuditRepositoryMockRecorder
	isgomock struct{}
}

// MockAuditRepositoryMockRecorder is the mock recorder for MockAuditRepository.
type MockAuditRepositoryMockRecorder struct {
	mock *MockAuditRepository
}

// NewMockAuditRepository creates a new mock instance.
func NewMockAuditRepository(ctrl *gomock.Controller) *MockAuditRepository {
	mock := &MockAuditRepository{ctrl: ctrl}
	mock.recorder = &MockAuditRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditRepository) EXPECT() *MockAuditRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockAuditRepository) Find(arg0 context.Context, arg1 *domain.AuditLog, arg2 ...option.QueryOption) ([]*domain.AuditLog, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*domain.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAuditRepositoryMockRecorder) Find(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAuditRepository)(nil).Find), varargs...)
}

// WithTrx mocks base method.
func (m *MockAuditRepository) WithTrx(tx *gorm.DB) AuditRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTrx", tx)
	ret0, _ := ret[0].(AuditRepository)
	return ret0
}

// WithTrx indicates an expected call of WithTrx.
func (mr *MockAuditRepositoryMockRecorder) WithTrx(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTrx", reflect.TypeOf((*MockAuditRepository)(nil).WithTrx), tx)
}
