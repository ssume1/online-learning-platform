package database

import (
	"fmt"
	"time"

	"github.com/sandbox-science/online-learning-platform/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs necessary migrations.
func InitDB(cfg entity.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode)

	var err error

	// Retry connection
	for i := 0; i < 3; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after retries: %v", err)
	}

	// Perform schema migrations
	return DatabaseMigration(DB)
}

// DatabaseMigration performs the automatic migration for the schema entities
func DatabaseMigration(DB *gorm.DB) (*gorm.DB, error) {
	err := DB.AutoMigrate(&entity.Account{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Migrated database!")

	return DB, nil
}
