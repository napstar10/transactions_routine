package services

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"transactions_routine/api/dto"
	"transactions_routine/database/models"
)

type TransactionService interface {
	CreateTransaction(transaction dto.CreateTransactionRequest) error
	GetTransaction(id int) (models.Transaction, error)
	GetTransactions(ctx context.Context, accountId uint, lastId uint, limit int) (dto.PaginatedTransactions, error)
}

type transactionService struct {
	db               *gorm.DB // Injected dependency
	transactionModel models.TransactionInterface
}

func NewTransactionService(db *gorm.DB, transactionModel models.TransactionInterface) TransactionService {
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

func (t transactionService) GetTransactions(ctx context.Context, accountId uint, lastId uint, limit int) (dto.PaginatedTransactions, error) {
	transactionCh := make(chan dto.PaginatedTransactions)
	errCh := make(chan error)

	go func() {
		defer close(transactionCh)
		defer close(errCh)

		transactions, err := t.transactionModel.GetTransactions(t.db, accountId, lastId, limit)

		if err != nil {
			errCh <- err
		}

		var nextCursor uint
		if len(transactions) > 0 {
			nextCursor = transactions[len(transactions)-1].TransactionID
		}

		transactionCh <- dto.PaginatedTransactions{
			Transactions: transactions,
			NextCursor:   nextCursor,
		}

	}()

	select {
	case paginatedTransactions := <-transactionCh:
		fmt.Println("paginatedTransactions:", paginatedTransactions)
		return paginatedTransactions, nil
	case err := <-errCh:
		fmt.Println("Error:", err)
		return dto.PaginatedTransactions{}, err
	case <-ctx.Done():
		return dto.PaginatedTransactions{}, ctx.Err()
	}
}
