package gorm

import (
	"fmt"
	"log"

	"github.com/khunmostz/be-liongate-go/app/adapter/config"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitPostgresDB initializes and returns a GORM DB instance for PostgreSQL
func InitPostgresDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
		cfg.Postgres.TimeZone,
	)

	// Configure GORM with improved settings
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	// Auto migrate the models in the correct order to avoid foreign key issues
	// First migrate tables without foreign keys, then tables with foreign keys
	if err := db.AutoMigrate(
		&domain.Users{},
		&domain.Animals{},
		&domain.PerformanceStage{},
	); err != nil {
		log.Fatal("Failed to auto migrate base models:", err)
	}

	// Then migrate tables with foreign keys
	if err := db.AutoMigrate(
		&domain.ShowRounds{},
		&domain.Bookings{},
	); err != nil {
		log.Fatal("Failed to auto migrate models with foreign keys:", err)
	}

	// Get the underlying SQL DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB:", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
