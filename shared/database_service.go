package shared

import (
	"log"
	"os"
	"reuros-api/internal/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully")

	// Enable the uuid-ossp extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&users.User{}); err != nil {
		return nil, err
	}

	log.Println("Database migration completed")

	return db, nil
}
