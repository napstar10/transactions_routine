package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"transactions_routine/api/dto"
	"transactions_routine/api/services"
)

type AccountHandler struct {
	// Injected service
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var request dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.accountService.CreateAccount(request.DocNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account created", "documentNum": account.DocumentNumber})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	account, err := h.accountService.GetAccount(uint(id))

	if err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "account": account})
}
