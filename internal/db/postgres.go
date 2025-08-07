// internal/db/db.go

package db

import (
	"fmt"
	"log"

	"exams/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	c := config.AppConfig
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
