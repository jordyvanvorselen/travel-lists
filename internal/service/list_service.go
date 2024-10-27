package service

import (
	"context"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
)

func CreateList(ctx context.Context, list domain.List) (domain.List, error) {
	return repository.SaveList(ctx, list)
}

func GetListByUUID(ctx context.Context, uuid string) (domain.List, error) {
	return repository.GetListByUUID(ctx, uuid)
}
