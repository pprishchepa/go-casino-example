// Code generated by MockGen. DO NOT EDIT.
// Source: wallet.go
//
// Generated by this command:
//
//	mockgen -source=wallet.go -destination=wallet_mock_test.go -package=v1_test
//

// Package v1_test is a generated GoMock package.
package v1_test

import (
	context "context"
	reflect "reflect"

	domain "github.com/pprishchepa/go-casino-example/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockWalletService is a mock of WalletService interface.
type MockWalletService struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceMockRecorder
}

// MockWalletServiceMockRecorder is the mock recorder for MockWalletService.
type MockWalletServiceMockRecorder struct {
	mock *MockWalletService
}

// NewMockWalletService creates a new mock instance.
func NewMockWalletService(ctrl *gomock.Controller) *MockWalletService {
	mock := &MockWalletService{ctrl: ctrl}
	mock.recorder = &MockWalletServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletService) EXPECT() *MockWalletServiceMockRecorder {
	return m.recorder
}

// CreditMoney mocks base method.
func (m *MockWalletService) CreditMoney(ctx context.Context, entry domain.CreditEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditMoney", ctx, entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditMoney indicates an expected call of CreditMoney.
func (mr *MockWalletServiceMockRecorder) CreditMoney(ctx, entry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditMoney", reflect.TypeOf((*MockWalletService)(nil).CreditMoney), ctx, entry)
}

// DebitMoney mocks base method.
func (m *MockWalletService) DebitMoney(ctx context.Context, entry domain.DebitEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebitMoney", ctx, entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DebitMoney indicates an expected call of DebitMoney.
func (mr *MockWalletServiceMockRecorder) DebitMoney(ctx, entry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebitMoney", reflect.TypeOf((*MockWalletService)(nil).DebitMoney), ctx, entry)
}

// GetBalance mocks base method.
func (m *MockWalletService) GetBalance(ctx context.Context, walletID int) (*domain.WalletBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, walletID)
	ret0, _ := ret[0].(*domain.WalletBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockWalletServiceMockRecorder) GetBalance(ctx, walletID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockWalletService)(nil).GetBalance), ctx, walletID)
}
