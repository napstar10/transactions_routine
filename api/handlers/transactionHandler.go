package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transactions_routine/api/dto"
	"transactions_routine/api/services"
)

type TransactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transactionService services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var request dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.transactionService.CreateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var request dto.GetTransactionsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paginatedTransactions, err := h.transactionService.GetTransactions(request.AccountId, request.LastId, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": paginatedTransactions})
}
