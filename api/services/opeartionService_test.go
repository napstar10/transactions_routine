package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"transactions_routine/api/dto"
	"transactions_routine/database/models"
)

type OperationTypeServiceTestSuite struct {
	suite.Suite
	service            OperationService
	db                 *gorm.DB
	operationTypeModel *models.MockOperationTypeModel
}

func TestMyOperationTypeServiceSuite(t *testing.T) {
	suite.Run(t, new(OperationTypeServiceTestSuite))
}

func (suite *OperationTypeServiceTestSuite) SetupTest() {
	suite.db = new(gorm.DB)
	suite.operationTypeModel = new(models.MockOperationTypeModel)
	suite.service = NewOperationService(suite.db, suite.operationTypeModel)
}

func (suite *OperationTypeServiceTestSuite) TestOperationTypeService_CreateOperation_1() {

	// Mocking CreateOperationType setting expectation for the mock
	suite.operationTypeModel.On("CreateOperationType", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	data := dto.CreateOperationTypeRequest{
		OperationTypeId: 0,
		Description:     "Withdrawal",
	}

	// Call the service method to be tested
	err := suite.service.CreateOperationType(data)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *OperationTypeServiceTestSuite) TestOperationTypeService_CreateOperation_2() {

	// Mocking CreateOperationType & setting expectation for the mock
	suite.operationTypeModel.On("CreateOperationType", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("check Constraint violation"))

	data := dto.CreateOperationTypeRequest{
		OperationTypeId: 0,
		Description:     "Withdrawal",
	}

	// Call the CreateOperationType service method
	err := suite.service.CreateOperationType(data)

	// Assertions
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "check Constraint violation", err.Error())
}
