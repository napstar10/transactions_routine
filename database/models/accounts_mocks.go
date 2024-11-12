package models

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockAccountsModel struct {
	mock.Mock
}

func (m *MockAccountsModel) CreateAccount(db *gorm.DB, documentNumber string) Account {
	args := m.Called(db, documentNumber)
	return args.Get(0).(Account)
}

func (m *MockAccountsModel) FetchAccount(db *gorm.DB, id uint) (Account, error) {
	args := m.Called(db, id)
	return args.Get(0).(Account), args.Error(1)
}
