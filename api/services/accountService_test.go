package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"transactions_routine/database/models"
)

type AccountServiceTestSuite struct {
	suite.Suite
	service      AccountService
	db           *gorm.DB
	accountModel *models.MockAccountsModel
}

func TestMyAccountServiceSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

func (suite *AccountServiceTestSuite) SetupTest() {
	suite.db = new(gorm.DB)
	suite.accountModel = new(models.MockAccountsModel)
	suite.service = NewAccountService(suite.db, suite.accountModel)
}

func (suite *AccountServiceTestSuite) TestAccountService_GetAccount_1() {
	accountID := uint(1)
	expectedAccount := models.Account{AccountID: accountID, DocumentNumber: "12345678900"}

	// Mocking FetchAccount and Setting expectation for the mock
	suite.accountModel.On("FetchAccount", mock.Anything, mock.Anything).Return(expectedAccount, nil)

	// Call the service method
	result, err := suite.service.GetAccount(accountID)

	var expectedRes interface{} = map[string]interface{}{
		"accountId":  accountID,
		"documentNo": "12345678900",
	}

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRes, result)
	assert.Nil(suite.T(), err)
}

func (suite *AccountServiceTestSuite) TestAccountService_GetAccount_2() {
	accountID := uint(1)
	expectedAccount := models.Account{AccountID: accountID, DocumentNumber: "12345678900"}

	// Mocking FetchAccount and Setting expectation for the mock
	suite.accountModel.On("FetchAccount", mock.Anything, mock.Anything).Return(expectedAccount, nil)

	// Call the service method
	result, err := suite.service.GetAccount(accountID)

	var expectedRes interface{} = map[string]interface{}{
		"accountId":  accountID,
		"documentNo": "12345678999",
	}

	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), expectedRes, result, "Expected mismatch due to different document number")
	assert.Nil(suite.T(), err)
}
