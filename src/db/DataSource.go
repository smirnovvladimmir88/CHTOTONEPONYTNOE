/*
 * DataSource.go
 * Copyright (c) ti-bone 2023-2024
 */

package db

import (
	"feedbackBot/src/config"
	"feedbackBot/src/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection - the DB connection
var Connection *gorm.DB

// Init - initializes the DB connection, auto-migrates the users table and stores the connection in Connection
func Init() {
	var err error
	Connection, err = gorm.Open(postgres.Open(config.CurrentConfig.DbDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	log.SetOutput(os.Stdout)

	log.Println("Trying to auto-migrate users table...")
	err = Connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Errorf("failed to auto-migrate users table: %w", err))
	}

	log.Println("Auto-migrated users table.")

	log.Println("Trying to auto-migrate notes table...")
	err = Connection.AutoMigrate(&models.Note{})
	if err != nil {
		panic(fmt.Errorf("failed to auto-migrate notes table: %w", err))
	}

	log.Println("Auto-migrated notes table.")

	log.Println("Trying to auto-migrate messages table...")
	err = Connection.AutoMigrate(&models.Message{})
	if err != nil {
		panic(fmt.Errorf("failed to auto-migrate messages table: %w", err))
	}

	log.Println("Auto-migrated messages table, successfully connected to the DB.")
}

// GetAllAdmins - извлекает всех администраторов из базы данных
func GetAllAdmins() ([]*models.User, error) {
	var admins []*models.User
	res := Connection.Where("is_admin = ?", true).Find(&admins)
	if res.Error != nil {
			return nil, res.Error
	}
	return admins, nil
}