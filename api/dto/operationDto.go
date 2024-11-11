package dto

type CreateOperationTypeRequest struct {
	OperationTypeId uint   `json:"operation_type_id"`
	Description     string `json:"description"`
}
