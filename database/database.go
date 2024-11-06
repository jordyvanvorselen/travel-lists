package database

import (
	"os"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	db.AutoMigrate(&domain.List{}, &domain.ListItem{})

	return db
}
