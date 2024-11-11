package models

import (
	"errors"
	"gorm.io/gorm"
)

//type AccountInterface interface {
//	CreateAccount(db *gorm.DB, documentNumber string) Account
//	FetchAccount(db *gorm.DB, id int) (Account, error)
//}

type Account struct {
	AccountID      uint          `gorm:"primaryKey;column:account_id"`
	DocumentNumber string        `gorm:"column:document_number;not null;unique"`
	Transactions   []Transaction `gorm:"foreignKey:AccountID"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) CreateAccount(db *gorm.DB, documentNumber string) Account {
	account := Account{DocumentNumber: documentNumber}
	db.Create(&account)
	return account
}

func (a *Account) FetchAccount(db *gorm.DB, id int) (Account, error) {
	var account Account
	if err := db.First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return account, err // Not found case
		}
		return account, err // Other error cases
	}
	return account, nil
}
