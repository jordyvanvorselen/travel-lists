package service

import (
	"context"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
)

func CreateList(ctx context.Context, list domain.List) (domain.List, error) {
	return repository.Save(ctx, list)
}

func GetListByUUID(ctx context.Context, uuid string) (domain.List, error) {
	return repository.GetByUUID(ctx, uuid)
}
