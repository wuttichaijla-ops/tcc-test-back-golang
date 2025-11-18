package datasource

import (
	"log"
	"os"
	"test-back-golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the GORM database connection.
// Default: PostgreSQL via DATABASE_URL env (e.g. postgres://user:pass@localhost:5432/dbname?sslmode=disable)
func InitDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("[datasource] WARNING: DATABASE_URL is not set, database will not be connected")
		return
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[datasource] failed to connect database: %v", err)
	}

	// Auto-migrate schemas (product only)
	if err := db.AutoMigrate(&models.ProductCode{}); err != nil {
		log.Fatalf("[datasource] failed to migrate database: %v", err)
	}

	DB = db
	log.Println("[datasource] database connected successfully")
}
