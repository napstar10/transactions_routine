package services

import (
	"gorm.io/gorm"
	"transactions_routine/database/models"
)

type AccountService interface {
	CreateAccount(documentNo string) (models.Account, error)
	GetAccount(id uint) (interface{}, error)
}

type accountService struct {
	// Injected dependencies
	db           *gorm.DB
	accountModel models.AccountInterface
}

func NewAccountService(db *gorm.DB, accountModel models.AccountInterface) AccountService {
	return &accountService{db: db, accountModel: accountModel}
}

func (a accountService) CreateAccount(documentNo string) (models.Account, error) {
	account := a.accountModel.CreateAccount(a.db, documentNo)
	return account, nil

}

func (a accountService) GetAccount(id uint) (interface{}, error) {
	account, err := a.accountModel.FetchAccount(a.db, id)

	if err != nil {
		return models.Account{}, err
	}

	data := map[string]interface{}{
		"accountId":  account.AccountID,
		"documentNo": account.DocumentNumber,
	}
	return data, nil
}
