package models

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockOperationTypeModel struct {
	mock.Mock
}

func (m *MockOperationTypeModel) CreateOperationType(db *gorm.DB, operationTypeId uint, description string) error {
	args := m.Called(db, operationTypeId, description)
	return args.Error(0)
}
