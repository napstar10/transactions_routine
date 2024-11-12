// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"transactions_routine/api/handlers"
	"transactions_routine/api/services"
	"transactions_routine/config"
	"transactions_routine/database"
	"transactions_routine/database/models"
)

// Injectors from wire.go:

func InitDependencies() (*Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := database.NewDB(configConfig)
	if err != nil {
		return nil, err
	}
	accountInterface := models.NewAccount()
	accountService := services.NewAccountService(db, accountInterface)
	accountHandler := handlers.NewAccountHandler(accountService)
	transactionInterface := models.NewTransaction()
	transactionService := services.NewTransactionService(db, transactionInterface)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	operationTypeInterface := models.NewOperationType()
	operationService := services.NewOperationService(db, operationTypeInterface)
	operationHandler := handlers.NewOperationHandler(operationService)
	server := NewServer(accountHandler, transactionHandler, operationHandler)
	return server, nil
}
