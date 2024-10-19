package repository

import (
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
)

func Save(list domain.List) (domain.List, error) {
	return list, nil
}
