package services

import (
	"gorm.io/gorm"
	"transactions_routine/api/dto"
	"transactions_routine/database/models"
)

type OperationService interface {
	CreateOperation(request dto.CreateOperationTypeRequest) error
}

type operationService struct {
	db                 *gorm.DB
	operationTypeModel *models.OperationType
}

func NewOperationService(db *gorm.DB, operationTypeModel *models.OperationType) OperationService {
	return &operationService{db: db, operationTypeModel: operationTypeModel}
}

func (o operationService) CreateOperation(request dto.CreateOperationTypeRequest) error {
	err := o.operationTypeModel.CreateOperationType(o.db, request.OperationTypeId, request.Description)

	if err != nil {
		return err
	}

	return nil
}
