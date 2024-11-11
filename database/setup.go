package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"transactions_routine/config"
	"transactions_routine/database/models"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate will create the tables if they don't exist
	if err := db.AutoMigrate(&models.Account{}, &models.OperationType{}, &models.Transaction{}); err != nil {
		panic("failed to migrate database schema")
	}

	return db, nil
}
