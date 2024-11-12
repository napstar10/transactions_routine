package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transactions_routine/api/dto"
	"transactions_routine/api/services"
)

type OperationHandler struct {
	// Injected service
	operationService services.OperationService
}

func NewOperationHandler(accountService services.OperationService) *OperationHandler {
	return &OperationHandler{operationService: accountService}
}

func (h *OperationHandler) CreateOperationType(c *gin.Context) {
	var request dto.CreateOperationTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.operationService.CreateOperationType(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation Type created"})
}
