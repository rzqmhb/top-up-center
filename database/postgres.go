package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	// get db creds from loaded env
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")
)

// connecting to postgres database using gorm with the provided credentials
func ConnectDB() (*gorm.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
									host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}