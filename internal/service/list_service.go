package service

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
)

func CreateList(list domain.List) (domain.List, error) {
	return repository.Save(list)
}
