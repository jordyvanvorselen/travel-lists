package repository

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"gorm.io/gorm"
)

func SaveListItem(db *gorm.DB, listItem domain.ListItem) (domain.ListItem, error) {
	if db := db.Create(&listItem); db.Error != nil {
		return domain.ListItem{}, db.Error
	}

	return listItem, nil
}

func GetListItemsByListId(db *gorm.DB, id uint) ([]domain.ListItem, error) {
	var items []domain.ListItem

	if db := db.Where("list_id = ?", id).Find(&items); db.Error != nil {
		return nil, db.Error
	}

	return items, nil
}
