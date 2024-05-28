package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// get DSN for postgres database from loaded .env file
var dsn = os.Getenv("DSN")

// connecting to postgres database using gorm with the provided credentials
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}