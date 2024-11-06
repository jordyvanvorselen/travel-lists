package service

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
	"gorm.io/gorm"
)

func CreateList(db *gorm.DB, list domain.List) (domain.List, error) {
	return repository.SaveList(db, list)
}

func GetListByUUID(db *gorm.DB, uuid string) (*domain.List, error) {
	return repository.GetListByUUID(db, uuid)
}
