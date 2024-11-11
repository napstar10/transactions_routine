package services

import (
	"gorm.io/gorm"
	"transactions_routine/api/dto"
	"transactions_routine/database/models"
)

type TransactionService interface {
	CreateTransaction(transaction dto.CreateTransactionRequest) error
	GetTransaction(id int) (models.Transaction, error)
}

type transactionService struct {
	db               *gorm.DB // Injected dependency
	transactionModel *models.Transaction
}

func NewTransactionService(db *gorm.DB, transactionModel *models.Transaction) TransactionService {
	//todo : enable when constraint needs to be updated
	//if err := transactionModel.AddConstraints(db); err != nil {
	//	panic(err)
	//}
	return &transactionService{db: db, transactionModel: transactionModel}
}

func (t transactionService) CreateTransaction(transactionRequest dto.CreateTransactionRequest) error {
	_, err := t.transactionModel.CreateTransaction(t.db, transactionRequest.AccountId, transactionRequest.OperationTypeId, transactionRequest.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t transactionService) GetTransaction(id int) (models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
