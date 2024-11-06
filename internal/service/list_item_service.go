package service

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
	"gorm.io/gorm"
)

func CreateListItem(db *gorm.DB, listItem domain.ListItem) (domain.ListItem, error) {
	return repository.SaveListItem(db, listItem)
}

func GetListItemsByListId(db *gorm.DB, id uint) ([]domain.ListItem, error) {
	return repository.GetListItemsByListId(db, id)
}
