package dto

import "transactions_routine/database/models"

type CreateTransactionRequest struct {
	AccountId       uint    `json:"account_id"`
	OperationTypeId uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type GetTransactionsRequest struct {
	AccountId uint `json:"account_id"`
	LastId    uint `json:"last_id"` // for cursor based pagination
}

type PaginatedTransactions struct {
	Transactions []models.Transaction
	NextCursor   uint // The ID of the last transaction on this page, used as the next cursor
}
