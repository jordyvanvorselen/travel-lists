package repository

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"gorm.io/gorm"
)

func SaveList(db *gorm.DB, list domain.List) (domain.List, error) {
	if err := db.Create(&list).Error; err != nil {
		return domain.List{}, db.Error
	}

	return list, nil
}

func GetListByUUID(db *gorm.DB, uuid string) (*domain.List, error) {
	var list *domain.List

	if err := db.First(&list, "uuid = ?", uuid).Error; err != nil {
		return &domain.List{}, db.Error
	}

	return list, nil
}
