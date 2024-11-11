package dto

type CreateTransactionRequest struct {
	AccountId       uint    `json:"account_id"`
	OperationTypeId uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
