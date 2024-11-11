package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	TransactionID   uint      `gorm:"primaryKey;column:transaction_id"`
	AccountID       uint      `gorm:"column:account_id"`        // Foreign key to Accounts
	OperationTypeID uint      `gorm:"column:operation_type_id"` // Foreign key to OperationTypes
	Amount          float64   `gorm:"column:amount"`
	EventDate       time.Time `gorm:"column:event_date"`
	//Account         Account       `gorm:"foreignKey:AccountID"`
	//OperationType   OperationType `gorm:"foreignKey:OperationTypeID"`
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) AddConstraints(db *gorm.DB) error {
	// Apply GORM migrations
	err := db.AutoMigrate(&Transaction{})
	if err != nil {
		return err
	}

	return db.Exec(`
        ALTER TABLE transactions
        ADD CONSTRAINT check_amount_positive_for_type_4
		CHECK (
			(operation_type_id = 4 AND amount > 0) OR
			(operation_type_id != 4 AND amount < 0)
		)`).Error
}

func (t *Transaction) CreateTransaction(db *gorm.DB, accountID uint, operationTypeID uint, amount float64) (Transaction, error) {
	transaction := Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}
	if err := db.Create(&transaction).Error; err != nil {

		return Transaction{}, err
	}

	return transaction, nil
}
