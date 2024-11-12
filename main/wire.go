package main

import (
	"github.com/google/wire"
	"transactions_routine/api/handlers"
	"transactions_routine/api/services"
	"transactions_routine/config"
	"transactions_routine/database"
	"transactions_routine/database/models"
)

func InitDependencies() (*Server, error) {
	wire.Build(
		config.NewConfig, // Provides Config
		database.NewDB,   // Provides Database
		models.NewAccount,
		models.NewTransaction,
		models.NewOperationType,
		services.NewAccountService,     // Provides AccountService
		services.NewTransactionService, // Provides TransactionService
		services.NewOperationService,   // Provides OperationService
		handlers.NewAccountHandler,     // Provides AccountHandler
		handlers.NewTransactionHandler, // Provides TransactionHandler
		handlers.NewOperationHandler,   // Provides OperationHandler
		NewServer,                      // Provides Server (your main server setup)
	)
	return &Server{}, nil
}
