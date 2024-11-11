package services

import (
	"gorm.io/gorm"
	"transactions_routine/database/models"
)

type AccountService interface {
	CreateAccount(documentNo string) (models.Account, error)
	GetAccount(id int) (interface{}, error)
}

type accountService struct {
	db           *gorm.DB // Injected dependency (e.g., DB connection)
	accountModel *models.Account
}

func NewAccountService(db *gorm.DB, accountModel *models.Account) AccountService {
	return &accountService{db: db, accountModel: accountModel}
}

func (a accountService) CreateAccount(documentNo string) (models.Account, error) {
	account := a.accountModel.CreateAccount(a.db, documentNo)
	return account, nil

}

func (a accountService) GetAccount(id int) (interface{}, error) {
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
