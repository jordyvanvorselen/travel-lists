package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s port=5432 dbname=travel-lists sslmode=%s",
		os.Getenv("PSQL_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	return db
}
