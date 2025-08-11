package db

import (
	"go-service/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(
		&models.Aircraft{},
		&models.Route{},
		&models.FlightOp{},
		&models.Schedule{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}