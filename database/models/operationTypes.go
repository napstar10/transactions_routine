package models

import "gorm.io/gorm"

const (
	NormalPurchase           uint = 1
	PurchaseWithInstallments      = 2
	Withdrawal                    = 3
	CreditVoucher                 = 4
)

type OperationType struct {
	OperationTypeID uint          `gorm:"primaryKey;column:operation_type_id"`
	Description     string        `gorm:"column:description"`
	Transactions    []Transaction `gorm:"foreignKey:OperationTypeID"`
}

type OperationTypeInterface interface {
	CreateOperationType(db *gorm.DB, operationTypeId uint, description string) error
}

func NewOperationType() OperationTypeInterface {
	return &OperationType{}
}

func (o *OperationType) CreateOperationType(db *gorm.DB, operationTypeId uint, description string) error {
	operation := OperationType{OperationTypeID: operationTypeId, Description: description}
	if err := db.Create(&operation).Error; err != nil {
		return err
	}
	return nil
}
