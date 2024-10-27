package service

import (
	"context"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/repository"
)

func CreateListItem(ctx context.Context, listItem domain.ListItem) (domain.ListItem, error) {
	return repository.SaveListItem(ctx, listItem)
}

func GetListItemsByListId(ctx context.Context, id int) ([]domain.ListItem, error) {
	return repository.GetListItemsByListId(ctx, id)
}
