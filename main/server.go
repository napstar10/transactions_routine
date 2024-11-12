package main

import (
	"github.com/gin-gonic/gin"
	"transactions_routine/api/handlers"
	"transactions_routine/config"
)

type Server struct {
	router             *gin.Engine
	accountHandler     *handlers.AccountHandler
	transactionHandler *handlers.TransactionHandler
	config             *config.Config
}

func NewServer(conf *config.Config, accountHandler *handlers.AccountHandler, transactionHandler *handlers.TransactionHandler, operationHandler *handlers.OperationHandler) *Server {
	router := gin.Default()

	// Register routes
	v1 := router.Group("/v1")
	{
		v1.POST("/accounts", accountHandler.CreateAccount)
		v1.GET("/accounts/:id", accountHandler.GetAccount)
		v1.POST("/transactions", transactionHandler.CreateTransaction)
		v1.POST("/operation/types", operationHandler.CreateOperationType)
	}

	return &Server{router: router, config: conf}
}

func (s *Server) Run(port string) error {
	return s.router.Run(port)
}
